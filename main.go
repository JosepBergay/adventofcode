package main

import (
	"aoc2022"
	"fmt"
	"os"
	"strconv"
)

const CookieName = "SESSION_COOKIE"

type chanResult struct {
	day     int
	message string
}

type responses struct {
	success chan<- chanResult
	errors  chan<- chanResult
}

func main() {
	cookie := os.Getenv(CookieName)

	if cookie == "" {
		fmt.Println("No cookie :(")
		return
	}

	days := make([]int, 0)

	for _, v := range os.Args[1:] {
		dayX, err := strconv.Atoi(v)

		if err != nil {
			fmt.Printf("Unknown argument %v skipping", v)
			continue
		}
		days = append(days, dayX)
	}

	if len(days) == 0 {
		fmt.Println("Nothing to do :(")
		return
	}

	success := make(chan chanResult, len(days))
	errors := make(chan chanResult, len(days))

	resps := responses{success, errors}

	runDays(days, cookie, resps)

	doneCount := 0
	for {
		select {
		case res := <-success:
			fmt.Printf("Day %v result: %v\n", res.day, res.message)
		case res := <-errors:
			fmt.Printf("Day %v error: %v\n", res.day, res.message)
		}
		doneCount++
		if doneCount >= len(days) {
			break
		}
	}
}

func runDays(days []int, cookie string, r responses) {
	fmt.Printf("Running days %v\n", days)

	for _, day := range days {
		go runDay(day, cookie, r)
	}
}

func runDay(dayNum int, cookie string, r responses) {

	fmt.Printf("Running day %v\n", dayNum)

	result := chanResult{day: dayNum}

	input, err := aoc2022.FetchInput(dayNum, cookie)

	if err != nil {
		result.message = fmt.Sprintf("Error fetching %v: %v", dayNum, err.Error())
		r.errors <- result
		return

	}

	day, ok := aoc2022.Days[dayNum]
	if !ok {
		result.message = "not implemented"
		r.errors <- result
		return
	}

	res, err := day.Exec(input)

	if err != nil {
		result.message = fmt.Sprintf("[Exec]: %v", err.Error())
		r.errors <- result
		return
	}

	result.message = fmt.Sprintf("[Part1]: %v, [Part2]: %v", res.Part1, res.Part2)
	r.success <- result
}