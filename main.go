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
	"os"
	"os/signal"
	"syscall"
	"time"

	tm "github.com/buger/goterm"
	"github.com/jwalton/gchalk"
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
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		tm.Clear()             // Clear screen
		tm.MoveCursor(1, 1)    // Reset cursor
		tm.Printf("\033[?25h") // Show cursor
		tm.Flush()             // Mandatory flush
		os.Exit(0)             // Exit
	}()

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

		// Reset temp value on each loop
		tmp := drops[:0]

		// Capture resize
		if height != tm.Height() || width != tm.Width() {
			// Update height and width to new height
			height = tm.Height()
			width = tm.Width()

			// Loop through all drops...
			for idx := range drops {
				drop := drops[idx]

				// ...and ignore if off screen
				if drop.x >= width || drop.y >= height {
					continue
				}

				// Add to tmp slice
				tmp = append(tmp, Drop{drop.char, drop.speed, drop.x, drop.y})
			}

			drops = tmp
		}

		// For each drop
		for idx := range drops {
			// Get reference to drop and fall
			drop := &drops[idx]
			fallDrop(drop)

			// If y value of drop is more than height, skip loop
			if drop.y > height {
				continue
			}

			// Append to tmp array only if previous gaurd clause is false
			tmp = append(tmp, Drop{drop.char, drop.speed, drop.x, drop.y})

			// Move cursor to location and print character to screen
			tm.MoveCursor(drop.x, drop.y)
			if drop.char == LIGHT {
				tm.Print(gchalk.BrightBlack(drop.char))
				continue
			}
			tm.Print(drop.char)
		}

		drops = tmp

		// Generate new drops at the end of each loop
		for _, dropType := range dropTypes {
			ints := generateMultipleRandomNumbers(getDropsPerLine(width), 0, width)

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

func getDropsPerLine(w int) int {
	if w > 150 {
		return 2
	}
	return 1
}

// Generate c random numbers of range min to max
func generateMultipleRandomNumbers(c, min, max int) []int {
	nums := make([]int, c)

	for i := 0; i < c; i++ {
		nums[i] = rand.Intn(max-min+1) + min
	}

	return nums
}
