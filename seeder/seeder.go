package seeder

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"torrent/handshake"
	"torrent/message"
	"torrent/torrentfile"
)

type Request struct {
	Index      int
	BlockBegin int
	Begin      int
	End        int
	BlockSize  int
}

// Read parses a message from a stream. Returns `nil` on keep-alive message
func readMessage(r io.Reader) (*message.Message, error) {
	lengthBuf := make([]byte, 4)
	_, err := io.ReadFull(r, lengthBuf)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthBuf)

	// keep-alive message
	if length == 0 {
		return nil, nil
	}

	messageBuf := make([]byte, length)
	_, err = io.ReadFull(r, messageBuf)
	if err != nil {
		return nil, err
	}

	m := message.Message{
		ID:      message.MessageID(messageBuf[0]),
		Payload: messageBuf[1:],
	}

	return &m, nil
}

func SendUnchoke(conn net.Conn) error {
	msg := message.Message{ID: message.MsgUnchoke}
	_, err := conn.Write(msg.Serialize())
	return err
}

// Some Remaining shit todo here
func parseRequest(torrent *torrentfile.Torrent, msg *message.Message) (*Request, error) {
	if msg.ID != message.MsgRequest {
		return nil, fmt.Errorf("Expected REQUEST (ID %d), got ID %d", message.MsgRequest, msg.ID)
	}

	if len(msg.Payload) < 12 {
		return nil, fmt.Errorf("Payload too short. %d < 8", len(msg.Payload))
	}
	MaxBound := torrent.PieceLength * len(torrent.PieceHashes)
	index := int(binary.BigEndian.Uint32(msg.Payload[0:4]))
	blockStart := int(binary.BigEndian.Uint32(msg.Payload[4:8]))
	blockSize := int(binary.BigEndian.Uint32(msg.Payload[8:12]))
	begin := index*torrent.PieceLength + blockStart
	end := begin + blockSize

	if end > MaxBound {
		end = MaxBound
		blockSize = end - begin
	}
	request := Request{
		Index:      index,
		BlockBegin: blockStart,
		Begin:      begin,
		End:        end,
		BlockSize:  blockSize,
	}

	return &request, nil
}

func FormatPiece(request *Request, data []byte) *message.Message {
	payload := make([]byte, 8+len(data))
	binary.BigEndian.PutUint32(payload[0:4], uint32(request.Index))
	binary.BigEndian.PutUint32(payload[4:8], uint32(request.BlockBegin))
	for idx := 0; idx < len(data); idx++ {
		payload[8+idx] = data[idx]
	}
	return &message.Message{ID: message.MsgPiece, Payload: payload}
}

func Upload(torrent *torrentfile.Torrent, msg *message.Message, conn net.Conn) error {
	request, err := parseRequest(torrent, msg)
	if err != nil {
		return err
	}
	file, err := os.Open(torrent.Name)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("could not read file due to unexpected error", err)
	}
	data := make([]byte, request.BlockSize)
	_, err = file.ReadAt(data, int64(request.Begin))

	if err != nil {
		return fmt.Errorf("uploading Interrupted due to unexpected error", err)
	}

	message := FormatPiece(request, data)

	conn.Write(message.Serialize())
	log.Printf("Piece %d uploaded for file %s", request.Index, torrent.Name)

	// conn.Write()
	return nil
}

func handleConnection(torrent *torrentfile.Torrent, conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	res, err := handshake.Read(reader)

	if err != nil {
		return
	}
	conn.Write(res.Serialize())
	Payload := torrent.Bitfield
	bitField := message.Message{ID: message.MsgBitfield, Payload: Payload}
	conn.Write(bitField.Serialize())
	_, err = readMessage(reader)
	if err != nil {

		return
	}

	_, err = readMessage(reader)
	if err != nil {
		return
	}

	// maybe add a checker if the number of goroutines hasn't been overloaded
	error := SendUnchoke(conn)

	if error != nil {
		return
	}

	for {
		requestMessage, err := readMessage(reader)
		if err != nil {
			fmt.Errorf("could not read message", requestMessage, err)
			return
		}
		go Upload(torrent, requestMessage, conn)
	}
}

func HandleSeed(torrent *torrentfile.Torrent, Port uint16) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))

	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	log.Printf("Listening on localhost and port: %d ", Port)

	for {
		conn, err := ln.Accept()
		if err == nil {
			log.Println("Accepted Connection", conn.RemoteAddr().String())
			go handleConnection(torrent, conn)
		}
	}
}