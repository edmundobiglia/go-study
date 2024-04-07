package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Please enter one word!")
		return
	}

	word := strings.ToLower(os.Args[1])
	wordLen := len(word)
	var evenWordLen int

	if wordLen%2 == 0 {
		evenWordLen = wordLen
	} else {
		evenWordLen = wordLen - 1
	}

	halfLen := evenWordLen / 2

	for letterPosition := 0; letterPosition <= halfLen; letterPosition++ {
		if word[letterPosition] != word[wordLen-(letterPosition+1)] {
			fmt.Printf("👎 NO, \"%s\" is NOT a palindrome.", word)
			break
		} else if letterPosition+1 == halfLen {
			fmt.Printf("👍 YES, \"%s\" IS a palindrome!", word)
		}
	}
}
