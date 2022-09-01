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
	// 在售 | 停售
	Status string `gorm:"column:status" db:"status" json:"status" form:"status"`
}