package main

import (
	"fmt"
	"math/rand"
	"time"
)

const SIZE = 10

type Player struct {
	name      string
	energy    int
	monastery bool
}

func fillPlayersArray() [SIZE]Player {
	//false - first monastery
	//true - second monastery
	var monTmp bool = false

	arr := [SIZE]Player{}

	//set seed as timeNow to generate every time new values
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < SIZE; i++ {
		//the first part of the players is from the FIRST monastery
		//the second part of the players is from the SECOND monastery
		if i >= SIZE/2 {
			monTmp = true
		}
		arr[i].name = fmt.Sprint("player_", i)
		arr[i].energy = rand.Intn(100) + 1
		arr[i].monastery = monTmp

		fmt.Println(arr[i])
	}
	return arr
}
func main() {
	const size = 10
	fillPlayersArray()

}
