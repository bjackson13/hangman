package authentication

import (
	"github.com/bjackson13/hangman/models/user"
	"golang.org/x/crypto/bcrypt"
	tokenGen "github.com/brianvoe/sjwt"
	"time"
	"errors"
)

var SUPER_DUPER_SECRET_KEY []byte = []byte("OMG_50_5p00ky") //this will be moved... 

/*AuthenticateUserLogin - authenticates a users credentials and returns the user or an error*/
func AuthenticateUserLogin(username string, password string, requestIP string, requestUserAgent string) (*user.User, error) {
	userRepo, _ := user.NewRepo()

	user, err := userRepo.GetUser(username)
	if err != nil {
		return nil, err
	}
	
	validPwd := compareHashToString(user.GetPassword(), password)
	if !validPwd {
		return nil, errors.New("Invalid login")
	}
	
	// Update our user identifiers in a go routine to save on time responding to the user request
	go func() {
		userRepo.UpdateUserIdentifiers(user.UserID, requestIP, requestUserAgent, time.Now().Unix())
		userRepo.Close()
	}()
	
	return user, nil
	
}

/*GenerateSessionToken use user identifying info to generate a session token*/
func GenerateSessionToken(validUser user.User) string {
	token := tokenGen.New()

	token.Set("id", validUser.UserID)
	token.Set("username", validUser.Username)
	token.Set("ip", validUser.IP)
	token.Set("useragent", validUser.UserAgent)
	token.Set("lastlogin", validUser.LastLogin)

	//set expiration
	token.SetExpiresAt(time.Unix(validUser.LastLogin, 0).Add(time.Hour * 24))

	finaltoken := token.Generate(SUPER_DUPER_SECRET_KEY)

	return finaltoken
}

/*VerifyAndProcessToken for a given token verify it is untampered, unexpired, 
	and finally get the token data and compare it to request data. 
	Return User if valid or error if not*/
func VerifyAndProcessToken(token string, requestIP string, requestUA string) (*user.User, error) {
	if !verifySessionToken(token) { return nil, errors.New("Invalid Session Token") }
	if !checkSessionExpired(token) { return nil, errors.New("Expired Session Token") }
	parsedUser, err := parseSessionToken(token)
	if err != nil { return nil, err }
	if parsedUser.IP != requestIP || parsedUser.UserAgent != requestUA { return nil, errors.New("Bad Request") }
	return parsedUser, nil
}

/*VerifySessionToken verifies the given session token is valid and returns values*/
func verifySessionToken(token string) bool {
	return tokenGen.Verify(token, SUPER_DUPER_SECRET_KEY)
}

/*CheckSessionExpired Verify if a token is expired or not*/
/*True if expired.
	False if not expired*/
func checkSessionExpired(token string) bool {
	claim, err := tokenGen.Parse(token)
	// if we can't parse the token assume expired
	if err != nil {
		return true
	}
	//make sure session is stil valid
	return claim.Validate() == nil
}

/*ParseSessionToken parse session token, supress warning*/
func parseSessionToken(token string) (*user.User, error) {
	claim, err := tokenGen.Parse(token)
	if err != nil {
		return nil, err
	}

	username,_ := claim.GetStr("username")
	ip,_ := claim.GetStr("ip")
	useragent,_ := claim.GetStr("useragent")
	lastlogin,_ := claim.GetInt("lastlogin")

	parsedUser := user.NewUser(username, "", ip, useragent, int64(lastlogin))
	parsedUser.UserID,_ = claim.GetInt("id")

	return parsedUser, err
}

/*Internal function to comparing 2 strings hash values*/
func compareHashToString(hashedTxt string, plainTxt string) bool {   
	byteHash := []byte(hashedTxt)    
	byteTxt := []byte(plainTxt)
	err := bcrypt.CompareHashAndPassword(byteHash, byteTxt)
    if err != nil {
        return false
    }
    
	return true
}
