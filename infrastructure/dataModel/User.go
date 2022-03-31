package datamodel

type UserType int

type UserModel struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	ContactNo string `db:"contact_no"`
}
