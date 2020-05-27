package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 12)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	var target int32

	if target64, err := strconv.ParseInt(scanner.Text(), 10, 32); err != nil {
		panic(err)
	} else {
		target = int32(target64)
	}

	numbers := make(map[int32]int)

	for scanner.Scan() {
		var num int32

		if num64, err := strconv.ParseInt(scanner.Text(), 10, 32); err != nil {
			panic(err)
		} else {
			num = int32(num64)
		}

		num_times, ok := numbers[num]
		if !ok {
			numbers[num] = 0
			num_times = 0
		}

		numbers[num] = num_times + 1
	}

	result := 0

	for key, _ := range numbers {
		rest := target - key

		rest_value, ok := numbers[rest]
		if rest == key && ok && rest_value > 1 {
			result = 1
			break
		} else if rest != key && ok {
			result = 1
			break
		}
	}

	ioutil.WriteFile("output.txt", []byte(strconv.Itoa(result)), 0644)
}
