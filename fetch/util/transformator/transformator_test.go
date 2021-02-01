package transformator_test

import (
	"testing"

	"github.com/ramailh/backend/fetch/util/transformator"
)

func TestInterfacesToInts(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []interface{}
		output []int
	}{
		{
			desc:   "test 1",
			input:  []interface{}{1, 2, 3, 4, 5},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			desc:   "test 2",
			input:  []interface{}{nil},
			output: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ints := transformator.NewTransformator().FromInterfaces(tC.input).ToInts()

			for i := 0; i < len(ints); i++ {
				if ints[i] != tC.output[i] {
					t.Error("error: output not matched")
				}
			}

			t.Logf("success: %v", ints)
		})
	}
}

func TestStringsToInts(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []interface{}
		output []int
	}{
		{
			desc:   "test 1",
			input:  []interface{}{"1", "2", "3", "4", "5"},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			desc:   "test 2",
			input:  []interface{}{nil},
			output: []int{},
		},
		{
			desc:   "test 3",
			input:  []interface{}{"1", "2", "  3"},
			output: []int{1, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ints := transformator.NewTransformator().FromStrings(tC.input).ToInts()

			for i := 0; i < len(ints); i++ {
				if ints[i] != tC.output[i] {
					t.Error("error: output not matched")
				}
			}

			t.Logf("success: %v", ints)
		})
	}
}
