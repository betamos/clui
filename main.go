package main

import (
	"fmt"
	//"github.com/maxmclau/gput"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type symbol struct {
	color color.Attribute
	sym   string
}

type Theme []string

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
		theme:     []string{"◝", "◞", "◟", "◜"},
		tickSpeed: 200 * time.Millisecond,
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
		fmt.Print(s.theme[s.tickCount%len(s.theme)])
	}
	fmt.Printf(" %v", s.message)
	s.mtx.Unlock()

}

func (s *Spinner) ui() {
	s.ticker = time.NewTicker(s.tickSpeed)
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

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		signal.Stop(c)
		fmt.Print("\u001b[2K") // clear line
		fmt.Print("\u001b[0G") // to column 0
		color.Set(color.FgRed)
		fmt.Print("✘")
		color.Unset()
		fmt.Println(" file copy aborted")
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(sig)

	}()
	s := NewSpinner("copying files")
	s.Show()
	time.Sleep(time.Second)
	s.Update("uploading directory")
	time.Sleep(time.Second)
	s.Fail("failed upload")
	//s.Success("uploaded 122 files")
	fmt.Println("other output")
	time.Sleep(time.Second)
}
