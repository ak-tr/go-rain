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
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	tm "github.com/buger/goterm"
)

const (
	HEAVY  string = "|"
	MEDIUM string = ":"
	LIGHT  string = "."
)

// Types
type Drop struct {
	char        string
	speed, x, y int
}

func (d *Drop) fall() {
	d.y += d.speed
}

type Drops []Drop

func (ds Drops) add(d Drop) Drops {
	ds = append(ds, d)
	return ds
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		tm.Clear()                                    // Clear screen
		tm.MoveCursor(1, 1)                           // Reset cursor
		tm.Print("\033[?25hhttps://github.com/ak-tr") // Show cursor
		tm.Flush()                                    // Mandatory flush
		os.Exit(0)                                    // Exit
	}()

	// Print escape code to hide cursor
	tm.Printf("\033[?25l")

	// Set seed for rand
	rand.Seed(time.Now().UnixNano())

	// Get height and width of terminal
	height := tm.Height()
	width := tm.Width()

	// Create array of drops
	var drops Drops
	var dropTypes = []string{HEAVY, MEDIUM, LIGHT}

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
			drop.fall()

			// If y value of drop is more than height, skip loop
			if drop.y > height {
				continue
			}

			// Append to tmp array only if previous gaurd clause is false
			tmp = append(tmp, Drop{drop.char, drop.speed, drop.x, drop.y})

			// Move cursor to location and print character to screen
			tm.MoveCursor(drop.x, drop.y)
			if drop.char != HEAVY {
				tm.Printf("\033[37m%s", drop.char)
				continue
			}
			tm.Printf("\033[97m%s", drop.char)
		}

		drops = tmp

		// Generate new drops at the end of each loop
		for _, dropType := range dropTypes {
			for i := 0; i < getDropsPerLine(width); i++ {
				j := generateRandomNumber(0, width)
				drops = drops.add(Drop{dropType, getSpeed(dropType), j, 0})
			}
		}

		// Flush required
		tm.Flush()

		// Sleep for 50ms
		time.Sleep(time.Millisecond * 50)
	}
}

// Get the speed of the drop based on the drop type
// Heavy drops fall by 1 cell per loop
// Light drops fall by 3 cells per loop
func getSpeed(dropType string) int {
	switch dropType {
	case HEAVY:
		return generateRandomNumber(1, 2)
	case MEDIUM:
		return generateRandomNumber(2, 3)
	case LIGHT:
		return generateRandomNumber(3, 4)
	default:
		return 1
	}
}

// Get width of screen, divide by 100 and then round up
// If value is more than 1, return random number between 1 and 2
// otherwise return 1
func getDropsPerLine(w int) int {
	if int(math.Ceil(float64(w)/100)) > 1 {
		return generateRandomNumber(1, 2)
	}
	return 1
}

// Generate c random numbers of range min to max
func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
