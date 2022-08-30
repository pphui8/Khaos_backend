package test

import (
	"database/sql"
	"khaos/database"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	result, err := db.Query("select id, username, descript from user")
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

func TestDBSumUser(t *testing.T) {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	result, err := db.Query("SELECT COUNT(*) FROM user")
	if err != nil {
		t.Error(err)
	}
	for result.Next() {
		count := 0
        if err := result.Scan(&count); err != nil {
            t.Error(err)
        }
		t.Log(count)
	}
}

func TestGetUserDEtail(t *testing.T) {
	result, err := database.GetUsersDetail("1")
	t.Log("res:", result,"err", err)
	result, err = database.GetUsersDetail("2")
	t.Log("res:", result,"err", err)
}