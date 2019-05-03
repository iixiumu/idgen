package main

type MysqlPessimism struct {
}

func (MysqlPessimism) GenID() uint64 {
	ins := &MyID{}
	db := MyDB.Begin()
	defer db.Commit()
	err := db.Raw("select * from my_id where id = ? for update", 2).Scan(ins).Error
	if err != nil {

	}
	db.Table("my_id").Where("id = ? and my_id = ?", ins.ID, ins.MyID).Update(MyID{
		MyID: ins.MyID + 1,
	})
	return ins.MyID + 1
}
