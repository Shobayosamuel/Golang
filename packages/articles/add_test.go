package artist

import "testing"

func TestAdd(t *testing.T){
	expected := 4
	actual := Add(2, 2)

	if actual != expected {
		t.Errorf("Function does not add up, Expected: %d, Actual: %d", expected, actual)
	}
}

func TestMakeExcited(t *testing.T) {
	expected := "OMG SO EXCITING!"
	actual := MakeExcited("omg so exciting!")
	if actual != expected {
		t.Errorf("Average was incorrect! Expected: %s, Actual: %s", expected, actual)
	}
}