package authentication

import (
	"github.com/bjackson13/hangman/models"
	"github.com/bjackson13/hangman/models/user"
	"golang.org/x/crypto/bcrypt"
	tokenGen "github.com/brianvoe/sjwt"
	"time"
	"errors"
)

var SUPER_DUPER_SECRET_KEY []byte = []byte("OMG_50_5p00ky") //this will be moved... 

/*AuthenticateUserLogin - authenticates a users credentials and returns the user or an error*/
func AuthenticateUserLogin(username string, password string) (*user.User, error) {
	dbConn, err := dbconn.Connect() 
	if err != nil {
		return nil, err
	}
	defer dbConn.Connection.Close()
	
	userRepo := user.NewRepo(dbConn)
	user, err := userRepo.GetUser(username) 

	validPwd := compareHashToString(user.GetPassword(), password)
	if validPwd {
		return user, nil
	}
	
	return nil, errors.New("Invalid login")
}

/*GenerateSessionToken use user identifying info to generate a session token*/
func GenerateSessionToken(username string, timestamp time.Time, ip string, useragent string) string {
	
	token := tokenGen.New()
	token.Set("username", username)
	token.Set("ip", ip)
	token.Set("useragent", useragent)

	//set expiration
	token.SetExpiresAt(timestamp.Add(time.Hour * 24))

	finaltoken := token.Generate(SUPER_DUPER_SECRET_KEY)

	return finaltoken
}

/*VerifySessionToken verifies the given session token is valid and returns values*/
func VerifySessionToken(token string) bool {
	return tokenGen.Verify(token, SUPER_DUPER_SECRET_KEY)
}

/*ParseSessionToken parse session token, supress warning*/
func ParseSessionToken(token string) tokenGen.Claims {
	claim, _ := tokenGen.Parse(token)
	return claim
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
