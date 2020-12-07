package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

func Output(args []string) {

	if len(args) < 2 {

		noArgument()
		outputUsage()
		errorExit()

	} else if len(args) < 4 {

		notEnoughArguments()
		outputUsage()
		errorExit()

	} else if len(args) > 4 {

		tooManyArguments()
		outputUsage()
		errorExit()

	}

	if formatCheck(args) == false {

		formatFileUnknown()
		outputUsage()
		errorExit()

	}

	var text = args[1]
	var format = args[2]
	var output = outputFlag(args)
	var result = ""

	var file, err = os.Create(output)

	if err != nil {

		errorError()
		errorExit()

	}

	defer file.Close()

	if newlineCheck(text) { // == true

		var splitText = strings.Split(text, "\\n")

		for _, v := range splitText {

			for i := 0; i < 8; i++ {

				for _, v2 := range v {

					result += scanner(format, 1+int(v2-32)*9+i)

				}

				fmt.Fprintln(file, result)
				result = ""

			}

		}

	} else { // if newlineCheck(text) == false

		for i := 0; i < 8; i++ {

			for _, v := range text {

				result += scanner(format, 1+int(v-32)*9+i)

			}

			fmt.Fprintln(file, result)
			result = ""

		}

	}

}
