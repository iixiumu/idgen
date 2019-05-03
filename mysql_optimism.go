package main

type MysqlOptimism struct {
}

func (MysqlOptimism) GenID() uint64 {
	for true {
		ins := &MyID{}
		err := MyDB.Where("id = ?", 1).First(ins).Error
		if err != nil {

		}
		db := MyDB.Table("my_id").Where("id = ? and my_id = ?", ins.ID, ins.MyID).Update(MyID{
			MyID: ins.MyID + 1,
		})
		if db.RowsAffected == 1 {
			return ins.MyID + 1
		}
	}
	return 0
}
