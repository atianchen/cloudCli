package db

import (
	badger "github.com/dgraph-io/badger/v3"
	"log"
	"time"
)

type BadgerDb struct {
	badgerDb *badger.DB
}

func (b *BadgerDb) Connect(cfg DbConfig) {
	dbFile := cfg.(FileDbConfig).DbFile
	db, err := badger.Open(badger.DefaultOptions(dbFile))
	if err != nil {
		log.Fatal(err)
	}
	b.badgerDb = db
}

func (b *BadgerDb) Set(key string, value string, expireTime time.Duration) error {
	return b.badgerDb.Update(func(txn *badger.Txn) error {
		if expireTime > 0 {
			e := badger.NewEntry([]byte(key), []byte(value)).WithTTL(expireTime)
			return txn.SetEntry(e)
		} else {
			return txn.Set([]byte(key), []byte(value))
		}
	})
}

func (b *BadgerDb) GetBytes(key string) ([]byte, error) {
	var byteContent []byte
	err := b.badgerDb.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		byteContent, err = item.ValueCopy(nil)
		return err
	})
	return byteContent, err
}

func (b *BadgerDb) Get(key string) (string, error) {
	var byteContent []byte
	err := b.badgerDb.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		byteContent, err = item.ValueCopy(nil)
		return err
	})
	if byteContent != nil {
		return string(byteContent), nil
	} else {
		return "", err
	}
}

func (b *BadgerDb) Remove(key string) error {
	return b.badgerDb.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (b *BadgerDb) Release() {
	if b.badgerDb != nil {
		b.badgerDb.Close()
	}
}
