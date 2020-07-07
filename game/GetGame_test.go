package game

import (
	"reflect"
	"testing"
)

func TestGetGame(t *testing.T) {
	Begin()
	got, err := GetGame(0)
	expected := Game{0, [4]int{}, "IN PROGRESS", make([]Round, 1)}

	if err != nil {
		t.Errorf("Got %v erro but expected nil", err)
	}

	if got.ID != expected.ID {
		t.Errorf("Got id %v Expected id %v", got.ID, expected.ID)
	}
	if got.Answer != expected.Answer {
		t.Errorf("Got answer %v but Expected %v", got.Answer, expected.Answer)
	}
	if reflect.DeepEqual(got.Rounds, expected.Rounds) == false {
		t.Errorf("Got rounds %v Expected %v", got.Rounds, expected.Rounds)
	}
	if got.Status != expected.Status {
		t.Errorf("Got status %v Expected %v", got.Status, expected.Status)
	}
	Games = Games[:0]
}
