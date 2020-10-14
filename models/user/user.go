package user

/*User - struct to load data into from USer table in DB*/
type User struct {
	UserID int
	Username string
	IP string
	UserAgent string
	LastLogin int64
	password string
}

/*NewUser - create and return a ference to a user struct with the provided attributes set*/
func NewUser(username string, password string, ip string, useragent string, login int64) *User {
	user := new(User)
	user.Username = username

	/*We don't need to use the SetPassword function, 
		but in case any more logic works it's way into the function we should*/
	user.SetPassword(password) 

	user.IP = ip
	user.UserAgent = useragent
	user.LastLogin = login
	return user
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