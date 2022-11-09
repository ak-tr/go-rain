# go-rain

![Screen Recording 2022-11-08 at 03 06 08](https://user-images.githubusercontent.com/62529128/200467049-2b26a0ed-36cc-4f78-b8e8-8d22ee614035.gif)

This is a very basic implementation of a terminal rain effect in Go using the [goterm library](https://github.com/buger/goterm).

Inspired by [nkleeman](https://github.com/nkleemann)'s [implemenatation](https://github.com/nkleemann/ascii-rain) in C.

## Dependencies

As mentioned above, this program relies on `goterm` and `gookit/color` to print the rain drops to the terminal.  If you decide you want to run the program, it should auto install any dependencies when you perform `go run` command.

## Installation

Download the repo:

`$ git clone https://github.com/ak-tr/go-rain`

Move into directory:

`$ cd go-rain`

Run:

`$ go run main.go`

Alternatively, build with:

`$ go build`

Then, run:

`$ ./go-rain`

## Issues

Windows support added. Should have no issues there.

MacOS users, if you encounter an `xcrun: error: invalid active developer path` error, run:

`$ xcode-select --install`

<sup><sub>Also my first time using Golang, I appreciate any feedback...</sub></sup>
