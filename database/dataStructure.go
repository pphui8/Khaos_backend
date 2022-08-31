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