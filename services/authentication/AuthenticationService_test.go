package authentication

import (
	"testing"
	"time"
	"github.com/bjackson13/hangman/models/user"
)

func TestAuthenticateUserLogin(t *testing.T) {
	//test user in test DB
	username := "bren"
	password := "bren"

	user, err := AuthenticateUserLogin(username, password, "192.1.1.1", "chrome")
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

	user, err := AuthenticateUserLogin(username, password, "", "")
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

	user, err := AuthenticateUserLogin(username, password, "", "")
	if err == nil {
		t.Errorf("Should have received error while logging in")
	}
	
	if user != nil {
		t.Errorf("Failed to return error, got user: %v", user)
	}

}

func TestGenerateVerifyparseSessionToken(t *testing.T) {
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", time.Now().Unix())
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) //dereference user, don't want values to be changed
	if token == "" {
		t.Errorf("invalid token generated")
	}

	valid := verifySessionToken(token)
	if !valid {
		t.Errorf("could not verify token")
	}

	parsedUser, err := parseSessionToken(token)
	if parsedUser == nil || err != nil {
		t.Errorf("invalid token could not be parsed: %v", err.Error())
	}

	if parsedUser.Username != testUser.Username || parsedUser.IP != testUser.IP || parsedUser.UserAgent != testUser.UserAgent || parsedUser.UserID != testUser.UserID {
		t.Errorf("invalid token parsed")
	}
}

func TestParseBadTokenShouldFail(t *testing.T) {
	token := "blah"
	testUser, err := parseSessionToken(token)
	if err == nil || testUser != nil {
		t.Errorf("Error should have been returned for bad session")
	}
}

func TestTamperedTokenFailsToVerify(t *testing.T) {
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", time.Now().Unix())
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) 

	//tamper with token string
	token = token + "blahblah"

	valid := verifySessionToken(token)
	if valid {
		t.Errorf("Token should not be verified")
	}
}

func TestcheckSessionExpiredForUnexpiredToken(t *testing.T) {
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", time.Now().Unix())
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) 

	valid := checkSessionExpired(token)
	if !valid {
		t.Errorf("Token should be unexpired")
	}
}

func TestcheckSessionExpiredForExpiredToken(t *testing.T) {
	timestamp := time.Now().Add(time.Hour * 26 * -1).Unix()// timestamp is over 24 hours ago
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", timestamp)
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) 

	valid := checkSessionExpired(token)
	if valid {
		t.Errorf("Token should be expired")
	}
}

func TestVerifyAndProcessTokenForGoodToken(t *testing.T) {
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", time.Now().Unix())
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) 

	reqIP := "192.168.1.1"
	reqUA := "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)"
	verifiedUser, err := VerifyAndProcessToken(token, reqIP, reqUA)

	if err != nil {
		t.Log(err)
		t.Error("Unable to parse user from token")
	}

	if testUser.Username != verifiedUser.Username || testUser.UserID != verifiedUser.UserID {
		t.Error("Incorrect values returned")
	}
}

func TestVerifyAndProcessTokenForBadToken(t *testing.T) {
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", time.Now().Unix())
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) 
	token = token + "blahblah"

	_, err := VerifyAndProcessToken(token, "", "")
	if err == nil {
		t.Error("Should have returned error when verifying token")
	}
}

func TestVerifyAndProcessTokenForBadRequestParams(t *testing.T) {
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", time.Now().Unix())
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) 

	reqIP := "129.21.12.12"
	reqUA := "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)"
	_, err := VerifyAndProcessToken(token, reqIP, reqUA)

	if err == nil {
		t.Error("Should have returned error when verifying request params")
	}
}

func TestVerifyAndProcessTokenForExpiredToken(t *testing.T) {
	timestamp := time.Now().Add(time.Hour * 26 * -1).Unix()// timestamp is over 24 hours ago
	testUser := user.NewUser("auth", "", "192.168.1.1", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0)", timestamp)
	testUser.UserID = 101

	token := GenerateSessionToken(*testUser) 

	_, err := VerifyAndProcessToken(token, "", "")
	if err == nil {
		t.Error("Should have returned error when verifying token expiration")
	}
}