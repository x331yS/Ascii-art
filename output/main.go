package main

import (
	asciiart ".."
	"fmt"
	"os"
	"strings"

)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("\nYou have to write for example \"YourText\" typeOfWriting --output=nameOfNew.txt")
		return
	}

	typeOfWriting := ""
	file := ""
	args := os.Args[1:]
	for _, v := range args {

		switch v {
		case "thinkertoy" :
			typeOfWriting = "thinkertoy.txt"
		case "shadow" :
			typeOfWriting = "shadow.txt"
		case "standard" :
			typeOfWriting = "standard.txt"

		}

		if len(v) > 9 && v[:9] == "--output=" {
			file = v[9:]
		}
	}

	s := os.Args[1]
	for _, v := range os.Args[2:] {
		if v != "standard" && v != "shadow" && v != "thinkertoy" && (len(v) > 9 && v[:9] != "--output=") {
			s += " " + v
		}
	}
	if typeOfWriting != "standard.txt" && typeOfWriting!="shadow.txt" && typeOfWriting != "thinkertoy.txt" {
		fmt.Println("\nYou have to write for example \"YourText\" typeOfWriting --output=nameOfNew.txt")
		return
	}

	prev := '0'
	newLine := false
	for _, v := range s {
		if v == 'n' && prev == '\\' {
			newLine = true
		}
		prev = v
	}

	f, e := os.Create(file)
	if e != nil {
		fmt.Println("\nYou have to write for example \"YourText\" typeOfWriting --output=nameOfNew.txt")
		os.Exit(0)
	}
	defer f.Close()

	result := ""
	if newLine {

		args := strings.Split(s, "\\n")
		for _, v := range args {

			for i := 0; i < 8; i++ {

				for _, w := range v {

					result += asciiart.ScanLine(typeOfWriting,1 + int(w-32)*9 + i)

				}

				fmt.Fprintln(f, result)
				result = ""
			}

		}

	} else {

		for i := 0; i < 8; i++ {

			for _, v := range s {

				result += asciiart.ScanLine(typeOfWriting,1 + int(v-32)*9 + i)

			}

			fmt.Fprintln(f, result)
			result = ""
		}

	}

}