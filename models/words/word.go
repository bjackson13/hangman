package words

import "strings"

/*GameWord - represents the word our game is centered around*/
type GameWord struct {
	WordID           int
	Length           int
	correctGuesses   []string
	incorrectGuesses []string
}

/*NewGameWord create new game word*/
func NewGameWord(wordLength int) *GameWord {
	word := new(GameWord)
	word.Length = wordLength
	word.correctGuesses = make([]string, word.Length)
	return word
}

/*AddCorrectGuess add a correct guess to our word. Takes the letter to be added and the indexes to set as correct*/
func (word *GameWord) AddCorrectGuess(letter string, indexes []int) {
	/*Prevent more than a letter from being written at once by using [0]*/
	char := string(letter[0])
	for _,v := range indexes {
		word.correctGuesses[v] = char
	}
		
}

/*AddIncorrectGuess add an incorrect guess*/
func (word *GameWord) AddIncorrectGuess(letter string) {
	/*Prevent more than a letter from being written at once by using [0]*/
	word.incorrectGuesses = append(word.incorrectGuesses, string(letter[0]))
		
}

/*GetCorrectGuesses return correct guesses in CSV format*/
func (word *GameWord) GetCorrectGuesses() string {
	return strings.Join(word.correctGuesses, ",")
}

/*GetIncorrectGuesses return incorrect guesses in CSV format*/
func (word *GameWord) GetIncorrectGuesses() string {
	return strings.Join(word.incorrectGuesses, ",")
}

/*SetCorrectGuesses set correct guesses*/
func (word *GameWord) SetCorrectGuesses(guesses string) {
	word.correctGuesses = []string{guesses}
}

/*SetIncorrectGuesses set incorrect guesses*/
func (word *GameWord) SetIncorrectGuesses(guesses string) {
	word.incorrectGuesses = []string{guesses}
}

/*isCompleted returns if the word has been completely guessed or not*/
func (word *GameWord) isCompleted() bool {
	for _, v := range word.correctGuesses {
		if v == "" {
			return false
		}
	} 
	return true
}
