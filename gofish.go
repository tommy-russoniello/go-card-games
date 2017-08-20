package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

type player struct {
  deck *Deck
  name string
	books *Deck
}

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

var player1, player2 player
var pile *Deck
var current, other *player

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

func main () {
  initialize()
	count := 1
		for i := 0; i < 10; i++{
			fmt.Println("turn ", count, "...")
			fmt.Print(current.name, ": ")
			fmt.Println(current.deck.toStringWords())
			printRequestableCards(current)
			gofish := makeRequest()
			if (checkMatches()) {
				break;
			}
			if (gofish) {
				count++
				switchPlayer()
			}
		}
  end()
}

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

func initialize() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Player 1 name: ")
	player1.name, _ = reader.ReadString('\n')
	fmt.Print("Player 2 name: ")
	player2.name, _ = reader.ReadString('\n')

	player1.name = strings.TrimRight(player1.name, "\r\n")
	player2.name = strings.TrimRight(player2.name, "\r\n")

	pile = NewDeck()
	player1.deck = NewDeck()
	player2.deck = NewDeck()

	pile.fillStandardDeckShuffled()
	player1.deck.addTop(pile.drawTop(7)...)
	player2.deck.addTop(pile.drawTop(7)...)

	current = &player1
	other = &player2
}

/*----------------------------------------------------------------------*/

func switchPlayer () {
	current, other = other, current
}

/*----------------------------------------------------------------------*/

func printRequestableCards (player *player) {
	// TODO
}

/*----------------------------------------------------------------------*/

func makeRequest () bool {
	return false
	// TODO
}

/*----------------------------------------------------------------------*/

func checkMatches () bool {
	return false
	// TODO
}

/*----------------------------------------------------------------------*/

func end () {
  os.Exit(0)
}

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
