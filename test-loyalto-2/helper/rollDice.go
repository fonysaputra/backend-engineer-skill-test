package helper

import (
	
	"math/rand"
	"sync"
	"time"
)

var onlyOnce sync.Once

func RollDice(dice []int) int {
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	return dice[rand.Intn(len(dice))]
}