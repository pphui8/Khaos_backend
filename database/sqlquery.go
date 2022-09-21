package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// get the number of users
func GetSummary() Summary {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	var summary Summary = Summary{
		Usernumber: 0,
		Ordernumber: 0,
	}
	var count int
	result, err := db.Query("SELECT COUNT(*) FROM user")
	if err != nil {
		panic(err)
	}
	for result.Next() {
        if err := result.Scan(&count); err != nil {
            panic(err)
        }
	}
	summary.Usernumber = count
	result, err = db.Query("SELECT COUNT(*) FROM orders")
	if err != nil {
		panic(err)
	}
	for result.Next() {
        if err := result.Scan(&count); err != nil {
            panic(err)
        }
	}
	summary.Ordernumber = count
	return summary
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

// del a user by id
func DelUser(id string) error {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	// delete data in orders
	ordersqueryStr := fmt.Sprintf(`
	DELETE FROM orders
	WHERE userid=%s
	`, id)
	_, err = db.Query(ordersqueryStr)
	if err != nil {
		return errors.New("delete user failed")
	}
	// delete data in user
	queryStr := fmt.Sprintf(`
	DELETE FROM user
	WHERE id=%s`, id)
	_, err = db.Query(queryStr)
	if err != nil {
		return errors.New("delete user failed")
	}
	return nil
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
	select id, productname, price, descript, stock, sale, img, type, status
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
			&listData[*len].Img,
			&listData[*len].Type,
			&listData[*len].Status); err != nil {
            panic(err)
        }
		*len += 1
	}
}

// add a product
func AddProduct(productDetail ProductDetail) error {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		return err
	}
	queryStr := fmt.Sprintf(`
	INSERT INTO product (productname, descript, price, stock, sale, img, type, status)
	VALUES ("%s", "%s", %d, %d, %d, "%s", "%s", "%s");
	`,
	productDetail.Productname,
	productDetail.Descript,
	productDetail.Price,
	productDetail.Stock,
	productDetail.Sale,
	productDetail.Img,
	productDetail.Type,
	productDetail.Status)
	_, err = db.Query(queryStr)
	if err != nil {
		return err
	}
	return err
}

// del a product by id
func DelProduct(id string) error {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	// delete data in orders
	ordersqueryStr := fmt.Sprintf(`
	DELETE FROM orders
	WHERE productid=%s
	`, id)
	_, err = db.Query(ordersqueryStr)
	if err != nil {
		return errors.New("delete product failed")
	}
	queryStr := fmt.Sprintf(`
	DELETE FROM product
	WHERE id=%s`, id)
	_, err = db.Query(queryStr)
	if err != nil {
		return errors.New("delete product failed")
	}
	return nil
}

// orders list
// limit: 50 lines
// listData: return value, len: value lengthss
func GetOrderList(listData *[53]ListOrder, len *int) {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	result, err := db.Query(`
	select id, userid, username, productid, productname, price, number, date, location, status
	from orders
	limit 50`)
	if err != nil {
		panic(err)
	}
	for result.Next() {
        if err := result.Scan(
			&listData[*len].Id,
			&listData[*len].Userid,
			&listData[*len].Username,
			&listData[*len].Productid,
			&listData[*len].Productname,
			&listData[*len].Price,
			&listData[*len].Number,
			&listData[*len].Date,
			&listData[*len].Location,
			&listData[*len].Status,); err != nil {
            panic(err)
        }
		*len += 1
	}
}

// orders list
// limit: 50 lines
// listData: return value, len: value lengthss
// func FindOrderListById(listData *[53]ListOrder, len *int, userId string) {
// 	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
// 	if err != nil {
// 		panic(err)
// 	}
// 	result, err := db.Query(`
// 	select id, userid, username, productid, productname, price, number, date, location, status
// 	from orders
// 	where id=` + userId + `
// 	limit 50`)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for result.Next() {
//         if err := result.Scan(
// 			&listData[*len].Id,
// 			&listData[*len].Userid,
// 			&listData[*len].Username,
// 			&listData[*len].Productid,
// 			&listData[*len].Productname,
// 			&listData[*len].Price,
// 			&listData[*len].Number,
// 			&listData[*len].Date,
// 			&listData[*len].Location,
// 			&listData[*len].Status,); err != nil {
//             panic(err)
//         }
// 		*len += 1
// 	}
// }

// del an order by id
func DelOrder(id string) error {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	queryStr := fmt.Sprintf(`
	DELETE FROM orders
	WHERE id=%s`, id)
	_, err = db.Query(queryStr)
	if err != nil {
		return errors.New("delete order failed")
	}
	return nil
}

// get announcement list
// limit: 50 lines
// listData: return value, len: value lengthss
func GetAnnouncementList(listData *[53]ListAnnouncement, len *int) {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	result, err := db.Query(`
	select id, title, content, date
	from announcement
	limit 50`)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		if err := result.Scan(
			&listData[*len].Id,
			&listData[*len].Title,
			&listData[*len].Content,
			&listData[*len].Date,); err != nil {
			panic(err)
		}
		*len += 1
	}
}

// add an announcement
func AddAnnouncement(announcementDetail AnnouncementDetail) error {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		return err
	}
	queryStr := fmt.Sprintf(`
	INSERT INTO announcement (title, content, date)
	VALUES ("%s", "%s", "%s");
	`,
	announcementDetail.Title,
	announcementDetail.Content,
	announcementDetail.Date)
	_, err = db.Query(queryStr)
	if err != nil {
		return err
	}
	return nil
}

// del an announcement by id
func DelAnnouncement(id string) error {
	db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	if err != nil {
		panic(err)
	}
	queryStr := fmt.Sprintf(`
	DELETE FROM announcement
	WHERE id=%s`, id)
	_, err = db.Query(queryStr)
	if err != nil {
		return errors.New("delete announcement failed")
	}
	return nil
}