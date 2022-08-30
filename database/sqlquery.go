package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// get the number of users
func GetUsersNumber() int {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	count := 0
	result, err := db.Query("SELECT COUNT(*) FROM user")
	if err != nil {
		panic(err)
	}
	for result.Next() {
        if err := result.Scan(&count); err != nil {
            panic(err)
        }
	}
	return count
}

// get a user`s detail
func GetUsersDetail(id string) (UserDetail, error) {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	queryStr := fmt.Sprintf(`
	select id, username, descript, registerdate, phone, privilege, publickey 
	from user
	where id=%s`, id)
	result, err := db.Query(queryStr)
	if err != nil {
		panic(err)
	}
	var userDetail UserDetail;
	for result.Next() {
        if err := result.Scan(
			&userDetail.Id,
			&userDetail.Username,
			&userDetail.Descript, 
			&userDetail.Registerdate, 
			&userDetail.Phone,
			&userDetail.Privilege,
			&userDetail.Publickey); err != nil {
            panic(err)
        }
	}
	if userDetail.Id == 0 {
		return UserDetail{}, errors.New("user not found")
	}
	return userDetail, nil
}

// get the user list
// limit: 100 lines
func GetUsersList() ([103]ListUserData, int) {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	result, err := db.Query("select id, username, registerdate, phone, privilege from user limit 100")
	if err != nil {
		panic(err)
	}
	var listData [103]ListUserData
	counter := 0
	for result.Next() {
        if err := result.Scan(
			&listData[counter].Id,
			&listData[counter].Username, 
			&listData[counter].Registerdate,
			&listData[counter].Phone,
			&listData[counter].Privilege); err != nil {
            panic(err)
        }
		counter += 1
	}
	return listData, counter
}