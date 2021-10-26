package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var db leveldb

func batchWrite() leveldb.Batch {
	// batch write
	batch := new(leveldb.Batch)
	batch.Put([]byte("foo"), []byte("value"))
	batch.Put([]byte("bar"), []byte("another value"))
	return batch
}

func main() {
	// db init
	db, err := leveldb.OpenFile("leveldb", nil)
	defer db.Close()

	batch = batchWrite()
	err = db.Write(batch, nil)
	log.Printf("%v", err)

	data, err := db.Get([]byte("foo"), nil)
	fmt.Printf(string(data) + "/n")
	log.Printf("%v", err)

}
