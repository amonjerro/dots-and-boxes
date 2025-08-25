package utils

import "testing"

type ReduceCase interface {
	TestReduce() int
	GetExpectedValue() int
}

type IntReduce struct {
	values        []int
	initialValue  int
	function      func(acc int, val int) int
	expectedValue int
}

func (i *IntReduce) TestReduce() int {
	return Reduce(i.values, i.function, i.initialValue)
}

func (i *IntReduce) GetExpectedValue() int {
	return i.expectedValue
}

type BoolReduce struct {
	values        []bool
	initialValue  int
	function      func(acc int, val bool) int
	expectedValue int
}

func (b *BoolReduce) TestReduce() int {
	return Reduce(b.values, b.function, b.initialValue)
}

func (b *BoolReduce) GetExpectedValue() int {
	return b.expectedValue
}

func TestReduce(t *testing.T) {
	testCases := []ReduceCase{
		&IntReduce{
			values:        []int{1, 2, 3, 4},
			initialValue:  0,
			function:      func(acc int, val int) int { return acc + val },
			expectedValue: 10,
		},
		&IntReduce{
			values:        []int{1, 2, 3, 4},
			initialValue:  9,
			function:      func(acc int, val int) int { return acc + val },
			expectedValue: 19,
		},
		&BoolReduce{
			values:       []bool{true, false, true, true},
			initialValue: 0,
			function: func(acc int, val bool) int {
				if !val {
					return acc * 0
				}
				return (1 + acc) * 2
			},
			expectedValue: 6,
		},
	}

	for _, tt := range testCases {
		if tt.TestReduce() != tt.GetExpectedValue() {
			t.Fatalf("Reduce calculation did not work. Expected %d, got %d", tt.TestReduce(), tt.GetExpectedValue())
		}
	}
}

func TestSet(t *testing.T) {
	set := &Set[int]{Elements: make(map[int]bool), Count: 0}

	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Remove(2)

	if set.Count != 2 {
		t.Fatalf("Set count off")
	}

	if set.Contains(2) {
		t.Fatalf("Item Removal Failed")
	}

	if !set.Contains(3) {
		t.Fatalf("Item retrieval failed")
	}

}
