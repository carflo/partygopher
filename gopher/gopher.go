package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell"
)

const (
	frameCount = 23
	delay      = 80 // milliseconds
)

var (
	// XTERM COLORS
	// https://jonasjacek.github.io/colors/
	colors = []tcell.Color{
		tcell.Color(202), //Magenta,
		tcell.Color(206), //HotPink
		tcell.Color(207), //MediumOrchid1
		tcell.Color(171), //Magenta2
		tcell.Color(111), //SkyBlue2
		tcell.Color(87),  //DarkSlateGray2
		tcell.Color(118), //Chartreuse1
		tcell.Color(220), //Gold1
		tcell.Color(209), //Salmon1
		tcell.Color(202), //OrangeRed1
	}
	screen tcell.Screen
)

func main() {
	loops := flag.Int("loops", 1, "Number of times to loop the animation")
	flag.Parse()

	screen = initializeScreen()
	frames := initializeData()
	for i := 0; i < *loops; i++ {
		for j := 0; j < frameCount; j++ {
			draw(frames[j], colors[j%len(colors)])
			time.Sleep(delay * time.Millisecond)
		}
	}

	screen.Fini()
}

func initializeScreen() tcell.Screen {
	styling := tcell.StyleDefault
	screen, err := tcell.NewScreen()
	errHandler(err)
	err = screen.Init()
	errHandler(err)

	screen.SetStyle(styling.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	screen.Clear()

	return screen
}

func initializeData() []string {
	frames := []string{}

	for i := 0; i < frameCount; i++ {
		frame, err := Asset("frame" + strconv.Itoa(i))
		errHandler(err)
		frames = append(frames, string(frame))
	}

	return frames
}

func errHandler(e error) {
	if e != nil {
		fmt.Printf("%v\n", e)
		os.Exit(1)
	}
}

func draw(frame string, color tcell.Color) {

	lines := strings.Split(frame, "\n")

	for x, line := range lines {
		for y, cell := range line {
			screen.SetCell(y, x, tcell.StyleDefault.Foreground(color), cell)
		}
	}
	screen.Show()
}
