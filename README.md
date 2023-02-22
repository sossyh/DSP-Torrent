



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
- 
