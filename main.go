package main

import (
	"aoc2022"
	"fmt"
	"os"
	"strconv"
	"time"
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
			fmt.Printf("Unknown argument %v\n", v)
			continue
		}
		days = append(days, dayX)
	}

	if len(days) == 0 {
		// Run them all
		for i := 1; i <= len(aoc2022.Days); i++ {
			days = append(days, i)
		}
	}

	success := make(chan chanResult, len(days))
	errors := make(chan chanResult, len(days))

	resps := responses{success, errors}

	runDays(days, cookie, resps)

	doneCount := 0
	for {
		select {
		case res := <-success:
			fmt.Printf("Day %v:\t%v\n", res.day, res.message)
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
	fmt.Println("Running days", days)

	for _, day := range days {
		go runDay(day, cookie, r)
	}
}

func runDay(dayNum int, cookie string, r responses) {
	result := chanResult{day: dayNum}

	day, ok := aoc2022.Days[dayNum]
	if !ok {
		result.message = "not implemented"
		r.errors <- result
		return
	}

	input, err := aoc2022.FetchInput(dayNum, cookie)

	if err != nil {
		result.message = fmt.Sprintf("Error fetching %v: %v", dayNum, err.Error())
		r.errors <- result
		return

	}

	start := time.Now()

	res, err := day.Exec(input)

	elapsed := time.Since(start)
	if elapsed > time.Second {
		elapsed = elapsed.Truncate(time.Millisecond)
	} else {
		elapsed = elapsed.Truncate(time.Microsecond)

	}

	if err != nil {
		result.message = fmt.Sprintf("[Exec]: %v", err.Error())
		r.errors <- result
		return
	}

	result.message = fmt.Sprintf("[Part1]: %v [Part2]: %v (%v)", res.Part1, res.Part2, elapsed)
	r.success <- result
}
