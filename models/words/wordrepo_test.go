package words

import (
	"testing"
	"os"
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
	id, err := wordRepo.AddWord("test")
	if err != nil {
		t.Errorf("Error inserting new word into DB, %s", err.Error())
	} else if id < 0 {
		t.Errorf("Error inserting new word into DB, %s", err.Error())
	}
	wordID = id
}

func TestUpdateWordGuesses(t *testing.T){
	word := NewGameWord("test", "ts", "qr")
	word.WordID = wordID
	err := wordRepo.UpdateWordGuesses(*word)
	if err != nil {
		t.Errorf("Error updating word guesses into DB, %s", err.Error())
	}
}

func TestGetWord(t *testing.T){
	word,err := wordRepo.GetWord(wordID)
	if err != nil {
		t.Errorf("Error getting word from DB, %s", err.Error())
	}

	if word.Word != "test" || word.GetCorrectGuesses() != "ts" || word.GetIncorrectGuesses() != "qr" {
		t.Errorf("Correct values not retireved from DB, %s", err.Error())
	} 
}

func TestUpdateword(t *testing.T){
	err := wordRepo.UpdateWord("DogsRKewl", wordID)
	if err != nil {
		t.Errorf("Error updating word guesses into DB, %s", err.Error())
	}

	word,err := wordRepo.GetWord(wordID)
	if err != nil {
		t.Errorf("Error getting updated word from DB, %s", err.Error())
	}

	if word.Word != "DogsRKewl" || word.GetCorrectGuesses() != "" || word.GetIncorrectGuesses() != "" {
		t.Errorf("Correct values not retireved from DB, %s", err.Error())
	} 
}

func TestDeleteWord(t *testing.T) {
	err := wordRepo.DeleteWord(wordID)
	if err != nil {
		t.Errorf("Error inserting new word into DB, %s", err.Error())
	}
}