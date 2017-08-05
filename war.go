package main
import "fmt"
import "os"
import "bufio"
import "strings"

var trash, deck1, deck2 *Deck
var player1, player2 string
var count int

func main(){
	initialize()
	for {
		fmt.Println("turn ", count, "...")
		var card1, card2 card
		card1, card2 = drawCards()
		fmt.Print(player1, "[", deck1.size(), " cards] draws: ", card1.toStringWords(), "\n")
		fmt.Print(player2, "[", deck2.size(), " cards] draws: ", card2.toStringWords(), "\n")
		var roundWinner, roundLoser, winnings *Deck
		roundWinner, roundLoser, winnings = compare(card1, card2)
		winnings.appendAtBottom(roundWinner)
		fmt.Print(getPlayer(roundWinner), " wins!\n")
		if (roundLoser.empty()) {
			end(roundWinner)
		}
		count++
	}

}

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

func drawCards() (card, card) {
	card1 := deck1.removeTop()
	card2 := deck2.removeTop()
	return card1, card2
}

func compare(card1 card, card2 card, winningsTemp ...*Deck) (*Deck, *Deck, *Deck) {
	if len(winningsTemp) == 0 {
		winnings := NewDeck()
	} else {
		winnings := winningsTemp[0]
	}
	winnings.addTop(card2)
	winnings.addTop(card1)
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

func thisIsWar(c card, winnings *Deck) (*Deck, *Deck, *Deck) {
	fmt.Print("both have", c.toStringWords(), "!\n THIS IS WAR\n")
	for i := 0; i < 3; i++ {
		if deck1.size() == 1 {
			break
		} else {
			winnings.addTop(deck1.removeTop())
		}
	}
	for i := 0; i < 3; i++ {
		if deck2.size() == 1 {
			break
		} else {
			winnings.addTop(deck2.removeTop())
		}
	}
	var card1, card2 card
	card1, card2 = drawCards()
	fmt.Print(player1, "[", deck1.size(), " cards] draws: ", card1.toStringWords(), "\n")
	fmt.Print(player2, "[", deck2.size(), " cards] draws: ", card2.toStringWords(), "\n")
	var winner, loser *Deck
	winner, loser = compare(card1, card2)
	return winner, loser, winnings
}

func end(d *Deck) {
	fmt.Println("blah blah blah player whatever wins")
	os.Exit(0)
}

func getPlayer(d *Deck) string {
	if d == deck1 {
		return player1
	} else if d == deck2 {
		return player2
	} else {
		return "oh, shit"
	}
}

func print(d *Deck) {
	var size int
	if (d == nil) {
		size = 0
	} else {
		size = d.size()
	}
	fmt.Println(d.toStringWords(), " size: ", size, "\n")
}
