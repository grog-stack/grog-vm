/* Very simple, cursor based display.

The screen is rectangular.
*/

package vm

import (
	"fmt"
	"runtime"
	"time"
)

const (
	RESET_CURSOR      = 0x00
	MOVE_CURSOR_RIGHT = 0x10
	MOVE_CURSOR_LEFT  = 0x11
	MOVE_CURSOR_DOWN  = 0x12
	MOVE_CURSOR_UP    = 0x13
	FRAMES_PER_SECOND = 5
)

type Surface struct {
	cols int
	rows int
}

type Coordinate struct {
	col int
	row int
}

type Display struct {
	size   Surface
	cursor Coordinate
	pixels [][]*pixel
}

func (display Display) Read() byte {
	return 0
}

func (display Display) Write(value byte) {
	switch value {
	case MOVE_CURSOR_RIGHT:
		fmt.Println("Move cursor right")
		display.cursor.col++
	case MOVE_CURSOR_LEFT:
		fmt.Println("Move cursor left")
		display.cursor.col--
	case MOVE_CURSOR_DOWN:
		fmt.Println("Move cursor down")
		display.cursor.row++
	case MOVE_CURSOR_UP:
		fmt.Println("Move cursor up")
		display.cursor.row--
	default:
		fmt.Printf("Invalid operation: %X\n", value)
	}
}

func NewDisplay(cols int, rows int) *Display {
	display := Display{size: Surface{cols: cols, rows: rows}}
	go display.Init()
	return &display
}

func (d *Display) Init() {
	runtime.LockOSThread()

	window := initGlfw()
	//defer glfw.Terminate()

	program := initOpenGL()

	d.pixels = makePixels(d.size.cols, d.size.rows)

	for !window.ShouldClose() {
		renderStartTime := time.Now()
		draw(d.pixels, window, program)
		time.Sleep(delay(renderStartTime))
	}
}

func makePixels(cols int, rows int) [][]*pixel {
	pixels := make([][]*pixel, rows, cols)
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			pixels[x] = append(pixels[x], newPixel(x, y, rows, cols))
		}
	}

	return pixels
}

func delay(renderStartTime time.Time) time.Duration {
	return time.Second/time.Duration(FRAMES_PER_SECOND) - time.Since(renderStartTime)
}
