package game

import (
	"reflect"
	"testing"
)

func TestListGames(t *testing.T) {

	game1 := Game{
		0,
		[4]int{1, 2, 3, 4},
		"IN PROGRESS",
		make([]Round, 1),
	}
	Games = append(Games, &game1)

	checkGame := ListGames()

	if game1.ID != checkGame[0].ID {
		t.Errorf("Got id %v chekGame[0] id %v", game1.ID, checkGame[0].ID)
	}
	if game1.Answer != checkGame[0].Answer {
		t.Errorf("Got answer %v but chekGame[0] %v", game1.Answer, checkGame[0].Answer)
	}
	if reflect.DeepEqual(game1.Rounds, checkGame[0].Rounds) == false {
		t.Errorf("Got rounds %v chekGame[0] %v", game1.Rounds, checkGame[0].Rounds)
	}
	if game1.Status != checkGame[0].Status {
		t.Errorf("Got status %v chekGame[0] %v", game1.Status, checkGame[0].Status)
	}
}
