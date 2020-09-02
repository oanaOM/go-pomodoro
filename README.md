### Minimalistic pomodoro timer

A command-line interface (CLI) tool that helps you manage time more efficiently by applying the [pomodoro technique](https://en.wikipedia.org/wiki/Pomodoro_Technique).

## Description

- Start the timer: 25minutes by default
- After time elapsed put a mark on your paper/task. A sound will be triggered.
- Take a short break: 5min by default
- After 4 pomodoroes have run, take a longer break

## Prerequisite

You need to have Go (Golang) installed locally. You can find full [here](https://golang.org/doc/install) documentation on how to do this

## Instalation steps

1. Clone this repository locally `https://github.com/oanaOM/go-pomodoro`
2. Go into the cloned directory `cd go-pomodoro`
3. Build the project `go build`
4. Run the timer `./go-pomodoro`


## Features
The user can manually change the pomodoro time and breaks duration by using the following flags:

`--duration` allows the user to change the duration of a pomodoro timer 

`--intervals` allows the user to change the number of intervals

`--shortbreak` allows the user to change the duration of a short break


