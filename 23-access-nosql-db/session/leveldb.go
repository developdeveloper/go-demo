package session

import "github.com/syndtr/goleveldb/leveldb"

//GetLevelDB 调用方需要 Close 掉
func GetLevelDB() *leveldb.DB {
	db, err := leveldb.OpenFile("/tmp/level.db", nil)
	if err != nil {
		panic(err)
	}
	return db
}
