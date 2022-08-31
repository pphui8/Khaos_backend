package database

type UserDetail struct {
	Id           int
	Username     string
	Descript     string
	Registerdate string
	Phone        string
	Privilege    string
}

type ListUserData struct {
	Id           int
	Username     string
	Registerdate string
	Phone        string
	Privilege    string
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
	Stock       uint   `gorm:"column:stock" db:"stock" json:"stock" form:"stock"`
	Sale        uint   `gorm:"column:sale" db:"sale" json:"sale" form:"sale"`
	Type        string `gorm:"column:type" db:"type" json:"type" form:"type"`
	Status      string `gorm:"column:status" db:"status" json:"status" form:"status"`
}

type ProductDetail struct {
	Id          uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Productname string `gorm:"column:productname" db:"productname" json:"productname" form:"productname"`
	Descript    string `gorm:"column:descript" db:"descript" json:"descript" form:"descript"`
	Price       uint   `gorm:"column:price" db:"price" json:"price" form:"price"`
	Stock       uint   `gorm:"column:stock" db:"stock" json:"stock" form:"stock"`
	Sale        uint   `gorm:"column:sale" db:"sale" json:"sale" form:"sale"`
	Img         string `gorm:"column:img" db:"img" json:"img" form:"img"`
	Type        string `gorm:"column:type" db:"type" json:"type" form:"type"`
	Status      string `gorm:"column:status" db:"status" json:"status" form:"status"`
}