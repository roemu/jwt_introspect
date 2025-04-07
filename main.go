package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

type IntrospecResult struct {
	Header  string `json:"header"`
	Payload string `json:"payload"`
}

type IntrospecParseResult struct {
	ExpiresAt UnixTime `json:"exp"`
	IssuedAt UnixTime `json:"iat"`
	NotBefore UnixTime `json:"nbf"`
	Subject string `json:"sid"`
	Issuer string `json:"iss"`
}

func main() {
	headerOutput := flag.Bool("header", false, "Standalone flag, if set, will only output header part of jwt")
	parsedOutput := flag.Bool("parsed", false, "Standalone flag, if set, will output partially parsed jwt token with human readable values")
	reader, err := DetermineReaderFromFlags()
	if err != nil {
		log.Fatalf("Error determining input source: %s\n", err)
	}

	res, err := IntrospectFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	if *parsedOutput {
		out, err := json.Marshal(ParsePayload(res.Payload))
		if err != nil {
			log.Fatal("Unable to unmarshal parsed jwt token: ", err)
		}
		fmt.Print(string(out))
		
	} else if *headerOutput {
		fmt.Print(res.Header)
	} else {
		fmt.Print(res.Payload)
	}
}

func ParsePayload(s string) IntrospecParseResult {
	var parsed IntrospecParseResult
	err := json.Unmarshal([]byte(s), &parsed)
	if err != nil {
		log.Fatal("Unable to parse jwt token into human readable format.")
	}
	return parsed
}

func IntrospectFromReader(r io.Reader) (IntrospecResult, error) {
	output, err := io.ReadAll(r)
	introSpecResult := IntrospecResult{}

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

func DetermineReaderFromFlags() (io.Reader, error) {
	stdinInput := flag.Bool("stdin", false, "Standalone flag, cannot be combined with other input flags")
	clipInput := flag.Bool("clipboard", false, "Standalone flag, cannot be combined with other input flags")
	fileInput := flag.String("file", "", "--file=<path> \nStandalone flag, cannot be combined with other input flags.")

	flag.Parse()
	if *stdinInput {
		return bufio.NewReader(os.Stdin), nil
	}
	if *clipInput {
		c := clipboard.New()
		clip, err := c.PasteText()
		if err != nil {
			return nil, fmt.Errorf("Unable to fetch clipboard data, please use other option")
		}
		return strings.NewReader(clip), nil
	}
	if *fileInput != "" {
		f, err := os.Open(*fileInput)
		if err != nil {
			return nil, fmt.Errorf("Cannot open file at %s: %s\n", *fileInput, err)
		}
		return f, nil
	}
	if flag.NArg() > 0 {
		return strings.NewReader(flag.Arg(0)), nil
	}

	return nil, errors.New("No input source chosen, use --help to view available commands")
}

type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)
	return nil
}
