package words

/*GameWord - represents the word our game is centered around*/
type GameWord struct {
	WordID int
	Word string
	correctGuesses []byte
	incorrectGuesses []byte
}


/*NewGameWord create new game word*/
func NewGameWord(wordText string, correctGuesses string, incorrectGuesses string) *GameWord{
	word := new(GameWord)
	word.Word = wordText
	word.correctGuesses = append(word.correctGuesses, correctGuesses...)
	word.incorrectGuesses = append(word.incorrectGuesses, incorrectGuesses...)
	return word
}

/*AddGuess add a guess for our word, true for correct, false if not*/
func (word *GameWord) AddGuess(isCorrectGuess bool, letter string) {
	/*Prevent more than a letter from being written at once*/
	if isCorrectGuess {
		word.correctGuesses = append(word.correctGuesses, letter[0])
	} else {
		word.incorrectGuesses = append(word.incorrectGuesses, letter[0])
	}
}

/*GetCorrectGuesses convert byte array to string for correct guesses*/
func (word *GameWord) GetCorrectGuesses() string {
	return string(word.correctGuesses[:])
}

/*GetIncorrectGuesses convert byte array to string for incorrect guesses*/
func (word *GameWord) GetIncorrectGuesses() string {
	return string(word.incorrectGuesses[:])
}
