package databases

import "testing"

func TestCreateStack(t *testing.T) {
	key := randString(12)
	res, err := db.CreateStack(key)
	if err != nil {
		t.Error("expected", res, "got error", err)
	}
	secondRes, secondErr := db.CreateStack(key)
	if secondErr == nil {
		t.Error("expected", "error message", "got", secondRes)
	}
}

func TestDelStack(t *testing.T) {
	key := randString(12)
	db.CreateStack(key)
	db.DelStack(key)
	res, err := db.CreateStack(key)
	if err != nil {
		t.Error("expected", res, "got error", err)
	}
}

func TestGetStack(t *testing.T) {
	key := randString(12)
	res, err := db.GetStack(key)
	if err == nil {
		t.Error("expected", "error", "got", res)
	}
	newStack, _ := db.CreateStack(key)
	secondRes, secondErr := db.GetStack(key)
	if secondErr != nil || secondRes != newStack {
		t.Error("expected", newStack, "got", secondErr)
	}
}
