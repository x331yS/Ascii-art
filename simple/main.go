package main

import (
	asciiart ".."
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("\nERROR: Please add an argument.")
		return
	}
	typeOfWriting := "standard.txt"

	s := os.Args[1]
	for _, v := range os.Args[2:] {
				s += " " + v
		}

	prev := '0'
	newLine := false
	for _, v := range s {
		if v == 'n' && prev == '\\' {
			newLine = true
		}
		prev = v
	}

	var result string

	if newLine {

		args := strings.Split(s, "\\n")
		for _, v := range args {

			for i := 0; i < 8; i++ {

				for _, w := range v {

					result += asciiart.ScanLine(typeOfWriting,1 + int(w-32)*9 + i)

				}

				fmt.Println(result)
				result = ""
			}

		}

	} else {

		for i := 0; i < 8; i++ {

			for _, v := range s {

				result += asciiart.ScanLine(typeOfWriting,1 + int(v-32)*9 + i)

			}

			fmt.Println(result)
			result = ""
		}

	}

}

