package aoc2022

import (
	"fmt"
	"io"
	"net/http"
)

type Day interface {
	Day() int

	Parse(input string) (string, error)

	Exec(parsed string) (string, error)
}

var Days = make(map[int]Day, 0)

func FetchInput(day int, cookie string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%v/input", day), nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	req.Header.Set("cookie", fmt.Sprintf("session=%v", cookie))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
