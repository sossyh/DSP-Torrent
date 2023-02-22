



Group Members:
- Hiwot Derese
- Sosna Worku
- Hayat Tofik
- Tseganesh Yifru
- Ekram Kedir


# Project Description
##   Bencode


- is a format in which the torrent file is written.To get useful parameters that can be used by the trackers and the other peers we need to apply different operations like parsing and marshaling.So under the bencode folder we have added he implementation for the parser and marshal and he corrponding test files for them.
- Parser.go - This file implements functionalities that are needed by the parser such as gettin the important informations from the .torrent file.After                                parsing the whole file , we can spot the URL of the tracker, the creation date (as a Unix timestamp), the name and size of the file, and a                                big binary blob containing the SHA-1 hashes of each piece, which are equally-sized parts of the file we want to download. The exact size of                              a piece varies between torrents, but they are usually somewhere between 256KB and 1MB. This means that a large file might be made up of                                  thousands of pieces. We’ll download these pieces from our peers, together, and we will get all the files needed.
- marshal.go - since the .torrent file is 'bee-encoded' we need to decode it before usig it .The fuction under this file implements the encoding and decoding logic.
##   torrent
- This folder contains  the  tracker,handshake ,peer and download.
- peer.go - used to fill it's bitfield ,read ,write and copyingpieces.
- handShake.go - we should check from the other peers  can communicate using the BitTorrent protocol,
is able to understand and respond to our messages,and 
has the file that we want, or at least knows what we’re talking about.
- tracker.go - Now that we have a list of peers, it’s time to connect with them and start downloading pieces! We can break down the process into a few steps. For each peer, we want to:
Start a TCP connection with the peer. This is like starting phone call.
Complete a two-way BitTorrent handshake. “Hello?” “Hello."
Exchange messages to download pieces. “I’d like piece #444 please
download.go .
- peer.go - This implements the peer it self and he fuctions that it use for communicating and sending for the other peers.Once we’ve completed the initial handshake, we can send and receive messages. Well, not quite—if the other peer isn’t ready to accept messages, we can’t send any until they tell us they’re ready. In this state, we’re considered choked by the other peer. They’ll send us an unchoke message to let us know that we can begin asking them for data. By default, we assume that we’re choked until proven otherwise.
Once we’ve been unchoked, we can then begin sending requests for pieces, and they can send us messages back containing pieces.
bitfield.go - it uses fuctions that help the peer to keep track of the pieces that it can seed .

# Running the project 
## main.go 
- In this file fuctions that implement structuring the data that is about to be sent to the tracker and download it.
## command for running 
Usage
import {subscribe} from '@github/paste-markdown'

// Subscribe the behavior to the textarea.
subscribe(document.querySelector('textarea[data-paste-markdown]'))
Using a library like selector-observer, the behavior can automatically be applied to any element matching a selector.

import {observe} from 'selector-observer'


