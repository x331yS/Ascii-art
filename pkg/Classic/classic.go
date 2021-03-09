package ascii_art

import (
	"../Ascii-pkg"
	"fmt"
	"strings"
)

func Classic(args []string)  {

	if len(args) < 2 {

		ascii_art.NoArgument()
		ascii_art.ClassicUsage()
		ascii_art.ErrorExit()

	} else if len(args) > 2 {

		ascii_art.TooManyArguments()
		ascii_art.ClassicUsage()
		ascii_art.ErrorExit()

	}

	var text = args[1]
	var format = "standard"
	var result = ""

	if ascii_art.NewlineCheck(text) { //  == true

		var splitText = strings.Split(text,"\\n")

		for _, v := range splitText {

			for i := 0; i < 8; i++ {

				for _, v2 := range v {

					result += ascii_art.Scanner(format, 1+int(v2-32)*9+i)

				}

				fmt.Println(result)
				result = ""

			}

		}

	} else { //if newlineCheck(text) == false

			for i := 0 ; i < 8 ; i++ {

				for _, v := range text {

					result += ascii_art.Scanner(format, 1 + int(v - 32) * 9 + i)

				}

				fmt.Println(result)
				result = ""

			}

	}

}