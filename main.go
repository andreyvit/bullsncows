package main

import (
	"bytes"
	"fmt"
	"log"
)

const size = 6

const digits = 10

var alphabet = []byte("0123456789")

type Comb [size]byte

func Parse(s string) Comb {
	var c Comb
	bb := []byte(s)
	if len(bb) != size {
		panic("wrong number of digits")
	}
	for i, b := range bb {
		d := bytes.IndexByte(alphabet, b)
		if d < 0 {
			panic("unknown digit")
		}
		c[i] = byte(d)
	}
	return c
}

func (c Comb) String() string {
	var s [size]byte
	for i, d := range c {
		s[i] = alphabet[d]
	}
	return string(s[:])
}

type Result struct {
	Bulls int
	Cows  int
}

func (r Result) Equal(o Result) bool {
	return r.Bulls == o.Bulls && r.Cows == o.Cows
}

func (r Result) String() string {
	if r.Bulls > 0 {
		if r.Cows > 0 {
			return fmt.Sprintf("%dБ %dК", r.Bulls, r.Cows)
		} else {
			return fmt.Sprintf("%dБ", r.Bulls)
		}
	} else {
		if r.Cows > 0 {
			return fmt.Sprintf("%dК", r.Cows)
		} else {
			return "0"
		}
	}
}

type Matches uint

func (m Matches) String() string {
	var b [3]rune
	for i := uint(0); i < 3; i++ {
		if (m & (1 << i)) != 0 {
			b[i] = '+'
		} else {
			b[i] = '-'
		}
	}
	return string(b[:])
}

func Match(move *Comb, orig *Comb, res Result) Matches {
	var r Matches
	if Compute1(move, orig).Equal(res) {
		r |= 1
	}
	if Compute2(move, orig).Equal(res) {
		r |= 2
	}
	if Compute3(move, orig).Equal(res) {
		r |= 4
	}
	return r
}

func Compute1(move *Comb, orig *Comb) Result {
	var result Result
	for mi, md := range *move {
		for oi, od := range *orig {
			if md == od {
				if mi == oi {
					result.Bulls++
				} else {
					result.Cows++
				}
			}
		}
	}
	return result
}

func Compute2(move *Comb, orig *Comb) Result {
	var result Result
	// var seen [digits]bool
	for mi, md := range *move {
		var bulls, cows int
		for oi, od := range *orig {
			if md == od {
				if mi == oi {
					bulls++
				} else {
					cows++
				}
			}
		}
		if bulls > 0 {
			result.Bulls++
		} else if cows > 0 {
			result.Cows++
		}
	}
	return result
}

func Compute3(move *Comb, orig *Comb) Result {
	var result Result
	// var seen [digits]bool
	for oi, od := range *orig {
		var bulls, cows int
		for mi, md := range *move {
			if md == od {
				if mi == oi {
					bulls++
				} else {
					cows++
				}
			}
		}
		if bulls > 0 {
			result.Bulls++
		} else if cows > 0 {
			result.Cows++
		}
	}
	return result
}

type Round struct {
	Move *Comb
	Result
}

var rounds = []*Round{
	&Round{&Comb{1, 2, 3, 4, 5, 6}, Result{0, 2}},
	&Round{&Comb{7, 8, 9, 0, 2, 3}, Result{0, 4}},
	&Round{&Comb{8, 7, 5, 3, 1, 0}, Result{2, 1}},
	&Round{&Comb{6, 0, 5, 9, 8, 0}, Result{0, 5}},
	&Round{&Comb{8, 7, 0, 6, 0, 4}, Result{2, 1}},
	&Round{&Comb{8, 9, 8, 6, 1, 9}, Result{3, 1}},
	&Round{&Comb{8, 7, 6, 5, 9, 9}, Result{3, 1}},
}

func possible(move *Comb) bool {
	for _, round := range rounds {
		if Match(round.Move, move, round.Result) == 0 {
			return false
		}
	}
	return true
}

type Finder struct {
	next Comb

	evaluated int
}

func (f *Finder) find(i int) {
	if i == size {
		f.evaluated++
		if possible(&f.next) {
			log.Printf("%v", f.next)
			for _, round := range rounds {
				m := Match(round.Move, &f.next, round.Result)
				log.Printf("    %v -> %v", round.Move, m)
			}
		}
		return
	}
	for d := byte(0); d < digits; d++ {
		f.next[i] = d
		f.find(i + 1)
	}
}

func main() {
	var f Finder
	f.find(0)
	log.Printf("done: %d evaluated", f.evaluated)
}
