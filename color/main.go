package main

import (
	asciiart ".."
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("\nYou have to write for example \"YourText\" --color=Yourcolor")
		fmt.Println("The color : red, green, yellow, blue, purple, cyan, black")
		return
	}
	typeOfWriting := "standard.txt"
		colorBlack := "\033[30m"
		colorRed := "\033[31m"
		colorGreen := "\033[32m"
		colorYellow := "\033[33m"
		colorBlue := "\033[34m"
		colorPurple := "\033[35m"
		colorCyan := "\033[36m"
		colorReset := "\033[0m"

	color:= ""
	args := os.Args[1:]
	for _, v := range args {


		if len(v) > 8 && v[:8] == "--color=" {
			color = v[8:]

			switch color {
			case "red":
				color = colorRed
			case "green":
				color = colorGreen
			case "yellow":
				color = colorYellow
			case "blue":
				color = colorBlue
			case "purple":
				color = colorPurple
			case "cyan":
				color = colorCyan
			case "black":
				color = colorBlack
			}
		}

	}
	/*if os.Args[1] == "coloring" && len(args) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entrez votre texte :  ")
		scanner.Scan()
		text := scanner.Text()
		v := text[0:]
		for range text {
			fmt.Println("Choisissez la couleur pour cet argument : ")
			fmt.Println(v)
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			yourcolor := scanner.Text()


			fmt.Println(yourcolor)
		}
	}*/



	s := os.Args[1]
	for _, v := range os.Args[2:] {
		if v != "red" && v != "green" && v != "yellow" && v != "blue" && v != "purple" && v != "cyan" && v != "black" && (len(v) > 8 && v[:8] != "--color=") {
			s += " " + v

		}
	}
	if color != "\033[31m" && color != "\033[32m" && color != "\033[33m" && color != "\033[34m" && color != "\033[35m" && color != "\033[36m" && color != "\033[30m" {
		fmt.Println("\nYou have to write for example \"YourText\" --color=Yourcolor")
		fmt.Println("The color : red, green, yellow, blue, purple, cyan, black")
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

	var result string

	if newLine {

		args := strings.Split(s, "\\n")
		for _, v := range args {

			for i := 0; i < 8; i++ {

				for _, w := range v {

					result += asciiart.ScanLine(typeOfWriting, 1+int(w-32)*9+i)

				}
				fmt.Println(color, result)
				result = ""
			}

		}

	} else {

		for i := 0; i < 8; i++ {

			for _, v := range s {

				result += asciiart.ScanLine(typeOfWriting, 1+int(v-32)*9+i)

			}
			fmt.Println(color, result)
			result = ""
		}

	}

	fmt.Println(colorReset)

}
