package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

	fileAcs, err := os.Create("asecendente.txt")
	defer fileAcs.Close()

	if err != nil {
		fmt.Println(err)
	}

	for i, str := range strings {
		fileAcs.WriteString(str)
		if i < len(strings)-1 {
			fileAcs.WriteString("\n")
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(strings)))

	fileDes, err := os.Create("descendente.txt")
	defer fileDes.Close()

	if err != nil {
		fmt.Println(err)
	}

	for i, str := range strings {
		fileDes.WriteString(str)
		if i < len(strings)-1 {
			fileDes.WriteString("\n")
		}
	}
}
