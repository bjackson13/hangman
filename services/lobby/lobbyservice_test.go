package lobby

import (
	"reflect"
	"testing"

	"github.com/bjackson13/hangman/models/user"
)

func TestService_addUser(t *testing.T) {

	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		service *Service
		args    args
		wantErr bool
	}{
		{
			name:    "add postman",
			args:    args{3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{}
			if err := service.addUser(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Service.addUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_getLobbyUsers(t *testing.T) {
	testUser := user.NewUser("postman", "", "", "", 0)
	testUser.UserID = 3
	tests := []struct {
		name    string
		service *Service
		want    []user.User
		wantErr bool
	}{
		{
			name:    "get postman",
			want:    []user.User{*testUser},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{}
			got, err := service.getLobbyUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.getLobbyUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.getLobbyUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_userIsInLobby(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		service *Service
		args    args
		want    bool
	}{
		{
			name: "Postman is in lobby",
			args: args{3},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{}
			if got := service.userIsInLobby(tt.args.userID); got != tt.want {
				t.Errorf("Service.userIsInLobby() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_removeUser(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		service *Service
		args    args
		wantErr bool
	}{
		{
			name: "Postman is in lobby",
			args: args{3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{}
			if err := service.removeUser(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Service.removeUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
