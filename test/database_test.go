package test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	result, err := db.Query("select id, username, descript from manager")
	if err != nil {
		t.Error(err)
	}
	counter := 0
	for result.Next() {
		var id, username, descript string
        if err := result.Scan(&id, &username, &descript); err != nil {
            t.Error(err)
        }
		t.Log(counter)
		t.Log("id:", id)
		t.Log("username:", username)
		t.Log("descript:", descript)
		counter += 1
	}
}