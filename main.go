package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	blockchain := NewBlockchain()
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(db)

	cli := CLI{blockchain}
	cli.Run()
}
