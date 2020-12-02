package ascii_art

import (
	"bufio"
	"fmt"
	"os"
)

func ScanLine(typeOfWriting string,num int) string {

	fileName, err := os.Open("../template/" + typeOfWriting)
	scanner := bufio.NewScanner(fileName)
	checkLine := 0
	printLine := ""

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	for scanner.Scan() {

		if checkLine == num {
			printLine = scanner.Text()
		}
		checkLine++

	}

	return printLine
}
