package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		want     []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
		{OpAdd, []int{}, []byte{byte(OpAdd)}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		if len(instruction) != len(tt.want) {
			t.Errorf("instruction has wrong length: got %d, wanted %d", len(instruction), len(tt.want))
		}

		for i, b := range tt.want {
			if instruction[i] != tt.want[i] {
				t.Errorf("wrong byte at position %d: got %d, wanted %d", i, instruction[i], b)
			}
		}
	}
}

func TestInstructionString(t *testing.T) {
	instructions := []Instructions{
		Make(OpAdd),
		Make(OpConstant, 2),
		Make(OpConstant, 65535),
	}

	expected := `0000 OpAdd
0001 OpConstant 2
0004 OpConstant 65535
`

	concatted := Instructions{}
	for _, instruction := range instructions {
		concatted = append(concatted, instruction...)
	}

	if concatted.String() != expected {
		t.Errorf("instructions incorrectly formatted: got %q, wanted %q", concatted.String(), expected)
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op        Opcode
		operands  []int
		bytesRead int
	}{
		{OpConstant, []int{65535}, 2},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		def, err := Lookup(byte(tt.op))
		if err != nil {
			t.Fatalf("definition not found: %q", err)
		}

		operandsRead, n := ReadOperands(def, instruction[1:])
		if n != tt.bytesRead {
			t.Fatalf("wrong number of bytes read: got %d, wanted %d", n, tt.bytesRead)
		}

		for i, want := range tt.operands {
			if operandsRead[i] != want {
				t.Errorf("wrong operand: got %d, wanted %d", operandsRead[i], want)
			}
		}
	}
}
