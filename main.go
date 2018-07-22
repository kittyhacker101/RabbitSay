// RabbitSay by kittyhacker101
package main

import (
	"bufio"
	"fmt"
	"github.com/Code-Hex/go-wordwrap"
	"github.com/bbrks/wrap"
	"os"
	"strings"
)

const (
	boxtop     = "|￣￣￣￣￣￣￣￣￣￣￣￣|\n"
	boxbottom  = "|＿＿＿＿＿＿＿＿＿＿＿＿|"
	boxsides   = "|"
	boxpadding = " "
	figure     = `
	       (\__/) ||
	       (•ㅅ•) ||
	      /  　  づ`
	textarea = 24 - (len(boxpadding) * 2)
)

// advancedWordWrap uses various word wrappers on the text.
func advancedWordWrap(input string, length int) string {
	input = strings.TrimSpace(input)
	input = wrap.Wrap(input, length)
	return wordwrap.WrapString(input, uint(length))
}

// padString adds padding to the string if required.
func padString(input string, padding string, length int) string {
	if len(input) >= length {
		return input
	}

	for (len(input) % length) != 0 {
		input = input + padding
	}
	return input
}

func genRabbit(input []string) string {
	rabbit := boxtop
	for i := range input {
		if input[i] == "" {
			continue
		}
		rabbit = rabbit + boxsides + boxpadding + padString(input[i], boxpadding, textarea) + boxpadding + boxsides + "\n"
	}
	return rabbit + boxbottom + figure
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text for your rabbit to say below.")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Unable to get console input!")
		return
	}

	// Trim the input, and then wrap the words
	input = advancedWordWrap(input, textarea)
	splitInput := strings.Split(input, "\n")
	rabbit := genRabbit(splitInput)

	fmt.Println("Warning : You may need to add/remove whitespace for the text to display properly.")
	fmt.Println(rabbit)
}
