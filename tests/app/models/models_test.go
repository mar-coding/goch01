package utils

import (
	"testing"

	"github.com/mar-coding/goch01/app/models"
)

type test struct {
	value string
	sum   int32
}

func initTestCase(a []test) []models.Str {
	tc := []models.Str{}

	str1 := &models.Str{
		Value: a[0].value,
		Sum:   a[0].sum,
	}
	tc = append(tc, *str1)

	str2 := &models.Str{
		Value: a[1].value,
		Sum:   a[1].sum,
	}

	tc = append(tc, *str2)

	return tc
}

func Test_Lesser(t *testing.T) {
	// Define test cases
	testCases := []test{
		{"332a91e2ccfddeeade6bcdf402923983bbda9751c70d3e6b7ceab0e3af6e735e", 135},
		{"36d644b5a663d1104dc03cafd789c1c5fe24a46feba972d6af0fb4511edecd65", 144},
	}
	tc := initTestCase(testCases)
	expectedOutput := true
	output := tc[0].Lesser(&tc[1])
	if expectedOutput != output {
		t.Errorf("Expected '%t', but got '%t'", expectedOutput, output)
	}
}

func Test_Greater(t *testing.T) {
	// Define test cases
	testCases := []test{
		{"332a91e2ccfddeeade6bcdf402923983bbda9751c70d3e6b7ceab0e3af6e735e", 135},
		{"36d644b5a663d1104dc03cafd789c1c5fe24a46feba972d6af0fb4511edecd65", 144},
	}
	tc := initTestCase(testCases)
	expectedOutput := false
	output := tc[0].Greater(&tc[1])
	if expectedOutput != output {
		t.Errorf("Expected '%t', but got '%t'", expectedOutput, output)
	}
}

func Test_Equal(t *testing.T) {
	// Define test cases
	testCases := []test{
		{"f8dbb755df217e7e022fffb38f3ae0de9ccf7f12ce6f50860f64073d9bbd3f8c", 144},
		{"36d644b5a663d1104dc03cafd789c1c5fe24a46feba972d6af0fb4511edecd65", 144},
	}
	tc := initTestCase(testCases)
	expectedOutput := true
	output := tc[0].Equal(&tc[1])
	if expectedOutput != output {
		t.Errorf("Expected '%t', but got '%t'", expectedOutput, output)
	}
}

func Test_Swap(t *testing.T) {
	// Define test cases
	inputCases := []test{
		{"332a91e2ccfddeeade6bcdf402923983bbda9751c70d3e6b7ceab0e3af6e735e", 135},
		{"36d644b5a663d1104dc03cafd789c1c5fe24a46feba972d6af0fb4511edecd65", 144},
	}
	input := initTestCase(inputCases)

	expectedCase := []test{
		{"36d644b5a663d1104dc03cafd789c1c5fe24a46feba972d6af0fb4511edecd65", 144},
		{"332a91e2ccfddeeade6bcdf402923983bbda9751c70d3e6b7ceab0e3af6e735e", 135},
	}
	expected := initTestCase(expectedCase)
	input[0].Swap(&input[1])
	is_swap_worked := func(a []models.Str, b []models.Str) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if !a[i].Equal(&b[i]) {
				return false
			}
		}
		return true
	}
	if !is_swap_worked(input, expected) {
		t.Errorf("Expected '%v', but got '%v'", input, expected)
	}
}
