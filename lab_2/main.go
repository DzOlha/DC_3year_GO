/*
с) Завдання про Шлях Кулака. На сивих схилах Гімалаїв стоять два древніх буддистських монастиря:

	Гуань-Інь і Гуань-Янь. Щороку в день зішестя на землю боддісатви Араватті ченці обох монастирів
	 збираються на спільне свято і показують своє вдосконалення на Шляху Кулака.
	 Всіх тих, що змагаються ченців розбивають на пари,
	 переможці пар б'ються потім між собою і так далі, до фінального поєдинку.
	 Монастир, ченці якого перемогли у фінальному бою,
	 забирає собі на зберігання статую боддісатви. Реалізувати багатопоточний додаток,
	 що визначає переможця. В якості вхідних даних використовується масив,
	 в якому зберігається кількість енергії Ци кожного ченця.
	 При вирішенні використовувати принцип дихотомії.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const SIZE = 16

type Player struct {
	name      string
	energy    int
	monastery string
}

func fillPlayersArray() [SIZE]Player {
	//M1 - first monastery
	//M2 - second monastery
	var monTmp string = "M1"

	arr := [SIZE]Player{}

	//set seed as timeNow to generate every time new values
	rand.Seed(time.Now().UnixMicro())
	fmt.Println("-----------------Initial Players------------------")
	for i := 0; i < SIZE; i++ {
		//the first part of the players is from the FIRST monastery
		//the second part of the players is from the SECOND monastery
		if i >= SIZE/2 {
			monTmp = "M2"
		}
		arr[i].name = fmt.Sprint("player_", i)
		arr[i].energy = rand.Intn(100) + 1
		arr[i].monastery = monTmp

		fmt.Println(arr[i])
	}
	fmt.Println("-----------------Initial Players------------------")
	return arr
}

func oneCompetition(one Player, two Player, c chan Player) {
	if one.energy > two.energy {
		one.energy += two.energy
		two.energy = 0
		c <- one
	} else {
		two.energy += one.energy
		one.energy = 0
		c <- two
	}
}

func main() {
	players := fillPlayersArray()
	var halfSize int = SIZE / 2
	var size int = SIZE
	c := make(chan Player, halfSize)

	for {
		for i := 0; i < halfSize; i++ {
			go oneCompetition(players[i], players[size-i-1], c)
		}
		for i := 0; i < halfSize; i++ {
			players[i] = <-c
			fmt.Println(players[i])
			if i == halfSize-1 {
				fmt.Println(fmt.Sprint("---------------------Above are Winners: Round 1/", size/2, "-------------------------"))
			}
		}
		size = halfSize
		if size == 1 {
			break
		}
		halfSize = halfSize / 2
	}

}
