# go challenge

A Go program that generates random strings, hashes them using SHA256, sums up all the numbers in the resulting hash strings, and adds them to an ordered slice if the sum is greater than 50, for further usage.

---
## run

To run the Example Project, you'll need to have Go installed on your machine(it is based on go v1.18).

Once you have Go installed, you can run the project using the following command:

`go run cmd/main.go`

---
### The available options are:

- `-h, --help`: Display help information.
- `-algo`
    Algorithm for sort slice:[bubble|merge] (default "Insert in sorted array")
- `-benchmark`
    Benchmark Mode:[true|false] (default false)
    It checks the performance between parallel merge sort and bubble sort to determine which one is faster.
- `-len`
    Length of each string[int] (default 20)
- `-str`
    Number of random string[int] (default 10)