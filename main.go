package main

import (
	"bufio"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

type IntroSpecResult struct {
	Header  string `json:"header"`
	Payload string `json:"payload"`
}

func main() {
	var reader io.Reader
	stdinInput := flag.Bool("stdin", false, "Standalone flag, cannot be combined with other input flags")
	clipInput := flag.Bool("clipboard", false, "Standalone flag, cannot be combined with other input flags")
	fileInput := flag.String("file", "", "--file=<path> \nStandalone flag, cannot be combined with other input flags.")
	headerOutput := flag.Bool("header", false, "Standalone flag, if set, will only output header part of jwt")

	flag.Parse()
		log.Println(os.Args)

	if *stdinInput {
		reader = bufio.NewReader(os.Stdin)
	}
	if *clipInput {
		c := clipboard.New()
		clip, err := c.PasteText()
		if err != nil {
			log.Fatalf("Unable to fetch clipboard data, please use other option")
		}
		reader = strings.NewReader(clip)
	}
	if *fileInput != "" {
		f, err := os.Open(*fileInput)
		if err != nil {
			log.Fatalf("Cannot open file at %s: %s\n", *fileInput, err)
		}
		reader = f
	}
	if len(os.Args) >= 2 && !*stdinInput && !*clipInput && *fileInput == "" {
		reader = strings.NewReader(os.Args[1])
	}
	if len(os.Args) == 1 && !*stdinInput && !*clipInput && *fileInput == "" {
		log.Fatalf("No jwt token provided, use --stdin/--clipboard/--file=<path> or provide as argument to command")
	}
	res, err := IntrospectFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	if *headerOutput {
		fmt.Print(res.Header)
	} else {
		fmt.Print(res.Payload)
	}
}

func IntrospectFromReader(r io.Reader) (IntroSpecResult, error) {
	output, err := io.ReadAll(r)
	introSpecResult := IntroSpecResult{}

	s := strings.Split(string(output), ".")
	if len(s) != 3 {
		return introSpecResult, errors.New("Malformed jwt token")
	}

	header, err := base64.RawStdEncoding.DecodeString(s[0])
	if err != nil {
		return introSpecResult, errors.New(fmt.Sprintf("Error decoding header of jwt token: %s, %s", err, s[0]))
	}
	introSpecResult.Header = string(header)

	payload, err := base64.RawStdEncoding.DecodeString(s[1])
	if err != nil {
		return introSpecResult, errors.New(fmt.Sprintf("Error decoding payload of jwt token: %s", err))
	}
	introSpecResult.Payload = string(payload)

	return introSpecResult, nil
}
