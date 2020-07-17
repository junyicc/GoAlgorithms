package main

import (
	"fmt"
	"time"
)

// use three processor to print A, B, C
// the result must be "ABCABCABC..."

func printA(waitC chan struct{}, waitA chan struct{}) {
	for {
		select {
		case <-waitC:
			fmt.Print("A")
			waitA <- struct{}{}
		}
	}
}

func printB(waitA chan struct{}, waitB chan struct{}) {
	for {
		select {
		case <-waitA:
			fmt.Print("B")
			waitB <- struct{}{}
		}
	}
}

func printC(waitB chan struct{}, waitC chan struct{}) {
	for {
		select {
		case <-waitB:
			fmt.Print("C")
			waitC <- struct{}{}
		}
	}
}

func coordinator() {
	waitA := make(chan struct{})
	waitB := make(chan struct{})
	waitC := make(chan struct{})

	go printC(waitB, waitC)
	go printB(waitA, waitB)
	go printA(waitC, waitA)

	waitC <- struct{}{}
	time.Sleep(5 * time.Second)
}

func main() {
	coordinator()
}
