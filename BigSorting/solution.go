package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	len   int
	value string
}

func sortItem(lst []Item) {
	sort.SliceStable(lst, func(i, j int) bool {
		itemi, itemj := lst[i], lst[j]
		switch {
		case itemi.len == itemj.len:
			return itemi.value < itemj.value
		case itemi.len != itemj.len:
			return itemi.len < itemj.len
		default:
			return true
		}
	})
}

// Complete the bigSorting function below.
func bigSorting(unsorted []string) []string {
	var lst []Item
	for _, k := range unsorted {
		var item Item
		item.len = len(k)
		item.value = k
		lst = append(lst, item)
	}

	sortItem(lst)
	var sorted []string
	for _, k := range lst {
		sorted = append(sorted, k.value)
	}

	return sorted

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var unsorted []string

	for i := 0; i < int(n); i++ {
		unsortedItem := readLine(reader)
		unsorted = append(unsorted, unsortedItem)
	}

	result := bigSorting(unsorted)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func twoSum(nums []int, target int) []int {
	mapReturn := make(map[int]int)
	for k, v := range nums {
		numNeed := target - v
		if value, ok := mapReturn[numNeed]; ok {
			return []int{value, k}
		} else {
			mapReturn[v] = k
		}
	}
	return []int{}
}
