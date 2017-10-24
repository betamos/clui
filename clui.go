package clui

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
	"time"
)

type Spinner struct {
	theme     []string
	tickSpeed time.Duration
	tickCount int
	status    string // enum?
	message   string
	ticker    *time.Ticker
	mtx       *sync.Mutex
}

func NewSpinner(message string) *Spinner {
	s := &Spinner{
		theme:     []string{"⠛", "⠹", "⢸", "⣰", "⣤", "⣆", "⡇", "⠏"},
		tickSpeed: 150 * time.Millisecond,
		tickCount: 0,
		status:    "working",
		message:   message,
		mtx:       &sync.Mutex{},
	}
	return s
}

func (s *Spinner) redraw() {
	s.mtx.Lock()
	fmt.Print("\u001b[2K") // clear line
	fmt.Print("\u001b[0G") // to column 0
	if s.status == "success" {
		color.Set(color.FgGreen)
		fmt.Print("✔")
		color.Unset()
	} else if s.status == "fail" {
		color.Set(color.FgRed)
		fmt.Print("✘")
		color.Unset()

	} else {
		color.Set(color.FgYellow)
		fmt.Print(s.theme[s.tickCount%len(s.theme)])
		color.Unset()
	}
	fmt.Printf(" %v", s.message)
	s.mtx.Unlock()

}

func (s *Spinner) ui() {
	for {
		select {
		case <-s.ticker.C:
			s.tickCount++
			if s.status != "working" {
				break
			}
			s.redraw()
		}
	}
}

func (s *Spinner) Show() {
	s.ticker = time.NewTicker(s.tickSpeed)
	go s.ui()
}

func (s *Spinner) Fail(message string) {
	s.ticker.Stop()
	s.message = message
	s.status = "fail"
	s.redraw()
	fmt.Println()
}

func (s *Spinner) Success(message string) {
	s.ticker.Stop()
	s.message = message
	s.status = "success"
	s.redraw()
	fmt.Println()
}

func (s *Spinner) Update(message string) {
	s.message = message
	s.redraw()
}
