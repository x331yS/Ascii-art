package ascii_art

import (
	"fmt"
	"strings"
)

func Justify(args []string) {

	if len(args) < 2 {

		noArgument()
		justifyUsage()
		errorExit()

	} else if len(args) < 4 {

		notEnoughArguments()
		justifyUsage()
		errorExit()

	} else if len(args) > 4 {

		tooManyArguments()
		justifyUsage()
		errorExit()

	}

	if formatCheck(args) == false {

		formatFileUnknown()
		justifyUsage()
		errorExit()

	}

	var text = args[1]
	var format = args[2]
	var align = justifyFlag(args)
	var terminalWidth = terminalWidth()
	var result = ""

	if align == "justify" {

		justifyJustify(text, result, format, terminalWidth)

	} else if align == "left" || align == "right" || align == "center" {

		if newlineCheck(text) == true {

			var splitText = strings.Split(text, "\\n")

			for _, v := range splitText {

				for i := 0; i < 8; i++ {

					for _, v2 := range v {

						result += scanner(format, 1+int(v2-32)*9+i)

					}

					switch align {

					case "left":
						{

							fmt.Println(result)
							result = ""

						}

					case "right":
						{

							fmt.Print(fillSpaces(terminalWidth-len(result)), result)
							result = ""

						}

					case "center":
						{

							fmt.Print(fillSpaces((terminalWidth-len(result))/2), result, fillSpaces((terminalWidth-len(result))/2), "\n")
							result = ""

						}

					}
				}
			}

		} else if newlineCheck(text) == false {

			for i := 0; i < 8; i++ {

				for _, v := range text {

					result += scanner(format, 1+int(v-32)*9+i)

				}

				switch align {

				case "left":
					{

						fmt.Println(result)
						result = ""

					}

				case "right":
					{

						fmt.Print(fillSpaces(terminalWidth-len(result)), result)
						result = ""

					}

				case "center":
					{

						fmt.Print(fillSpaces((terminalWidth-len(result))/2), result, fillSpaces((terminalWidth-len(result))/2), "\n")
						result = ""

					}

				}

			}

		}

	}

}
