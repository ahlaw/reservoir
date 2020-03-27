package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func reservoirSample(r io.Reader, w io.Writer, capacity int) {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(r)
	reservoir := make([]string, 0, capacity)
	i := 0
	next := capacity
	fcapacity := float64(capacity)

	W := math.Exp(math.Log(rand.Float64()) / fcapacity)

	for scanner.Scan() {
		if len(reservoir) < cap(reservoir) {
			data := scanner.Text()
			reservoir = append(reservoir, data)
		} else if i == next {
			data := scanner.Text()
			reservoir[rand.Intn(len(reservoir))] = data
			next += int(math.Log(rand.Float64())/math.Log(1-W)) + 1
			W *= math.Exp(math.Log(rand.Float64()) / fcapacity)
		}
		i++
	}
	for _, item := range reservoir {
		fmt.Fprintln(w, item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <capacity>\n", os.Args[0])
		os.Exit(1)
	}

	capacity, err := strconv.Atoi(os.Args[1])
	if err != nil || capacity < 0 {
		fmt.Printf("Error: %v is an invalid capacity size\n", os.Args[1])
		os.Exit(1)
	}

	reservoirSample(os.Stdin, os.Stdout, capacity)
}
