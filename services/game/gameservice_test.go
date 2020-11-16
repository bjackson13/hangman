package games

import (
	"reflect"
	"testing"

	games "github.com/bjackson13/hangman/models/game"
)

func TestService_GetUserGame(t *testing.T) {
	//game := &games.Game{33,-1,2,3,""} //uncomment and change game value if needing to test again
	type args struct {
		userID int
	}
	tests := []struct {
		name string
		s    *Service
		args args
		want *games.Game
	}{
		{
			name: "Postman has no games",
			args: args{3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if got := s.GetUserGame(tt.args.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetUserGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_MakeGuess(t *testing.T) {
	type args struct {
		gameID int
		userID int
		guess  string
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		{
			name: "Make a guess for valid guesser",
			args: args{1,1,"g"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if err := s.MakeGuess(tt.args.gameID, tt.args.userID, tt.args.guess); (err != nil) != tt.wantErr {
				t.Errorf("Service.MakeGuess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_EndGame(t *testing.T) {
	type args struct {
		gameID int
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		{
			name:    "Remove nonexistent game",
			args:    args{1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if err := s.EndGame(tt.args.gameID); (err != nil) != tt.wantErr {
				t.Errorf("Service.EndGame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
