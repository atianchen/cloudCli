package db

import (
		"time"
		"log"
		badger "github.com/dgraph-io/badger/v3"
	)

var  badgerDb *badger.DB

type BadgerDbHelper struct{

}

func (b *BadgerDbHelper) Connect(dbFile string){
	  db, err := badger.Open(badger.DefaultOptions(dbFile))
	  if err != nil {
		  log.Fatal(err)
	  }
	  badgerDb = db
}

func (b *BadgerDbHelper) Set(key string,value string,expireTime time.Duration) error {
	return badgerDb.Update(func(txn *badger.Txn) error {
	  	if expireTime>0 {
	  		  e := badger.NewEntry([]byte(key), []byte(value)).WithTTL(expireTime)
	  		  return txn.SetEntry(e) 
	  		}else{
	  		  return txn.Set([]byte(key), []byte(value))
	  		}
	})
}

func (b *BadgerDbHelper) Get(key string) (string,error){
  var byteContent []byte
	err := badgerDb.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		byteContent, err = item.ValueCopy(nil)
		return err
	})
	if (byteContent!=nil){
		return string(byteContent),nil
	}else{
		return "",err
	}
}

func (b *BadgerDbHelper) Remove(key string) error {
	return badgerDb.Update(func(txn *badger.Txn) error {
	  	  return txn.Delete([]byte(key))
	})
}

func (b *BadgerDbHelper)  Release(){
	if badgerDb!=nil {
		badgerDb.Close()
	}
}