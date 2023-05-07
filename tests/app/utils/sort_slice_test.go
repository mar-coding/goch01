package utils

import (
	"testing"

	"github.com/mar-coding/goch01/app/models"
	"github.com/mar-coding/goch01/app/utils"
)

func initArray() []models.Str {
	arrs := []models.Str{}
	newStr1 := models.Str{
		Value: "d8cb19c3ec2a3caa65b13373fc7840193e534c6cd044ed973f9baafeab2116cd",
		Sum:   150,
	}
	arrs = append(arrs, newStr1)

	newStr2 := models.Str{
		Value: "4fad107eececaea01a257052c186ca84b2a5f2e8504ccc32d3235e55c7366e4b",
		Sum:   141,
	}
	arrs = append(arrs, newStr2)

	newStr3 := models.Str{
		Value: "b39e19d73c402122a7cce627bd3c92aa5cb53652cbbe10ecec2db314d7edf8df",
		Sum:   131,
	}
	arrs = append(arrs, newStr3)

	newStr4 := models.Str{
		Value: "46bd63af5eeea892a5c5ed4322e76b603258bd57a64f188bc7734027aaaf69aa",
		Sum:   185,
	}
	arrs = append(arrs, newStr4)
	return arrs
}

func sortArray() []models.Str {
	arrs := []models.Str{}

	newStr4 := models.Str{
		Value: "46bd63af5eeea892a5c5ed4322e76b603258bd57a64f188bc7734027aaaf69aa",
		Sum:   185,
	}
	arrs = append(arrs, newStr4)
	newStr1 := models.Str{
		Value: "d8cb19c3ec2a3caa65b13373fc7840193e534c6cd044ed973f9baafeab2116cd",
		Sum:   150,
	}
	arrs = append(arrs, newStr1)

	newStr2 := models.Str{
		Value: "4fad107eececaea01a257052c186ca84b2a5f2e8504ccc32d3235e55c7366e4b",
		Sum:   141,
	}
	arrs = append(arrs, newStr2)

	newStr3 := models.Str{
		Value: "b39e19d73c402122a7cce627bd3c92aa5cb53652cbbe10ecec2db314d7edf8df",
		Sum:   131,
	}
	arrs = append(arrs, newStr3)

	return arrs

}

func Test_MergeSort(t *testing.T) {
	input := initArray()

	// Test case 1
	expectedOutput := sortArray()
	output := utils.MergeSort(input)

	is_equal := func(a []models.Str, b []models.Str) bool {
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
	if !is_equal(output, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, output)
	}
}

func Test_BubbleSort(t *testing.T) {
	input := initArray()

	// Test case 1
	expectedOutput := sortArray()
	output := utils.BubbleSort(input)

	is_equal := func(a []models.Str, b []models.Str) bool {
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
	if !is_equal(output, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, output)
	}
}

func Test_InsertSorted(t *testing.T) {
	input := initArray()

	// Test case 1
	expectedOutput := sortArray()

	output := []models.Str{}
	for _, i := range input {
		output = utils.InsertSorted(output, i)
	}

	is_equal := func(a []models.Str, b []models.Str) bool {
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
	if !is_equal(output, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, output)
	}
}
