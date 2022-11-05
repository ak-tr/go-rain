# go-rain

![](https://user-images.githubusercontent.com/62529128/200142769-3650694b-3191-483c-94b7-3ff574e83562.gif)


This is a very basic implementation of a terminal rain effect in Go using the [goterm library](https://github.com/buger/goterm).

Inspired by [nkleeman](https://github.com/nkleemann)'s [implemenatation](https://github.com/nkleemann/ascii-rain) in C.

## Dependencies

As mentioned above, this program relies on `goterm` and `gchalk` to print the rain drops to the terminal.  If you decide you want to run the program, it should auto install any dependencies when you perform `go run` command.

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

~~Currently does not work with standard Windows CMD or Powershell - use some form of terminal emulator like Cmder.~~

Windows support is unknown, works for some, doesn't work for others.

<sup><sub>Also my first time using Golang, I appreciate any feedback...</sub></sup>
