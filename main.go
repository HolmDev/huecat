package main

import (
	"regexp"
	"bufio"
	"fmt"
	"os"
	"io"
)



func main() {
	if os.Args[1] == "-h" || os.Args[1] == "--help"{
		fmt.Println("huecat v1.0. Colorify your command output!")
		fmt.Println("usage: huecat [-h]");
		fmt.Println("example: echo \"#ffffff\" | huecat");
		os.Exit(0);
	}

	reader := bufio.NewReader(os.Stdin);
	input := "";

	for {
		char, _, err := reader.ReadRune();
		if err != nil {
			if err == io.EOF {
				break;
			} else {
				panic(err)
			}
		}
		input += string(char);
	}

	hex24bit := regexp.MustCompile(`#([0-9A-Fa-f]){6}`);
	colorified := hex24bit.ReplaceAllStringFunc(input, colorify);
	fmt.Println(colorified);
}
