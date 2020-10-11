package user

/*User - struct to load data into from USer table in DB*/
type User struct {
	UserID int
	Username string
	IP string
	UserAgent string
}