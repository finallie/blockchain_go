package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// NumberToBytes converts any number to a byte slice
func NumberToBytes(value any) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, value)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
