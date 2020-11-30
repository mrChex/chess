package main

import (
	"github.com/mrChex/howToWinInChess/battlefield/models"
)

func main() {
	battlefield := models.NewBattlefield()

	performturn := func (from,to string) {
		if err := battlefield.Move(from, to); err != nil {
			panic(err)
		}
		battlefield.Print()
	}

	performturn("e2", "e4")
	performturn("e7", "e5")
	performturn("e4", "f5")
}

