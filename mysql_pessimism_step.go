package main

import "sync"

type MysqlPessimismStep struct {
}

const STEP = 10

var startID, endID uint64
var lock sync.Mutex

func init() {
	startID = 1
	endID = 0
}

func (MysqlPessimismStep) GenID() uint64 {
	lock.Lock()
	defer lock.Unlock()
	if startID < endID {
		startID++
	} else {
		ins := &MyID{}
		db := MyDB.Begin()
		defer db.Commit()
		err := db.Raw("select * from my_id where id = ? for update", 3).Scan(ins).Error
		if err != nil {

		}
		startID = ins.MyID + 1
		endID = ins.MyID + STEP
		db.Table("my_id").Where("id = ? and my_id = ?", ins.ID, ins.MyID).Update(MyID{
			MyID: ins.MyID + STEP,
		})
	}
	return startID
}
