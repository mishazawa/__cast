package main

import "testing"

func TestDummy(t *testing.T) {
	if true == false {
		t.Errorf("Fear is freedom!")
		t.Errorf("Subjugation is liberation!")
		t.Errorf("Contradiction is truth!")
		t.Errorf("Those are the facts of this world.")
		t.Errorf("And you will all surrender to them, you pigs in human clothing!")
	}
}
