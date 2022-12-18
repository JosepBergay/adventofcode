# My Advent of Code adventure.

The aim of this repo is to share my aoc exercises, maybe we can comment them up, learn in the process and most importantly have some fun!

## Features
* Autofetch inputs 🎉
* Output answers via console 🖨

## Instructions
If you intend to use it you should set an env var in the shape of `SESSION_COOKIE=123-replace-with-yours` containing the session id so it can properly fetch all inputs. You can grab yours from https://adventofcode.com/ once logged in (devtools ftw!). 

## 2022

### Build

```bash
go build
```

### Run

Pass which days you want to run. Or pass nothing to run them all.
```bash
adventofcode 1 2 # Will run day 1 and 2

...

Running days [1 2]
Day 2:  [Part1]: 13484 [Part2]: 13433 (450µs)
Day 1:  [Part1]: 70369 [Part2]: 203002 (328µs)
```

### Development

[air](https://github.com/cosmtrek/air) is used for hot reload so you can just do:
```bash
air 1 2 # Will run day 1 and 2 each time you save a file
```

## 2021

### Usage
`npm run start`. Don't forget to install dependencies first, you fool!

You can also pass a num argument to start to run only that day:

```
npm run start 2

...

start Running Day 2
✔ Done in 495 ms
ℹ Answer:
ℹ Day #2 part1: 1882980, part2: 1971232560
```

</br>
Made with ♥, NodeJS and Golang
