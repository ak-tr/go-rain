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
	HEAVY string = "|"
	LIGHT string = ":"
)

type Drop struct {
	char        string
	speed, x, y int
}

func main() {
	// Print escape code to hide cursor
	tm.Printf("\033[?25l")

	// Set seed for rand
	rand.Seed(time.Now().UnixNano())

	// Get height and width of terminal
	height := tm.Height()
	width := tm.Width()

	// Create array of drops
	var drops []Drop
	var dropTypes = []string{HEAVY, LIGHT}

	for { // Infinite loop
		// Clear screen on each loop
		tm.Clear()

		// Capture resize
		if height != tm.Height() || width != tm.Width() {
			// Update height and width to new height
			height = tm.Height()
			width = tm.Width()

			// Loop through all drops...
			for idx := range drops {
				drop := drops[idx]

				// ...and remove if off screen
				if drop.x >= width || drop.y >= height {
					removeDrop(drops, idx)
				}
			}
		}

		// For each drop
		for idx := range drops {
			// Get reference to drop and fall
			drop := &drops[idx]
			fallDrop(drop)

			// If drop falls off screen, remove it from drops array
			if drop.y > height {
				removeDrop(drops, idx)
				continue
			}

			// Move cursor to location and print character to screen
			tm.MoveCursor(drop.x, drop.y)
			tm.Printf(string(drop.char))
		}

		// Generate new drops at the end of each loop
		for _, dropType := range dropTypes {
			ints := generateMultipleRandomNumbers(2, 0, width)

			for _, j := range ints {
				drops = addDrop(Drop{dropType, getSpeed(dropType), j, 0}, drops)
			}
		}

		// Flush required
		tm.Flush()

		// Sleep for 50ms
		time.Sleep(time.Millisecond * 50)
	}
}

func fallDrop(d *Drop) {
	d.y += d.speed
}

// Add a pre-specified drop to the provided drops array
func addDrop(d Drop, ds []Drop) []Drop {
	ds = append(ds, d)
	return ds
}

// Remove a drop from the drop array by index
func removeDrop(s []Drop, i int) []Drop {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Get the speed of the drop based on the drop type
// Heavy drops fall by 1 cell per loop
// Light drops fall by 3 cells per loop
func getSpeed(dropType string) int {
	switch dropType {
	case HEAVY:
		return 1
	case LIGHT:
		return 3
	default:
		return 1
	}
}

// Generate c random numbers of range min to max
func generateMultipleRandomNumbers(c, min, max int) []int {
	nums := make([]int, c)

	for i := 0; i < c; i++ {
		nums[i] = rand.Intn(max-min+1) + min
	}

	return nums
}
