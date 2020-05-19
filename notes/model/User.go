package model

type Register struct {
	Username string
	Password string
}

type User struct {
	Id           int
	Username     string
	Password     string
	Gender       string
	Telphone     string
	Thumnail     string
	Email        string
	Introduction string
	Status       string
	Create_at    string
	Update_at    string
	Role         int
}

func (User) TableName() string{
	return "user"
}