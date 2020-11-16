package lobby

import (
	"github.com/bjackson13/hangman/models"
	"github.com/bjackson13/hangman/models/user"
	"sync"
	"database/sql"
)

/*Repo - Struct for CRUDing users from the database*/
type Repo struct {
	dbconn.Repo
}

/*NewRepo - Create new repo with access to mysql database*/
func NewRepo() (*Repo, error) {
	conn, err := dbconn.Connect() 
	if err != nil {
		return nil, err
	}

	repo := new(Repo)
	repo.DB = conn
	return repo, nil
}

/*AddLobbyUser add a user to the lobby bu UserID*/
func (repo *Repo) AddLobbyUser(userID int) error {
	lobbyStmt, err := repo.DB.Prepare("INSERT INTO LobbyUsers(UserId) VALUE (?)")
	defer lobbyStmt.Close()
	if err != nil {
		return err
	}

	_, err = lobbyStmt.Exec(userID)
	return err
}

/*GetAllLobbyUsers get all users in the lobby. Returns list of User structs with ID and USername filled out*/
func (repo *Repo) GetAllLobbyUsers() ([]user.User, error) {
	lobbyStmt, err := repo.DB.Prepare("SELECT User.UserId, User.Username FROM LobbyUsers INNER JOIN User ON LobbyUsers.UserId = User.UserId")
	defer lobbyStmt.Close()
	if err != nil {
		return nil, err
	}
	
	users := []user.User{}
	rows, err := lobbyStmt.Query()

	var wg sync.WaitGroup
	for rows.Next() {
		var u user.User
		err = rows.Scan(&u.UserID, &u.Username)
		if err != nil {
			break;
		}

		wg.Add(1)
		go func () {
			users = append (users, u)
			wg.Done()
		}()
	}

	wg.Wait()
	return users, err
}

/*UserIsInLobby check if a particular user id is in the lobby*/
func (repo *Repo) UserIsInLobby(userID int) (bool, error) {
	lobbyStmt, err := repo.DB.Prepare("SELECT UserId FROM LobbyUsers WHERE UserId = ?")
	defer lobbyStmt.Close()
	if err != nil {
		return false, err
	}
	
	/*An error is thrown if we try to scan an empty result set, 
		use this to determine if user is in lobby*/
	var user int
	err = lobbyStmt.QueryRow(userID).Scan(&user)
	return err == nil, err
}

/*RemoveLobbyUser remove user from the lobby*/
func (repo *Repo) RemoveLobbyUser(userID int) error {
	lobbyStmt, err := repo.DB.Prepare("DELETE FROM LobbyUsers WHERE UserId = ?")
	defer lobbyStmt.Close()
	if err != nil {
		return err
	}

	_, err = lobbyStmt.Exec(userID)
	return err
}

/*InviteUser invite user to play a game*/
func (repo *Repo) InviteUser(inviteeID int, inviterID int) error {
	lobbyStmt, err := repo.DB.Prepare("UPDATE LobbyUsers SET PendingInviteId = ? WHERE UserId = ?")
	defer lobbyStmt.Close()
	if err != nil {
		return err
	}

	_, err = lobbyStmt.Exec(inviterID, inviteeID)
	return err
}

/*RevokeInvite revoke invite to play. Set Pending invite user ID to NULL*/
func (repo *Repo) RevokeInvite(inviteeID int) error {
	lobbyStmt, err := repo.DB.Prepare("UPDATE LobbyUsers SET PendingInviteId = NULL WHERE UserId = ?")
	defer lobbyStmt.Close()
	if err != nil {
		return err
	}

	_, err = lobbyStmt.Exec(inviteeID)
	return err
}

/*CheckInvites check if a particular user id has any invites. Will return user ID who invited user or -1 for no invites*/
func (repo *Repo) CheckInvites(inviteeID int) (*string, int, error) {
	lobbyStmt, err := repo.DB.Prepare("SELECT User.Username, LobbyUsers.PendingInviteId FROM LobbyUsers JOIN User ON LobbyUsers.PendingInviteId = User.UserID WHERE LobbyUsers.UserId = ?")
	defer lobbyStmt.Close()
	if err != nil {
		return nil, -1, err
	}

	var pendingInviteUsername sql.NullString
	var pendingInviteUserID sql.NullInt32
	err = lobbyStmt.QueryRow(inviteeID).Scan(&pendingInviteUsername, &pendingInviteUserID)
	if !pendingInviteUsername.Valid || !pendingInviteUserID.Valid {
		return nil, -1, nil
	}
	
	inviterName := pendingInviteUsername.String
	inviterID := int(pendingInviteUserID.Int32)
	return &inviterName, inviterID, err
}
