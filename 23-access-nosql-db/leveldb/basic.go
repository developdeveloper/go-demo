package main

import (
	"23-access-nosql-db/session"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func keyValue() {
	db := session.GetLevelDB()
	defer db.Close()

	db.Put([]byte("1st"), []byte("1"), nil)
	data, _ := db.Get([]byte("1st"), nil)
	fmt.Println(string(data))

	db.Delete([]byte("1st"), nil)
	exists, _ := db.Has([]byte("1st"), nil)
	fmt.Println(exists)
	_, err := db.Get([]byte("1st"), nil)
	fmt.Println(err) // not found
	fmt.Println("------")
}

func iterAll() {
	db := session.GetLevelDB()
	defer db.Close()

	db.Put([]byte("3rd"), []byte("3"), nil)
	db.Put([]byte("2nd"), []byte("2"), nil)
	db.Put([]byte("1st"), []byte("1"), nil)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("%v %v\n", string(key), string(value))
	}

	iter.Release()
	fmt.Println("------")
}

func iterRange() {
	db := session.GetLevelDB()
	defer db.Close()

	iter := db.NewIterator(&util.Range{Start: []byte("2nd"), Limit: []byte("3rd")}, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("%v %v\n", string(key), string(value))
	}

	iter.Release()
	fmt.Println("------")
}

func seek() {
	db := session.GetLevelDB()
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	for ok := iter.Seek([]byte("2nd")); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("%v %v\n", string(key), string(value))
	}

	iter.Release()
	fmt.Println("------")
}

func batch() {
	db := session.GetLevelDB()
	defer db.Close()
	batch := new(leveldb.Batch)

	batch.Put([]byte("4th"), []byte("4"))
	batch.Put([]byte("5th"), []byte("5"))
	db.Write(batch, nil)
}

func main() {
	keyValue()
	iterAll()
	iterRange()
	seek()
	batch()
}
