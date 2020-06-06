package main

import (
	"bufio"
	"fmt"
	"os"
)

type Process struct {
	name string
	time int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, q int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &q)

	procs := make([]Process, n)
	for i := 0; i < n; i++ {
		var name string
		var time int
		fmt.Fscan(r, &name)
		fmt.Fscan(r, &time)
		procs[i] = Process{name, time}
	}

	var p Process
	elapsedtime := 0
	for len(procs) > 0 {
		p, procs = procs[0], procs[1:]
		if p.time > q {
			elapsedtime += q
			p.time -= q
			procs = append(procs, p)
		} else {
			elapsedtime += p.time
			fmt.Fprintf(w, "%s %d\n", p.name, elapsedtime)
		}
	}
}
