package utils

import (
	"testing"

	"github.com/mar-coding/goch01/app/utils"
)

func Test_GetHashStr(t *testing.T) {
	// test if my hash function works correctly
	// with sha256

	// Test case 1
	input1 := "kJnDmqwZmiJQPuaIBmKy"
	expectedOutput1 := "71bbfc6b489a7f6aea5c82ca92b57478690a5a8145fbc3aae10b057ee2756439"
	output1 := utils.GetHashStr(input1)
	if output1 != expectedOutput1 {
		t.Errorf("Expected %s, but got %s", expectedOutput1, output1)
	}

	// Test case 2
	input2 := "kJnDmqwaqe"
	expectedOutput2 := "a972f2898da6d27d5a2b3d10563ac3606fb8e7091134e72902a63fd1d13db6b6"
	output2 := utils.GetHashStr(input2)
	if output2 != expectedOutput2 {
		t.Errorf("Expected %s, but got %s", expectedOutput2, output2)
	}

}

func Test_GetRandomString(t *testing.T) {
	// Test if the GetRandomString function works correctly

	// Test case 1
	input1 := 0
	expectedOutput1 := 1
	output1 := len(utils.GetRandomString(input1))
	if output1 != expectedOutput1 {
		t.Errorf("Expected string length: %d, but got %d", expectedOutput1, output1)
	}

	// Test case 2
	input2 := -2
	expectedOutput2 := 1
	output2 := len(utils.GetRandomString(input2))
	if output2 != expectedOutput2 {
		t.Errorf("Expected string length: %d, but got %d", expectedOutput2, output2)
	}

	// Test case 3
	input3 := 5
	expectedOutput3 := 5
	output3 := len(utils.GetRandomString(input3))
	if output3 != expectedOutput3 {
		t.Errorf("Expected string length: %d, but got %d", expectedOutput3, output3)
	}
}

func Test_GetSumOfStr(t *testing.T) {
	// Test case 1
	input1 := "dc0783c1fbf74721558764422dd0eebba81b318b320865815e730a69216f4bf8"
	expectedOutput1 := 188
	output1 := utils.GetSumOfStr(input1)
	if output1 != expectedOutput1 {
		t.Errorf("Expected %d, but got %d", expectedOutput1, output1)
	}

	// Test case 2
	input2 := "442f254a"
	expectedOutput2 := 21
	output2 := utils.GetSumOfStr(input2)
	if output2 != expectedOutput2 {
		t.Errorf("Expected %d, but got %d", expectedOutput2, output2)
	}
}
