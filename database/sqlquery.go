package database

import (
	"database/sql"
	"errors"
	"fmt"
	"khaos/conf"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var myconf = conf.MyConf	
    //构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
    path := strings.Join([]string{myconf.Database.User, ":", myconf.Database.Password, "@tcp(", myconf.Database.Ip, ":", myconf.Database.Port, ")/", myconf.Database.Database, "?charset=utf8"}, "")
    //打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
    DB, _ = sql.Open("mysql", path)
    //设置数据库最大连接数
    DB.SetConnMaxLifetime(130)
    //设置上数据库最大闲置连接数
    DB.SetMaxIdleConns(10)
    //验证连接
    if err := DB.Ping(); err != nil {
        fmt.Println("open database fail")
        return
    }
    fmt.Println("connnect success")
}

// get the number of users
func GetSummary() Summary {
	var summary Summary = Summary{
		Usernumber: 0,
		Ordernumber: 0,
	}
	var count int
	result, err := DB.Query("SELECT COUNT(*) FROM user")
	if err != nil {
		panic(err)
	}
	for result.Next() {
        if err := result.Scan(&count); err != nil {
            panic(err)
        }
	}
	summary.Usernumber = count
	result, err = DB.Query("SELECT COUNT(*) FROM orders")
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
	queryStr := fmt.Sprintf(`
	select id, username, descript, registerdate, phone, privilege 
	from user
	where id=%s`, id)
	result, err := DB.Query(queryStr)
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
	// delete data in orders
	ordersqueryStr := fmt.Sprintf(`
	DELETE FROM orders
	WHERE userid=%s
	`, id)
	_, err := DB.Query(ordersqueryStr)
	if err != nil {
		return errors.New("delete user failed")
	}
	// delete data in post
	postqueryStr := fmt.Sprintf(`
	DELETE FROM post
	WHERE userid=%s
	`, id)
	_, err = DB.Query(postqueryStr)
	if err != nil {
		return errors.New("delete user failed")
	}
	// delete data in user
	queryStr := fmt.Sprintf(`
	DELETE FROM user
	WHERE id=%s`, id)
	_, err = DB.Query(queryStr)
	if err != nil {
		return errors.New("delete user failed")
	}
	return nil
}

// get the user list
// limit: 50 lines
// listData: return value, len: value lengthss
func GetUsersList(listData *[53]ListUserData, len *int) {
	result, err := DB.Query("select id, username, registerdate, phone, privilege from user limit 50")
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
	publickey = strings.Trim(publickey, " ")
	queryStr := fmt.Sprintf(`
	select id, username, descript, registerdate, phone, privilege 
    from user
    where publickey='%s' and privilege='manager'`, publickey)
	result, err := DB.Query(queryStr)
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
	result, err := DB.Query(`
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
	_, err := DB.Query(queryStr)
	if err != nil {
		return err
	}
	return err
}

// del a product by id
func DelProduct(id string) error {
	// delete data in orders
	ordersqueryStr := fmt.Sprintf(`
	DELETE FROM orders
	WHERE productid=%s
	`, id)
	_, err := DB.Query(ordersqueryStr)
	if err != nil {
		return errors.New("delete product failed")
	}
	queryStr := fmt.Sprintf(`
	DELETE FROM product
	WHERE id=%s`, id)
	_, err = DB.Query(queryStr)
	if err != nil {
		return errors.New("delete product failed")
	}
	return nil
}

// orders list
// limit: 50 lines
// listData: return value, len: value lengthss
func GetOrderList(listData *[53]ListOrder, len *int) {
	result, err := DB.Query(`
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

// del an order by id
func DelOrder(id string) error {
	queryStr := fmt.Sprintf(`
	DELETE FROM orders
	WHERE id=%s`, id)
	_, err := DB.Query(queryStr)
	if err != nil {
		return errors.New("delete order failed")
	}
	return nil
}

// get announcement list
// limit: 50 lines
// listData: return value, len: value lengthss
func GetAnnouncementList(listData *[53]ListAnnouncement, len *int) {
	result, err := DB.Query(`
	select id, title, content, date, img
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
			&listData[*len].Date,
			&listData[*len].Img); err != nil {
			panic(err)
		}
		*len += 1
	}
}

// add an announcement
func AddAnnouncement(announcementDetail AnnouncementDetail) error {
	queryStr := fmt.Sprintf(`
	INSERT INTO announcement (title, content, date, img)
	VALUES ("%s", "%s", "%s", "%s");
	`,
	announcementDetail.Title,
	announcementDetail.Content,
	announcementDetail.Date,
	announcementDetail.Img)
	_, err := DB.Query(queryStr)
	if err != nil {
		return err
	}
	return nil
}

// del an announcement by id
func DelAnnouncement(id string) error {
	queryStr := fmt.Sprintf(`
	DELETE FROM announcement
	WHERE id=%s`, id)
	_, err := DB.Query(queryStr)
	if err != nil {
		return errors.New("delete announcement failed")
	}
	return nil
}

// get posts list
// limit: 50 lines
// listData: return value, len: value lengthss
func GetPostsList(listData *[53]ListPost, len *int) {
	result, err := DB.Query(`
	select id, userid, username, title, content, browsenumber, date, legal, elite, img, tag
	from post
	limit 50`)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		if err := result.Scan(
			&listData[*len].Id,
			&listData[*len].Userid,
			&listData[*len].Username,
			&listData[*len].Title,
			&listData[*len].Content,
			&listData[*len].Browsenumber,
			&listData[*len].Date,
			&listData[*len].Legal,
			&listData[*len].Elite,
			&listData[*len].Img,
			&listData[*len].Tag); err != nil {
				panic(err)
			}
		*len += 1
	}
}

// get comment list by post id
// limit: 50 lines
// listData: return value, len: value lengthss
func GetCommentListByPostId(listData *[53]ListComment, len *int, postId string) {
	result, err := DB.Query(`
	select id, userid, username, postid, content, date, support, against
	from comment
	where postid = ` + postId + `
	limit 50`)
	if err != nil {
		panic(err)
	}
		for result.Next() {
		if err := result.Scan(
			&listData[*len].Id,
			&listData[*len].Userid,
			&listData[*len].Username,
			&listData[*len].Postid,
			&listData[*len].Content,
			&listData[*len].Date,
			&listData[*len].Support,
			&listData[*len].Against); err != nil {
				panic(err)
			}
		*len += 1
	}
}

// delete a post by id
func DelPost(id string) error {	
	// del comments
	queryStr := fmt.Sprintf(`
	DELETE FROM comment
	WHERE postid=%s`, id)
	_, err := DB.Query(queryStr)
	if err != nil {
		return errors.New("delete post failed")
	}
	// del post
	queryStr = fmt.Sprintf(`
	DELETE FROM post
	WHERE id=%s`, id)
	_, err = DB.Query(queryStr)
	if err != nil {
		return errors.New("delete post failed")
	}
	return nil
}

// delete a comment by id
func DelComment(id string) error {
	queryStr := fmt.Sprintf(`
	DELETE FROM comment
	WHERE id=%s`, id)
	_, err := DB.Query(queryStr)
	if err != nil {
		return errors.New("delete comment failed")
	}
	return nil
}