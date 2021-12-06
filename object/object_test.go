package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings %q and %q are identical but have different hash keys", hello1.Value, hello2.Value)
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings %q and %q are identical but have different hash keys", diff1.Value, diff2.Value)
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings %q and %q are different but have identical hash keys", hello1.Value, diff1.Value)
	}
}
