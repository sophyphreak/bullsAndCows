package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeGuess(t *testing.T) {
	game1 := Game{
		0,
		[4]int{1, 2, 3, 4},
		"IN PROGRESS",
		make([]Round, 0),
	}
	Games = make([]*Game, 0)
	Games = append(Games, &game1)
	r, err := MakeGuess(game1.ID, "1234")
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
	if game1.Status != "FINISHED" {
		t.Errorf("Expected game1 status to be FINISHED but instead received %s", game1.Status)
	}
	if len(game1.Rounds) != 1 {
		t.Errorf("Expected game1 to have 1 round but instead received %d", len(game1.Rounds))
	}

	game2 := Game{
		0,
		[4]int{1, 2, 3, 4},
		"IN PROGRESS",
		make([]Round, 0),
	}
	Games = append(Games, &game2)
	r, err = MakeGuess(game2.ID, "4321")
	if err != nil {
		t.Errorf("Expected no error but instead received %v", err)
	}
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
	if game2.Status != "IN PROGRESS" {
		t.Errorf("Expected game1 status to be IN PROGRESS but instead received %s", game1.Status)
	}
	if len(game2.Rounds) != 1 {
		t.Errorf("Expected game1 to have 1 round but instead received %d", len(game1.Rounds))
	}

	r, err = MakeGuess(game2.ID, "5678")
	if err != nil {
		t.Errorf("Expected no error but instead received %v", err)
	}
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
	if game2.Status != "IN PROGRESS" {
		t.Errorf("Expected game1 status to be IN PROGRESS but instead received %s", game1.Status)
	}
	if len(game2.Rounds) != 2 {
		t.Errorf("Expected game1 to have 2 round but instead received %d", len(game1.Rounds))
	}

	r, err = MakeGuess(game2.ID, "1243")
	if err != nil {
		t.Errorf("Expected no error but instead received %v", err)
	}
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
	if game2.Status != "IN PROGRESS" {
		t.Errorf("Expected game1 status to be IN PROGRESS but instead received %s", game1.Status)
	}
	if len(game2.Rounds) != 3 {
		t.Errorf("Expected game1 to have 3 round but instead received %d", len(game1.Rounds))
	}

	r, err = MakeGuess(game2.ID, "1234")
	if err != nil {
		t.Errorf("Expected no error but instead received %v", err)
	}
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
	if game2.Status != "FINISHED" {
		t.Errorf("Expected game1 status to be FINISHED but instead received %s", game1.Status)
	}
	if len(game2.Rounds) != 4 {
		t.Errorf("Expected game1 to have 4 round but instead received %d", len(game1.Rounds))
	}

	r, err = MakeGuess(-1, "1234")
	if err == nil {
		t.Errorf("Expected error but instead received %v", err)
	}
	if (r != Round{}) {
		t.Errorf("Expected empty Round but instead received %v", r)
	}
}
