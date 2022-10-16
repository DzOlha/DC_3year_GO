package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	ActionsAllowedPerThread = 20
)

func boolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func gardener(garden [][]bool, rwMutex *sync.RWMutex, exit chan int) {
	for i := 0; i < ActionsAllowedPerThread; i++ {
		rwMutex.Lock()
		for i := 0; i < len(garden); i++ {
			for j := 0; j < len(garden[0]); j++ {
				if garden[i][j] == false {
					garden[i][j] = true
				}
			}
		}
		rwMutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
	exit <- 1
}

func nature(garden [][]bool, rwMutex *sync.RWMutex, exit chan int) {
	for i := 0; i < ActionsAllowedPerThread; i++ {
		rwMutex.Lock()
		for i := 0; i < len(garden)*2; i++ {
			index1 := random.Intn(len(garden))
			index2 := random.Intn(len(garden[0]))
			garden[index1][index2] = !garden[index1][index2]
		}
		rwMutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
	exit <- 1
}

func monitor1(garden [][]bool, rwMutex *sync.RWMutex, exit chan int) {
	file, err := os.Create("gardenStatus.txt")

	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}

	defer file.Close()

	for i := 0; i < ActionsAllowedPerThread; i++ {
		rwMutex.RLock()
		for i := 0; i < len(garden); i++ {
			var line string
			for j := 0; j < len(garden[0]); j++ {
				line += boolToString(garden[i][j]) + " "
			}
			_, _ = file.WriteString(line + "\n")
		}
		rwMutex.RUnlock()
		_, _ = file.WriteString("\n\n")
		time.Sleep(500 * time.Millisecond)
	}
	exit <- 1
}

func monitor2(garden [][]bool, rwMutex *sync.RWMutex, exit chan int) {
	for i := 0; i < ActionsAllowedPerThread; i++ {
		rwMutex.RLock()
		for i := 0; i < len(garden); i++ {
			var line string
			for j := 0; j < len(garden[0]); j++ {
				line += boolToString(garden[i][j]) + " "
			}
			fmt.Println(line)
		}
		rwMutex.RUnlock()
		fmt.Println()
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}
	exit <- 1
}

func main() {
	var garden [][]bool
	var rwMutex sync.RWMutex
	exit := make(chan int, 4)

	for i := 0; i < 10; i++ {
		var row []bool
		for j := 0; j < 10; j++ {
			random = rand.New(rand.NewSource(time.Now().UnixNano()))
			row = append(row, random.Intn(2) != 0)
		}
		garden = append(garden, row)
	}

	go monitor2(garden, &rwMutex, exit)
	go monitor1(garden, &rwMutex, exit)
	go nature(garden, &rwMutex, exit)
	go gardener(garden, &rwMutex, exit)

	for i := 0; i < 4; i++ {
		<-exit
	}
}
