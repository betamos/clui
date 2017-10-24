package main

	// symbol{color.FgRed, "◴ ◷ ◶ ◵"},
	syms := []symbol{
		symbol{color.FgRed, "◷"},
		symbol{color.FgRed, "◶"},
		symbol{color.FgRed, "◵"},
		symbol{color.FgRed, "◴"},
	}
	syms = []symbol{
		symbol{color.FgRed, "G"},
		symbol{color.FgYellow, "G"},
		symbol{color.FgBlue, "G"},
		symbol{color.FgGreen, "G"},
	}
	syms = []symbol{
		symbol{color.FgYellow, "◝"},
		symbol{color.FgYellow, "◞"},
		symbol{color.FgYellow, "◟"},
		symbol{color.FgYellow, "◜"},
	}
	syms = []symbol{
		symbol{color.FgYellow, "⠛"},
		symbol{color.FgYellow, "⠹"},
		symbol{color.FgYellow, "⢸"},
		symbol{color.FgYellow, "⣰"},
		symbol{color.FgYellow, "⣤"},
		symbol{color.FgYellow, "⣆"},
		symbol{color.FgYellow, "⡇"},
		symbol{color.FgYellow, "⠏"},
	}
	// ⣏⣿⣉⣹
	//syms = []symbol{
	//	symbol{color.FgYellow, "◐"},
	//	symbol{color.FgYellow, "◓"},
	//	symbol{color.FgYellow, "◑"},
	//	symbol{color.FgYellow, "◒"},
	//}

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
	for i := 0; i < 100; i++ {
		if i == 21 {
			fmt.Fprintf(os.Stderr, "herro\n")
		}
		sym := syms[i%len(syms)]
		fmt.Print("\u001b[2K") // clear line
		fmt.Print("\u001b[0G") // to column 0
		// color.Set(sym.color)
		fmt.Printf("%v", sym.sym)
		// color.Unset()
		fmt.Printf(" copying files %v/100", i)
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Print("\u001b[2K") // clear line
	fmt.Print("\u001b[0G") // to column 0
	color.Set(color.FgGreen)
	fmt.Print("✔")
	color.Unset()
	fmt.Println(" copied 100 files")
