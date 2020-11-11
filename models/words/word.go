package words

import (
	"strings"
)

/*GameWord - represents the word our game is centered around*/
type GameWord struct {
	WordID           int
	Length           int
	correctGuesses   []string
	incorrectGuesses []string
}

var maxIncorrectGuesses = 7

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
	for i := 0; i < len(word.incorrectGuesses); i++ {
		if word.incorrectGuesses[i] == "" {
			word.incorrectGuesses[i] = string(letter[0])
			break
		}
	}
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
	word.correctGuesses = make([]string, word.Length)
	for i,v := range strings.Split(guesses, ",") {
		word.correctGuesses[i] = v
	}
}

/*SetIncorrectGuesses set incorrect guesses*/
func (word *GameWord) SetIncorrectGuesses(guesses string) {
	if word.Length < 7 {
		word.incorrectGuesses = make([]string, word.Length, word.Length)
	} else {
		word.incorrectGuesses = make([]string, 7, 7)
	}
	for i,v := range strings.Split(guesses, ",") {
		word.incorrectGuesses[i] = v
	}
}

/*IsCompleted returns if the word has been completely guessed or not*/
func (word *GameWord) IsCompleted() bool {
	for _, v := range word.correctGuesses {
		if v == "" {
			return false
		}
	} 
	return true
}

/*GuessLimitExceeded check if user has submitted too many incorrect guesses*/
func (word *GameWord) GuessLimitExceeded() bool {
	for i := 0; i < len(word.incorrectGuesses); i++ {
		if word.incorrectGuesses[i] == "" {
			return false
		}
	} 
	return true
}

/*GuessAlreadyMade check if guess was already made*/
func (word *GameWord) GuessAlreadyMade(guess string) bool {
	check := make(chan bool, 2)
	go func() {
		found := false
		for _,v := range word.correctGuesses {
			if v == string(guess[0]) {
				found = true
				break
			}
		}
		check <- found
	}()

	go func() {
		found := false
		for _,v := range word.incorrectGuesses {
			if v == string(guess[0]) {
				found = true
				break
			}
		}
		check <- found
	}()
	
	found1 := <- check 
	found2 := <- check

	return found1 || found2
}