package main

import (
	"fmt"
	"time"
)

type Output struct {
	W  time.Duration
	T  string
	N1 string
	N2 string
	P1 []string
	P2 []string
}

func (o *Output) Println() {
	fmt.Printf("W=%s, T=%s, N1=%s, P1=%s, N2=%s, P2=%s",
		o.W,
		o.T,
		o.N1,
		o.P1,
		o.N2,
		o.P2,
	)
}
