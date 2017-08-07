package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

var trash, deck1, deck2 *Deck
var player1, player2 string
var count int

func main(){
	initialize()

	for {
		fmt.Println("turn ", count, "...")
		card1, card2 := drawCards(3)

		roundWinner, roundLoser, winnings := compare(card1, card2)
		winnings.shuffle()
		winnings.appendAtBottom(roundWinner)

		fmt.Print(getPlayer(roundWinner), " wins that round!\n")
		if (roundLoser.empty()) {
			end(roundWinner)
		}
		count++
	}

}

/* 
 * take in player names
 * initalize decks and turn count
 */
func initialize() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Player 1 name: ")
	player1, _ = reader.ReadString('\n')
	fmt.Print("Player 2 name: ")
	player2, _ = reader.ReadString('\n')

	player1 = strings.TrimRight(player1, "\r\n")
	player2 = strings.TrimRight(player2, "\r\n")

	trash = NewDeck()
	deck1 = NewDeck()
	deck2 = NewDeck()

	trash.fillStandardDeckShuffled()
	trash.splitInto(deck1, deck2)

	count = 1
}

/* 
* players -> players to be affected
* *** 1 -> player 1
* *** 2 -> player 2
* *** 3 -> both players
* cardTemp -> optional default card for unaffected cards to be set to
*
* draws a card from affected players' decks and returns those cards
* returns cardTemp[0] card for unaffected players
* waits for user to press 'enter' and then prints players' cards
*/
func drawCards(players int, cardTemp ...card) (card, card) {
	var card1, card2 card
	if players != 2 {
		card1 = deck1.removeTop()
		if players == 1  && cardTemp != nil {
			card2 = cardTemp[0]
		}
	}
	if players != 1 {
		card2 = deck2.removeTop()
		if players == 2 && cardTemp != nil {
			card1 = cardTemp[0]
		}
	}

	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Print(player1, "[", deck1.size(), " card(s)] draws: ", card1.toStringWords(), "\n")
	fmt.Print(player2, "[", deck2.size(), " card(s)] draws: ", card2.toStringWords(), "\n")
	return card1, card2
}

/*
 * card1 -> player 1's card
 * card2 -> player 2's card
 * winningsTemp -> optional cumulative winnings deck
 * *** used for recursive calls from chained WAR sequences
 *
 * adds players cards to winnings deck (may be empty to start)
 * compares cards' values and returns all decks in this order:
 * *** 1) deck of the player who won the round
 * *** 2) deck of the player who lost the round
 * *** 3) deck of all cards won by the winning player in this round
 * triggers war if players' cards are equal in value and returns the results
 */
func compare(card1 card, card2 card, winningsTemp ...*Deck) (*Deck, *Deck, *Deck) {
	var winnings *Deck
	if winningsTemp == nil {
		winnings = NewDeck()
	} else {
		winnings = winningsTemp[0]
	}
	winnings.addTop(card2, card1)

	if card1.num == card2.num {
		return thisIsWar(card1, winnings)
	} else if card1.num == ACE {
		if card2.num == 2 {
			return deck2, deck1, winnings
		} else {
			return deck1, deck2, winnings
		}
	} else if card2.num == ACE {
		if card1.num == 2 {
			return deck1, deck2, winnings
		} else {
			return deck2, deck1, winnings
		}
	} else if card1.num > card2.num {
		return deck1, deck2, winnings
	} else {
		return deck2, deck1, winnings
	}
}

/*
 * c -> card with value that both players drew
 * winnings -> cumulative winnings for round
 *
 * prints out situation
 * draws up to 3 cards from each deck and adds them to the winnings
 * *** drawing will stop if player only has 1 card left in their deck
 * draws one more card from each deck and returns the result of their comparison
 */
func thisIsWar(c card, winnings *Deck) (*Deck, *Deck, *Deck) {
	cardType := c.toStringWords()
	fmt.Print("both have a(n)", cardType[:len(cardType) - 12], "!\n THIS IS WAR -> ")

	for i := 0; i < 3; i++ {
		if deck1.size() <= 1 {
			break
		} else {
			winnings.addTop(deck1.removeTop())
		}
	}
	for i := 0; i < 3; i++ {
		if deck2.size() <= 1 {
			break
		} else {
			winnings.addTop(deck2.removeTop())
		}
	}

	var card1, card2 card
	if deck1.empty() && !deck2.empty() {
		fmt.Print(winnings.size() + 1, " cards at stake...\n")
		card1, card2 = drawCards(2, c)
	} else if deck2.empty() && !deck1.empty() {
		fmt.Print(winnings.size() + 1, " cards at stake...\n")
		card1, card2 = drawCards(1, c)
	} else {
		fmt.Print(winnings.size() + 2, " cards at stake...\n")
		card1, card2 = drawCards(3)
	}

	return compare(card1, card2, winnings) 
}

/*
 * player -> string of winning player's name
 *
 * prints out victory message for winning player
 * exits program
 */
func end(winnerDeck *Deck) {
	player := getPlayer(winnerDeck)
	fmt.Println(player, "has all ", winnerDeck.size(), " cards! ", player, " wins!")
	os.Exit(0)
}

/*
 * d -> deck of either player
 *
 * returns name of player for given deck
 */
func getPlayer(d *Deck) string {
	if d == deck1 {
		return player1
	} else if d == deck2 {
		return player2
	} else {
		return "oh, shit"
	}
}
