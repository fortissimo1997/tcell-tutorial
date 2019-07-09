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

	st := tcell.StyleDefault

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			color := rand.Int31n(0x1000000)
			st = st.Foreground(tcell.NewHexColor(color)).Background(tcell.Color(231))
			s.SetCell(x, y, st, '@')
		}
	}
	s.Show()

	time.Sleep(10 * time.Second)

	s.Fini()
}
