package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mar-coding/goch01/app/models"
	"github.com/mar-coding/goch01/app/utils"
)

func flag_usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
}

func main() {

	var (
		benchmark       = flag.Bool("benchmark", false, "Benchmark Mode:[true|false] (default false)")
		RAND_STR_NUMBER = flag.Int("str", 10, "Number of random string[int]")
		STR_LENGTH      = flag.Int("len", 20, "Length of each string[int]")
		algorithm       = flag.String("algo", "Insert in sorted array", "Algorithm for sort slice:[bubble|merge]")
	)

	flag.Usage = flag_usage
	flag.Parse()

	//this one use for insert on sorted slice
	strs1 := []models.Str{}

	//this one use for use merge and bubble sorts(or benchmark)
	strs2 := []models.Str{}

	temp := ""
	temp_sum := 0
	for i := 0; i < *RAND_STR_NUMBER; i++ {
		temp = utils.GetHashStr(utils.GetRandomString(*STR_LENGTH))
		temp_sum = utils.GetSumOfStr(temp)
		if temp_sum > 50 {

		}
		newStr := models.Str{
			Value: temp,
			Sum:   int32(temp_sum),
		}
		strs1 = utils.InsertSorted(strs1, newStr)

		strs2 = append(strs2, newStr)

	}

	if *benchmark {
		utils.Benchmark(strs2)
	} else {
		if *algorithm == "merge" {
			utils.MergeSort(strs2)
			fmt.Println(strs2)
		} else if *algorithm == "bubble" {
			utils.BubbleSort(strs2)
			fmt.Println(strs2)
		} else {
			fmt.Println(strs1)
		}
	}

}
