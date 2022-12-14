package test

import (
	"database/sql"
	"khaos/conf"
	"khaos/database"
	"strings"
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
	t.Log("res:", result, "err", err)
	result, err = database.GetUsersDetail("2")
	t.Log("res:", result, "err", err)
}

func TestUserList(t *testing.T) {
	// db, err := sql.Open("mysql", "root:123212321@tcp(127.0.0.1:3306)/khaos?charset=utf8")
	// if err != nil {
	// 	panic(err)
	// }
	// result, err := db.Query("select id, username, registerdate, phone, privilege from user limit 100")
	// if err != nil {
	// 	t.Error(err)
	// }
	// var listData [103]database.ListUserData
	// counter := 0
	// for result.Next() {
    //     if err := result.Scan(
	// 		&listData[counter].Id,
	// 		&listData[counter].Username, 
	// 		&listData[counter].Registerdate,
	// 		&listData[counter].Phone,
	// 		&listData[counter].Privilege); err != nil {
    //         t.Error(err)
    //     }
	// 	t.Log(listData[counter])
	// 	counter += 1
	// }
	var result [53]database.ListUserData
	var len int
	database.GetUsersList(&result, &len)
	t.Log(len)
}

func TestDeluser(t *testing.T) {
	database.InitDB()
	err := database.DelUser("1")
	if err != nil {
		t.Error(err)
	}
}

func TestLogin(t *testing.T) {
	result, err := database.Login("72DCE6E8801031109C58A89389BAFB86 4")
	t.Log("result:", result, "err:", err)
}

func TestProductList(t *testing.T) {
	var result [53]database.ListProduct
	var len int
	database.GetProductList(&result, &len)
	t.Log("result:", result[:len], "len:", len)
}

func TestDelProduct(t *testing.T) {
	err := database.DelProduct("1")
	if err != nil {
		t.Error(err)
	}
}


func TestOrder(t *testing.T) {
	var result [53]database.ListOrder
	var len int
	database.GetOrderList(&result, &len)
	t.Log("result:", result[:len], "len:", len)
}

func TestAddProduct(t *testing.T) {
	var product database.ProductDetail
	product.Productname = "?????????"
	product.Price = 9999
	product.Descript = "?????????????????????????????????????????????"
	product.Stock = 1000
	product.Sale = 0
	product.Type = "?????????"
	product.Img = "none"
	product.Status = "??????"

	err := database.AddProduct(product)
	if err != nil {
		t.Error(err)
	}
}

// func TestFindOrderById(t *testing.T) {
// 	result, err := database.FindOrderById("1")
// 	t.Log("result:", result, "err:", err)
// }

func TestGetAnnocements(t *testing.T) {
	var result [53]database.ListAnnouncement
	var len int
	database.GetAnnouncementList(&result, &len)
	t.Log("result:", result[:len], "len:", len)
}

func TestAddAnnouncement(t *testing.T) {
	var announcement database.AnnouncementDetail
	announcement.Title = "?????????"
	announcement.Content = "?????????????????????????????????????????????"
	announcement.Date = "2019-12-12"
	announcement.Img = ""

	err := database.AddAnnouncement(announcement)
	if err != nil {
		t.Error(err)
	}
}

func TestDelAnnouncement(t *testing.T) {
	err := database.DelAnnouncement("2")
	if err != nil {
		t.Error(err)
	}
}

func TestGetPostList(t *testing.T) {
	var result [53]database.ListPost
	var len int = 0
	database.GetPostsList(&result, &len)
	t.Error("result:", result[:len], "len:", len)
}

func TestGetCommentListByPostId(t *testing.T) {
	var result [53]database.ListComment
	var len int = 0
	database.GetCommentListByPostId(&result, &len, "1")
	t.Error("result:", result[:len], "len:", len)
}

func TestInitDB(t *testing.T) {
	conf.InitConf()
	var myconf = conf.MyConf
	t.Error(myconf)
	t.Error(strings.Join([]string{myconf.Database.User, ":", myconf.Database.Password, "@tcp(", myconf.Database.Ip, ":", myconf.Database.Port, ")/", myconf.Database.Database, "?charset=utf8"}, ""))
}

func TestDelPost(t *testing.T) {
	err := database.DelPost("1")
	if err != nil {
		t.Error(err)
	}
}