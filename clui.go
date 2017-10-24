package clui

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
	"time"
)

type Task struct {
	theme   []string
	status  string // enum?
	message string
	mtx     *sync.Mutex
}

var (
	tasks     []*Task
	mtx       *sync.Mutex = &sync.Mutex{}
	ticker    *time.Ticker
	tickSpeed time.Duration = 150 * time.Millisecond
	tickCount int
)

func NewTask(message string) *Task {
	s := &Task{
		theme:   []string{"⠛", "⠹", "⢸", "⣰", "⣤", "⣆", "⡇", "⠏"},
		status:  "working",
		message: message,
		mtx:     &sync.Mutex{},
	}
	tasks = append(tasks, s)
	fmt.Println()
	go ui()
	return s
}

func redraw() {
	mtx.Lock()

	fmt.Print("\u001b[0G")               // to column 0
	fmt.Printf("\u001b[%vA", len(tasks)) // to row -n
	for _, t := range tasks {
		fmt.Print("\u001b[2K") // clear line
		fmt.Print("\u001b[0G") // to column 0
		if t.status == "success" {
			color.Set(color.FgGreen)
			fmt.Print("✔")
			color.Unset()
		} else if t.status == "fail" {
			color.Set(color.FgRed)
			fmt.Print("✘")
			color.Unset()

		} else {
			color.Set(color.FgYellow)
			fmt.Print(t.theme[tickCount%len(t.theme)])
			color.Unset()
		}
		fmt.Printf(" %v\n", t.message)
	}
	mtx.Unlock()

}

func ui() {
	if ticker != nil {
		return
	}
	ticker = time.NewTicker(tickSpeed)
	for range ticker.C {
		tickCount++
		redraw()
	}
}

func (s *Task) Fail(message string) {
	s.complete(message, "fail")
}
func (s *Task) Success(message string) {
	s.complete(message, "success")
}

func (s *Task) complete(message, status string) {
	//s.ticker.Stop()
	s.message = message
	s.status = status
	redraw()
}

func (s *Task) Update(message string) {
	s.message = message
	redraw()
}
