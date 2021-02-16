package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"test-loyalto-2/helper"

)

var dice = []int{1, 2, 3, 4, 5, 6}

var diceFromPreviousPlayer = 0
var nextPlayer = -1
var existingPlayer = 0

type Players struct {
	player        int
	remainingDice int
	score         int
	diceHistory   []int
	complete      bool
	queueDice     int
}


var allPlayer = make([]Players, 0)


func checkNextPlayer(index int) int {
	var existingPlayer []int
	for i, _ := range allPlayer {
		if !allPlayer[i].complete {
			existingPlayer = append(existingPlayer, i)
		}
	}
	if len(existingPlayer) > 1 {
		result := helper.FindIndexByValue(existingPlayer, index)
		if result == (len(existingPlayer) - 1) {
			nextPlayer = existingPlayer[0]
		} else {
			nextPlayer = existingPlayer[result+1]
		}
	}
	return nextPlayer
}

func checkStatusPlayer() {
	for i, _ := range allPlayer {
		if allPlayer[i].remainingDice <= 0 {
			allPlayer[i].complete = true
		}
	}
}

func checkUnCompletePlayer() int {
	result := 0
	for i, _ := range allPlayer {
		if !allPlayer[i].complete {
			result++
		}
	}
	return result
}


func rule(arr []int, index int) {
	myRealDice := allPlayer[index].remainingDice - allPlayer[index].queueDice
	for i, num := range arr {
		if num == 6 {
			allPlayer[index].score = allPlayer[index].score + 1
			allPlayer[index].remainingDice = allPlayer[index].remainingDice - 1
		} else if num == 1 {
			allPlayer[index].remainingDice = allPlayer[index].remainingDice - 1
			allPlayer[index].diceHistory = helper.RemoveElementByValueLimit(allPlayer[index].diceHistory, 1, 1)
			diceFromPreviousPlayer++
		}

		if myRealDice == i+1 {
			allPlayer[index].queueDice = 0
			break
		}
	}
	allPlayer[index].diceHistory = helper.RemoveElementByValue(allPlayer[index].diceHistory, 6)
	if diceFromPreviousPlayer > 0 && nextPlayer > -1 {
		for i := 0; i < diceFromPreviousPlayer; i++ {
			allPlayer[nextPlayer].diceHistory = append(allPlayer[nextPlayer].diceHistory, 1)
			allPlayer[nextPlayer].remainingDice = allPlayer[nextPlayer].remainingDice + 1
			allPlayer[nextPlayer].queueDice = allPlayer[nextPlayer].queueDice + 1
		}
		diceFromPreviousPlayer = 0
	}
}


func initializePlayer(player int, totalDice int) {
	for i := 1; i <= player; i++ {
		allPlayer = append(allPlayer, Players{i, totalDice, 0, nil, false, 0})
	}
}

func clearAndEvalute() {
	for index, _ := range allPlayer {
		if allPlayer[index].remainingDice <= 0 {
			fmt.Printf("Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu) \n", allPlayer[index].player, allPlayer[index].score)
		} else {
			fmt.Printf("Pemain #%d (%d): %d\n", allPlayer[index].player, allPlayer[index].score, allPlayer[index].diceHistory)
		}
		allPlayer[index].diceHistory = nil
	}
}

func evaluate() {
	fmt.Println("Setelah Evaluasi:  ")
	for index, _ := range allPlayer {
		if !allPlayer[index].complete {
			nextPlayer = checkNextPlayer(index)
			rule(allPlayer[index].diceHistory, index)
		}
	}
	clearAndEvalute()
	checkStatusPlayer()
}

func startGame() {
	var round = 1
	for existingPlayer > 1 {
		fmt.Println("Giliran ", round, " Lempar Dadu")
		for index, _ := range allPlayer {
			for i := 0; i < allPlayer[index].remainingDice; i++ {
				allPlayer[index].diceHistory = append(allPlayer[index].diceHistory, helper.RollDice(dice))
			}
			if allPlayer[index].remainingDice <= 0 {
				fmt.Printf("Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu) \n", allPlayer[index].player, allPlayer[index].score)
			} else {
				fmt.Printf("Pemain #%d (%d): %d\n", allPlayer[index].player, allPlayer[index].score, allPlayer[index].diceHistory)
			}
		}
		evaluate()
		fmt.Println("=====================================================================")
		existingPlayer = checkUnCompletePlayer()
		round++
	}
}

func FindMaxAndMin(players []Players) (min Players, max Players) {
	min = players[0]
	max = players[0]
	for _, player := range players {
		if player.score > max.score {
			max = player
		}
		if player.score < min.score {
			min = player
		}
	}
	return min, max
}

func FindEndPlayer(players []Players) (endPlayer Players) {
	endPlayer = players[0]

	for _, player := range players {
		if !player.complete  {
			endPlayer = player
		}

	}
	return endPlayer
}



func main() {
    var diceQty int
	var player int
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Masukkan Jumlah Pemain = ")
    scanner.Scan()
   	playerString := scanner.Text()
   	player, _ = strconv.Atoi(playerString)

	fmt.Print("Masukkan Jumlah Dadu = ")
	scanner.Scan()
	diceQtyString := scanner.Text()
	diceQty, _ = strconv.Atoi(diceQtyString)

	fmt.Printf("\nPemain = %d, Dadu = %d ",player , diceQty)
	fmt.Println("=====================================================================")
	if player == 0 && diceQty == 0 {
		fmt.Print("Jumlah Pemain Atau Dadu Belum Di Masukkan")
	}else{
		initializePlayer(player, diceQty)
		existingPlayer = checkUnCompletePlayer()
		startGame()
		_, max := FindMaxAndMin(allPlayer)
		endPlayer := FindEndPlayer(allPlayer)

		fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.",endPlayer.player )
		fmt.Printf("\nGame dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya dan menang lebih dulu.", max.player)
	}
}