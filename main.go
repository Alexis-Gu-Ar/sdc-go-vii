package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func writeStringsToFile(strings []string, fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	for i, str := range strings {
		file.WriteString(str)
		if i < len(strings)-1 {
			file.WriteString("\n")
		}
	}
}

func main() {
	n := 5
	scanner := bufio.NewScanner(os.Stdin)
	strings := make([]string, n)
	fmt.Printf("Dame %d cadenas\n", n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		strings[i] = scanner.Text()
	}

	sort.Strings(strings)
	writeStringsToFile(strings, "asecendente.txt")

	sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	writeStringsToFile(strings, "descendente.txt")
}
