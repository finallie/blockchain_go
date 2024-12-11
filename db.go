package main

import (
	"github.com/boltdb/bolt"
	"log"
	"sync"
)

const dbFile = "blockchain.db"

const bucketBlocks = "blocks"

const keyLastHash = "l"

var (
	_db   *bolt.DB
	_once sync.Once
)

func GetDb() *bolt.DB {
	if _db == nil {
		_once.Do(initDb)
	}
	return _db
}

func initDb() {
	if _db == nil {
		var err error
		_db, err = bolt.Open(dbFile, 0600, nil)
		if err != nil {
			log.Panic(err)
		}
	}

	err := _db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketBlocks))
		if b == nil {
			var err error
			b, err = tx.CreateBucket([]byte(bucketBlocks))
			if err != nil {
				return err
			}

			genesis := NewGenesisBlock()
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}

			err = b.Put([]byte(keyLastHash), genesis.Hash)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

func GetLastHash() []byte {
	var lastHash []byte

	err := GetDb().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketBlocks))
		lastHash = b.Get([]byte(keyLastHash))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return lastHash
}
