package aoc2022

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type DayResult struct {
	Part1 string
	Part2 string
}

type Day interface {
	Exec(input string) (*DayResult, error)
}

var Days = make(map[int]Day, 0)

const year = 2022

func FetchInput(day int, cookie string) (string, error) {
	filePath := createInputFilePath(year, day)

	_, err := os.Stat(filePath)
	if err != nil {
		body, err := fetchInput(day, cookie)
		if err != nil {
			return "", err
		}

		if err := writeToFile(filePath, body); err != nil {
			return "", err
		}

		return string(body), nil
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(b), err
}

func createInputFilePath(year, day int) string {
	return fmt.Sprintf("./src/%v/day%v.txt", year, day)
}

func fetchInput(day int, cookie string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%v/input", day), nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	req.Header.Set("cookie", fmt.Sprintf("session=%v", cookie))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func writeToFile(fileName string, body []byte) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	return nil
}
