package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
	select id, username, descript, registerdate, phone, privilege 
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
			&userDetail.Privilege); err != nil {
            panic(err)
        }
	}
	if userDetail.Id == 0 {
		return UserDetail{}, errors.New("user not found")
	}
	return userDetail, nil
}

// get the user list
// limit: 50 lines
// listData: return value, len: value lengthss
func GetUsersList(listData *[53]ListUserData, len *int) {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	result, err := db.Query("select id, username, registerdate, phone, privilege from user limit 50")
	if err != nil {
		panic(err)
	}
	for result.Next() {
        if err := result.Scan(
			&listData[*len].Id,
			&listData[*len].Username, 
			&listData[*len].Registerdate,
			&listData[*len].Phone,
			&listData[*len].Privilege); err != nil {
            panic(err)
        }
		*len += 1
	}
}

// Login (manager only)
func Login(publickey string) (UserDetail, error) {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	publickey = strings.Trim(publickey, " ")
	queryStr := fmt.Sprintf(`
	select id, username, descript, registerdate, phone, privilege 
    from user
    where publickey='%s' and privilege='manager'`, publickey)
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
			&userDetail.Privilege); err != nil {
            panic(err)
        }
	}
	if userDetail.Id == 0 {
		return UserDetail{}, errors.New("user not found")
	}
	return userDetail, nil
}

// product list
// limit: 50 lines
// listData: return value, len: value lengthss
func GetProductList(listData *[53]ListProduct, len *int) {
		db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	result, err := db.Query(`
	select id, productname, price, descript, stock, sale, type, status
	from product
	limit 50`)
	if err != nil {
		panic(err)
	}
	for result.Next() {
        if err := result.Scan(
			&listData[*len].Id,
			&listData[*len].Productname, 
			&listData[*len].Price,
			&listData[*len].Descript,
			&listData[*len].Stock,
			&listData[*len].Sale,
			&listData[*len].Type,
			&listData[*len].Status); err != nil {
            panic(err)
        }
		*len += 1
	}
}