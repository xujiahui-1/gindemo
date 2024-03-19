package main

type TMG struct {
	ID   uint
	Name string
}

func TransactionTest() {

	tx := GLOBAL_DB.Begin()
	tx.Create(&TMG{2, "加"})
	tx.Create(&TMG{3, "辉"})
	tx.Commit()

	tx.Rollback()
}
