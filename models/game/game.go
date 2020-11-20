package games

/*Game represents a game*/
type Game struct {
	GameID int
	WordID int
	GuessingUserID int
	WordCreatorID int
	PendingGuess string
}