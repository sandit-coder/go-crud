package repository

import (
	"container/list"
	"database/sql"

)

type Repo interface {
	Select(doc interface{}) (*list.List, error)
	Insert(doc interface{}) (int64, error)
	Update(doc interface{}) (int64, error)
	Remove(doc interface{}) (int64, error)
}


var DbConnection *sql.DB


func SetupRepo() (err error) {
	DbConnection, err = db.GetDatabase()
	return
}

//CloseRepo - close database connections
func CloseRepo() {
	if DbConnection != nil {
		defer DbConnection.Close()
	}
}
