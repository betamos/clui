package main

import (
	"fmt"
	"github.com/betamos/clui"
	"time"
)

func main() {
	fmt.Println("vim-go")
	a := clui.NewTask("trying something")
	time.Sleep(3 * time.Second)
	b := clui.NewTask("something else")
	time.Sleep(3 * time.Second)
	a.Fail("aborting")
	time.Sleep(3 * time.Second)
	b.Success("at least I made it")

}
