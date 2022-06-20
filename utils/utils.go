package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error, msg string) {
	if e != nil {
		fmt.Println(msg)
		panic(e)
	}
}

func ReadFile(filename string) string {
	dat, err := os.ReadFile(filename)
	check(err, "unable to open file "+filename)
	return strings.TrimSpace(string(dat))
}

func ReadFileToLines(filename string) []string {
	return strings.Split(ReadFile(filename), "\n")
}

func ReadFileToIntArray(filename string) []int {
	lines := ReadFileToLines(filename)
	var data []int
	for _, x := range lines {
		value, _ := strconv.Atoi(x)
		data = append(data, value)
	}
	return data
}

func DoesExist(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("File " + filename + " exists")
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File " + filename + " does NOT exist")
		return false
	} else {
		fmt.Println("ERROR " + filename)
		return false
	}
}

func Sum[T Number](nums ...T) T {
	var res T = 0
	for _, n := range nums {
		res += n
	}
	return res
}

type Number interface {
	int | float32
}

func Map[T Number | string](mapper func(T) T, data []T) []T {
	var output []T
	for _, x := range data {
		output = append(output, mapper(x))
	}
	return output
}
