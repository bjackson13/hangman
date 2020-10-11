package user

/*User - struct to load data into from USer table in DB*/
type User struct {
	UserID int
	Username string
	IP string
	UserAgent string
	password string
}

/*
	Password getters and setters are used to allow the password field 
		to remain unexported so it is not manipulated when not needed
*/

/*SetPassword - set the user password*/
func (user *User) SetPassword(password string) {
	user.password = password
}

/*GetPassword - get the user password*/
func (user *User) GetPassword() string {
	return user.password
}