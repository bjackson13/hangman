package authentication

import (
	"testing"
	"time"
	"github.com/bjackson13/hangman/models/user"
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
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", time.Now().Unix())
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) //dereference user, don't want values to be changed
	if token == "" {
		t.Errorf("invalid token generated")
	}

	valid := VerifySessionToken(token)
	if !valid {
		t.Errorf("could not verify token")
	}

	parsedUser, expiredSession := ParseSessionToken(token)
	if parsedUser == nil {
		t.Errorf("invalid token could not be parsed")
	}

	if expiredSession != nil {
		t.Errorf("Session expired")
	}

	if parsedUser.Username != testUser.Username || parsedUser.IP != testUser.IP || parsedUser.UserAgent != testUser.UserAgent || parsedUser.UserID != testUser.UserID || expiredSession != nil {
		t.Errorf("invalid token parsed")
	}
}

func TestParseBadTokenShouldFail(t *testing.T) {
	token := "blah"
	testUser, badSession := ParseSessionToken(token)
	if badSession == nil || testUser != nil {
		t.Errorf("Error should have been returned for bad session")
	}
}

func TestTamperedTokenFailsToVerify(t *testing.T) {
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", time.Now().Unix())
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) 

	//tamper with token string
	token = token + "blahblah"

	valid := VerifySessionToken(token)
	if valid {
		t.Errorf("Token should not be verified")
	}
}