package ascii_art

import (
	"../Ascii-pkg"
	"fmt"
	"strings"
)

func Justify(args []string) {

	if len(args) < 2 {

		ascii_art.NoArgument()
		ascii_art.JustifyUsage()
		ascii_art.ErrorExit()

	} else if len(args) < 4 {

		ascii_art.NotEnoughArguments()
		ascii_art.JustifyUsage()
		ascii_art.ErrorExit()

	} else if len(args) > 4 {

		ascii_art.TooManyArguments()
		ascii_art.JustifyUsage()
		ascii_art.ErrorExit()

	}

	if ascii_art.FormatCheck(args) == false {

		ascii_art.FormatFileUnknown()
		ascii_art.JustifyUsage()
		ascii_art.ErrorExit()

	}

	var text = args[1]
	var format = args[2]
	var align = ascii_art.JustifyFlag(args)
	var terminalWidth = ascii_art.TerminalWidth()
	var result = ""

	if align == "justify" {

		ascii_art.JustifyJustify(text, result, format, terminalWidth)

	} else if align == "left" || align == "right" || align == "center" {

		if ascii_art.NewlineCheck(text) == true {

			var splitText = strings.Split(text, "\\n")

			for _, v := range splitText {

				for i := 0; i < 8; i++ {

					for _, v2 := range v {

						result += ascii_art.Scanner(format, 1+int(v2-32)*9+i)

					}

					switch align {

					case "left":
						{

							fmt.Println(result)
							result = ""

						}

					case "right":
						{

							fmt.Print(ascii_art.FillSpaces(terminalWidth-len(result)), result)
							result = ""

						}

					case "center":
						{

							fmt.Print(ascii_art.FillSpaces((terminalWidth-len(result))/2), result, ascii_art.FillSpaces((terminalWidth-len(result))/2), "\n")
							result = ""

						}

					}
				}
			}

		} else if ascii_art.NewlineCheck(text) == false {

			for i := 0; i < 8; i++ {

				for _, v := range text {

					result += ascii_art.Scanner(format, 1+int(v-32)*9+i)

				}

				switch align {

				case "left":
					{

						fmt.Println(result)
						result = ""

					}

				case "right":
					{

						fmt.Print(ascii_art.FillSpaces(terminalWidth-len(result)), result)
						result = ""

					}

				case "center":
					{

						fmt.Print(ascii_art.FillSpaces((terminalWidth-len(result))/2), result, ascii_art.FillSpaces((terminalWidth-len(result))/2), "\n")
						result = ""

					}

				}

			}

		}

	}

}
