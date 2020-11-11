package words

import (
	"os"
	"testing"
)

var wordRepo *Repo
var wordID int

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err error
	wordRepo, err = NewRepo()
	if err != nil {
		panic(err.Error())
	}
}

func teardown() {
	err := wordRepo.Close()
	if err != nil {
		panic(err.Error())
	}
}

func TestAddWord(t *testing.T) {
	id, err := wordRepo.AddWord(4)
	if err != nil {
		t.Errorf("Error inserting new word into DB, %s", err.Error())
	} else if id < 0 {
		t.Errorf("Error inserting new word into DB")
	}
	wordID = id
}

func TestUpdateWordGuesses(t *testing.T) {
	word := GameWord{wordID,4,make([]string,1),make([]string,1)}
	word.SetCorrectGuesses("")
	word.SetIncorrectGuesses("")

	word.AddCorrectGuess("t", []int{0,3})
	word.AddIncorrectGuess("q")
	word.AddIncorrectGuess("r")

	err := wordRepo.UpdateWordGuesses(word)
	if err != nil {
		t.Errorf("Error updating word guesses into DB, %s", err.Error())
	}
}

func TestGetWord(t *testing.T) {
	word, err := wordRepo.GetWord(wordID)
	if err != nil {
		t.Errorf("Error getting word from DB, %s", err.Error())
	}

	if word.Length != 4 || word.GetCorrectGuesses() != "t,,,t" || word.GetIncorrectGuesses() != "q,r,," {
		t.Errorf("Correct values not retireved from DB, %v %s %s", word.Length, word.GetCorrectGuesses(), word.GetIncorrectGuesses())
	}
}

func TestUpdateword(t *testing.T) {
	err := wordRepo.UpdateWord(5, wordID)
	if err != nil {
		t.Errorf("Error updating word guesses into DB, %s", err.Error())
	}

	word, err := wordRepo.GetWord(wordID)

	if err != nil {
		t.Errorf("Error getting updated word from DB, %s", err.Error())
	}
	
	if word.Length != 5 || word.GetCorrectGuesses() != ",,,," || word.GetIncorrectGuesses() != ",,,," {
		t.Errorf("Correct values not retireved from DB: %v", word)
	}
}

func TestDeleteWord(t *testing.T) {
	err := wordRepo.DeleteWord(wordID)
	if err != nil {
		t.Errorf("Error inserting new word into DB, %s", err.Error())
	}
}

func TestIsCompleted(t *testing.T) {
	word := GameWord{wordID,4,make([]string,1),make([]string,1)}
	word.SetCorrectGuesses("")
	word.SetIncorrectGuesses("")

	word.AddCorrectGuess("t", []int{0,3})
	word.AddIncorrectGuess("q")
	word.AddIncorrectGuess("r")

	if word.IsCompleted() {
		t.Errorf("Word should not be completed: %s", word.GetCorrectGuesses())
	}

	t.Log(word.GuessLimitExceeded())
	word.AddCorrectGuess("e", []int{1})
	word.AddCorrectGuess("s", []int{2})

	if !word.IsCompleted() {
		t.Errorf("Word should not completed: %s", word.GetCorrectGuesses())
	}
}

func TestGuessAlreadyComplete(t *testing.T) {
	
	word := GameWord{wordID,4,make([]string,1),make([]string,1)}
	word.SetCorrectGuesses("")
	word.SetIncorrectGuesses("")

	word.AddCorrectGuess("t", []int{0,3})
	word.AddIncorrectGuess("q")
	word.AddIncorrectGuess("r")
	
	found1 := word.GuessAlreadyMade("t")
	found2 := word.GuessAlreadyMade("w")
	found3 := word.GuessAlreadyMade("r")

	if !found1 || found2 || !found3 {
		t.Errorf("Incorrect guess located: %q", word)
	}
}