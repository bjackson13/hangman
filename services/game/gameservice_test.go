package games

import (
	"reflect"
	"testing"

	"github.com/bjackson13/hangman/models/game"
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
