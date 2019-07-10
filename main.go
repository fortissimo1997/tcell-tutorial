package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	s.Clear()
	w, h := s.Size()

	st := tcell.StyleDefault.Background(tcell.ColorBlack)

	y := rand.Intn(h)

	for x := w - 1; x >= 0; x-- {
		color := rand.Int31n(0x1000000)
		stStr := st.Foreground(tcell.NewHexColor(color))
		s.Clear()
		s.Fill(' ', st)
		s.SetContent(x, y, '0', []rune{'+', '-', '*', '/', '%'}, stStr)
		s.Show()
		time.Sleep(100 * time.Millisecond)
	}

	s.Fini()
}
