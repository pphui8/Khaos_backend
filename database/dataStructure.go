package database

type UserDetail struct {
	Id           int
	Username     string
	Descript     string
	Registerdate string
	Phone        string
	Privilege    string
	Publickey    string
}

type ListUserData struct {
	Id           int
	Username     string
	Registerdate string
	Phone        string
	Privilege    string
}