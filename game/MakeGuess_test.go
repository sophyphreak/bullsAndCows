package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeGuess(t *testing.T) {
	game1 := game{
		0,
		[4]int{1, 2, 3, 4},
		"IN PROGRESS",
		make([]Round, 0),
	}
	Games = append(Games, &game1)
	r := MakeGuess(game1.id, "1234")
	if r.ExactMatches != 4 {
		t.Errorf("Expected 4 exact matches but instead received %d", r.ExactMatches)
	}
	if r.PartialMatches != 0 {
		t.Errorf("Expected 0 partial matches but instead received %d", r.PartialMatches)
	}
	if r.Guess != 1 {
		t.Errorf("Expected Guess to be 1 but instead received %d", r.Guess)
	}
	if fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)) != "time.Time" {
		t.Errorf("Expected type of TimeOfGuess to be time.Time but instead received %v", fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)))
	}
	if game1.status != "FINISHED" {
		t.Errorf("Expected game1 status to be FINISHED but instead received %s", game1.status)
	}
	if len(game1.rounds) != 1 {
		t.Errorf("Expected game1 to have 1 round but instead received %d", len(game1.rounds))
	}

	game2 := game{
		0,
		[4]int{1, 2, 3, 4},
		"IN PROGRESS",
		make([]Round, 0),
	}
	Games = append(Games, &game2)
	r = MakeGuess(game2.id, "4321")
	if r.ExactMatches != 0 {
		t.Errorf("Expected 0 exact matches but instead received %d", r.ExactMatches)
	}
	if r.PartialMatches != 4 {
		t.Errorf("Expected 4 partial matches but instead received %d", r.PartialMatches)
	}
	if r.Guess != 1 {
		t.Errorf("Expected Guess to be 1 but instead received %d", r.Guess)
	}
	if fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)) != "time.Time" {
		t.Errorf("Expected type of TimeOfGuess to be time.Time but instead received %v", fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)))
	}
	if game2.status != "IN PROGRESS" {
		t.Errorf("Expected game1 status to be IN PROGRESS but instead received %s", game1.status)
	}
	if len(game2.rounds) != 1 {
		t.Errorf("Expected game1 to have 1 round but instead received %d", len(game1.rounds))
	}

	r = MakeGuess(game2.id, "5678")
	if r.ExactMatches != 0 {
		t.Errorf("Expected 0 exact matches but instead received %d", r.ExactMatches)
	}
	if r.PartialMatches != 0 {
		t.Errorf("Expected 0 partial matches but instead received %d", r.PartialMatches)
	}
	if r.Guess != 2 {
		t.Errorf("Expected Guess to be 1 but instead received %d", r.Guess)
	}
	if fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)) != "time.Time" {
		t.Errorf("Expected type of TimeOfGuess to be time.Time but instead received %v", fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)))
	}
	if game2.status != "IN PROGRESS" {
		t.Errorf("Expected game1 status to be IN PROGRESS but instead received %s", game1.status)
	}
	if len(game2.rounds) != 2 {
		t.Errorf("Expected game1 to have 2 round but instead received %d", len(game1.rounds))
	}

	r = MakeGuess(game2.id, "1243")
	if r.ExactMatches != 2 {
		t.Errorf("Expected 2 exact matches but instead received %d", r.ExactMatches)
	}
	if r.PartialMatches != 2 {
		t.Errorf("Expected 2 partial matches but instead received %d", r.PartialMatches)
	}
	if r.Guess != 3 {
		t.Errorf("Expected Guess to be 3 but instead received %d", r.Guess)
	}
	if fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)) != "time.Time" {
		t.Errorf("Expected type of TimeOfGuess to be time.Time but instead received %v", fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)))
	}
	if game2.status != "IN PROGRESS" {
		t.Errorf("Expected game1 status to be IN PROGRESS but instead received %s", game1.status)
	}
	if len(game2.rounds) != 3 {
		t.Errorf("Expected game1 to have 3 round but instead received %d", len(game1.rounds))
	}

	r = MakeGuess(game2.id, "1234")
	if r.ExactMatches != 4 {
		t.Errorf("Expected 4 exact matches but instead received %d", r.ExactMatches)
	}
	if r.PartialMatches != 0 {
		t.Errorf("Expected 0 partial matches but instead received %d", r.PartialMatches)
	}
	if r.Guess != 4 {
		t.Errorf("Expected Guess to be 4 but instead received %d", r.Guess)
	}
	if fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)) != "time.Time" {
		t.Errorf("Expected type of TimeOfGuess to be time.Time but instead received %v", fmt.Sprint(reflect.TypeOf(r.TimeOfGuess)))
	}
	if game2.status != "FINISHED" {
		t.Errorf("Expected game1 status to be FINISHED but instead received %s", game1.status)
	}
	if len(game2.rounds) != 4 {
		t.Errorf("Expected game1 to have 4 round but instead received %d", len(game1.rounds))
	}
}
