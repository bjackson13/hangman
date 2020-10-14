package authentication

import (
	"testing"
	"time"
)

func TestAuthenticateUserLogin(t *testing.T) {
	//test user in test DB
	username := "auth"
	password := "auth"

	user, err := AuthenticateUserLogin(username, password)
	if err != nil {
		t.Errorf("Failed to authenticate user: %s", err.Error())
	}
	
	if user == nil || user.Username != username || user.IP != "192.1.1.1" || user.UserAgent != "chrome" {
		t.Errorf("Failed to retireve proper user details:\n %s\n%s\n%s\n", user.Username, user.IP, user.UserAgent)
	}
}

func TestInvalidUserLoginNoUserNoPass(t *testing.T) {
	username := ""
	password := ""

	user, err := AuthenticateUserLogin(username, password)
	if err == nil {
		t.Errorf("Should have received error while logging in")
	}
	
	if user != nil {
		t.Errorf("Failed to return error, got user: %v", user)
	}

}

func TestInvalidUserLoginValidUserInvlaidPass(t *testing.T) {
	username := "auth"
	password := "thisisinvalid"

	user, err := AuthenticateUserLogin(username, password)
	if err == nil {
		t.Errorf("Should have received error while logging in")
	}
	
	if user != nil {
		t.Errorf("Failed to return error, got user: %v", user)
	}

}

func TestGenerateVerifyParseSessionToken(t *testing.T) {
	timestamp := time.Now() //mock timestamp
	username := "auth"
	ip := "192.168.1.1"
	useragent :=  "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0"

	token := GenerateSessionToken(username, timestamp, ip, useragent)
	if token == "" {
		t.Errorf("invalid token generated")
	}

	valid := VerifySessionToken(token)
	if !valid {
		t.Errorf("could not verify token")
	}

	parsedtoken := ParseSessionToken(token)
	if parsedtoken == nil {
		t.Errorf("invalid token could not be parsed")
	}

	parsedName,_ := parsedtoken.GetStr("username")
	parsedIP,_  := parsedtoken.GetStr("ip")
	parsedUA,_  := parsedtoken.GetStr("useragent")

	if parsedName != username || parsedIP != ip || parsedUA != useragent || parsedtoken.Validate() != nil {
		t.Errorf("invalid token parsed")
	}
}