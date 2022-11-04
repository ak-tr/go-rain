/*

                       .::.
                  .:'  .:
        ,MMM8&&&.:'   .:'
       MMMMM88&&&&  .:'
      MMMMM88&&&&&&:'
      MMMMM88&&&&&&
    .:MMMMM88&&&&&&
  .:'  MMMMM88&&&&
.:'   .:'MMM8&&&'
:'  .:'
'::'    ak-tr

  Rain in your terminal...
	that's it.

*/

package main

import (
	"math/rand"
	"time"

	tm "github.com/buger/goterm"
)

const (
	Heavy string = "|"
	Light string = ":"
)

type Drop struct {
	char        string
	speed, x, y int
}

func main() {
	// Set seed for rand
	rand.Seed(time.Now().UnixNano())

	// Get height and width of terminal
	height := tm.Height()
	width := tm.Width()

	// Create array of drops
	var drops []Drop
	var dropTypes = []string{Heavy, Light}

	// For the height of the terminal
	// create a random drop
	for i := 0; i < height; i++ {
		for _, dropType := range dropTypes {
			ints := GenerateMultipleRandomNumbers(3, 0, width)

			for _, j := range ints {
				drops = append(drops, Drop{dropType, GetSpeed(dropType), j, i})
			}
		}
	}

	for _, drop := range drops {
		tm.MoveCursor(drop.x, drop.y)
		tm.Printf(string(drop.char) + "\033[?25l")
	}

	// Infinite loop
	for {
		tm.Clear()

		// For each drop
		for idx := range drops {
			// Get reference to drop and change location
			drop := &drops[idx]
			drop.y += drop.speed

			if drop.y > height {
				RemoveDrop(drops, idx)
			}

			// Move cursor to location and print to screen
			tm.MoveCursor(drop.x, drop.y+drop.speed)
			tm.Printf(string(drop.char))
		}

		for _, dropType := range dropTypes {
			ints := GenerateMultipleRandomNumbers(3, 0, width)

			for _, j := range ints {
				drops = AddDrop(Drop{dropType, GetSpeed(dropType), j, 0}, drops)
			}
		}

		tm.Flush()

		// Sleep for 50ms
		time.Sleep(time.Millisecond * 50)
	}
}

func GetSpeed(dropType string) int {
	switch dropType {
	case Heavy:
		return 1
	case Light:
		return 3
	}
	return 1
}

func AddDrop(d Drop, ds []Drop) []Drop {
	ds = append(ds, d)
	return ds
}

func RemoveDrop(s []Drop, i int) []Drop {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func GenerateMultipleRandomNumbers(c, min, max int) []int {
	nums := make([]int, c)

	for i := 0; i < c; i++ {
		nums[i] = rand.Intn(max-min+1) + min
	}

	return nums
}
