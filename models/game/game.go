package games

type Game struct {
	GameID int
	WordID int
	GuessingUserID int
	WordCreatorID int
	ChatID int
	PendingGuess *string
}