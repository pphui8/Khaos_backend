package database

type UserDetail struct {
	Id           int
	Username     string
	Descript     string
	Registerdate string
	Phone        string
	Privilege    string
}

type Summary struct {
	Usernumber      int
	Ordernumber     int
	PastOrderNumber []int
}

type ListUserData struct {
	Id           int
	Username     string
	Registerdate string
	Phone        string
	Privilege    string // user | manager
}

type LoginData struct {
	Id           int
	Username     string
	Registerdate string
	Phone        string
	Privilege    string
}

type ListProduct struct {
	Id          uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Productname string `gorm:"column:productname" db:"productname" json:"productname" form:"productname"`
	Price       int    `gorm:"column:price" db:"price" json:"price" form:"price"`
	Descript    string `gorm:"column:descript" db:"descript" json:"descript"`
	Stock       uint   `gorm:"column:stock" db:"stock" json:"stock" form:"stock"`
	Sale        uint   `gorm:"column:sale" db:"sale" json:"sale" form:"sale"`
	Type        string `gorm:"column:type" db:"type" json:"type" form:"type"`
	Img         string `gorm:"column:img" db:"img" json:"img" form:"img"`
	// 在售 | 停售
	Status string `gorm:"column:status" db:"status" json:"status" form:"status"`
}

type ProductDetail struct {
	Id          uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Productname string `gorm:"column:productname" db:"productname" json:"productname" form:"productname"`
	Price       int    `gorm:"column:price" db:"price" json:"price" form:"price"`
	Descript    string `gorm:"column:descript" db:"descript" json:"descript"`
	Stock       uint   `gorm:"column:stock" db:"stock" json:"stock" form:"stock"`
	Sale        uint   `gorm:"column:sale" db:"sale" json:"sale" form:"sale"`
	Type        string `gorm:"column:type" db:"type" json:"type" form:"type"`
	Img         string `gorm:"column:img" db:"img" json:"img" form:"img"`
	// 在售 | 停售
	Status string `gorm:"column:status" db:"status" json:"status" form:"status"`
}

type ListOrder struct {
	Id          uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Userid      uint   `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`
	Username    string `gorm:"column:username" db:"username" json:"username" form:"username"`
	Productid   uint   `gorm:"column:productid" db:"productid" json:"productid" form:"productid"`
	Productname string `gorm:"column:productname" db:"productname" json:"productname" form:"productname"`
	Price       uint   `gorm:"column:price" db:"price" json:"price" form:"price"`
	Number      uint   `gorm:"column:number" db:"number" json:"number" form:"number"`
	Date        string `gorm:"column:date" db:"date" json:"date" form:"date"`
	Location    string `gorm:"column:location" db:"location" json:"location" form:"location"`
	Status      string `gorm:"column:status" db:"status" json:"status" form:"status"`
}

type ListAnnouncement struct {
	Id      uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Title   string `gorm:"column:title" db:"title" json:"title" form:"title"`
	Content string `gorm:"column:content" db:"content" json:"content" form:"content"`
	Date    string `gorm:"column:date" db:"date" json:"date" form:"date"`
	Img     string `gorm:"column:img" db:"img" json:"img" form:"img"`
}

type AnnouncementDetail struct {
	Id      uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Title   string `gorm:"column:title" db:"title" json:"title" form:"title"`
	Content string `gorm:"column:content" db:"content" json:"content" form:"content"`
	Date    string `gorm:"column:date" db:"date" json:"date" form:"date"`
	Img     string `gorm:"column:img" db:"img" json:"img" form:"img"`
}

type ListPost struct {
	Id           uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Userid       uint   `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`
	Username     string `gorm:"column:username" db:"username" json:"username" form:"username"`
	Title        string `gorm:"column:title" db:"title" json:"title" form:"title"`
	Content      string `gorm:"column:content" db:"content" json:"content" form:"content"`
	Browsenumber uint   `gorm:"column:browseNumber" db:"browseNumber" json:"browseNumber" form:"browseNumber"`
	Date         string `gorm:"column:date" db:"date" json:"date" form:"date"`
	Legal        int8   `gorm:"column:legal" db:"legal" json:"legal" form:"legal"`
	Elite        int8   `gorm:"column:elite" db:"elite" json:"elite" form:"elite"`
	Img          string `gorm:"column:img" db:"img" json:"img" form:"img"`
	Tag          string `gorm:"column:tag" db:"tag" json:"tag" form:"tag"`
}

type ListComment struct {
	Id       uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Userid   uint   `gorm:"column:userid" db:"userid" json:"userid" form:"userid"`
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"`
	Postid   uint   `gorm:"column:postid" db:"postid" json:"postid" form:"postid"`
	Content  string `gorm:"column:content" db:"content" json:"content" form:"content"`
	Date     string `gorm:"column:date" db:"date" json:"date" form:"date"`
	Support  uint   `gorm:"column:support" db:"support" json:"support" form:"support"`
	Against  uint   `gorm:"column:against" db:"against" json:"against" form:"against"`
}