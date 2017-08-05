package main

import "time"
import "math/rand"
import "bytes"
import "strconv"

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

//card type (num)
const (
ACE = 1
//                       number cards already have correct name
JACK = 11
QUEEN = 12
KING = 13
)

//card suit
const (
SPADES = iota + 1
HEARTS
DIAMONDS
CLUBS
)

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

type card struct {
	num int
	suit int
}

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

type Deck struct {
	cards []card
	r *rand.Rand
}

/*----------------------------------------------------------------------*/

//creates new Deck
func NewDeck () *Deck {
	d := new(Deck)
	d.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return d

}
/*----------------------------------------------------------------------*/

//returns size of Deck
func (d *Deck) size() int {
	return len(d.cards)
}

/*----------------------------------------------------------------------*/

//returns whether or not deck is empty as a boolean
func (d *Deck) empty() bool {
	return d.size() == 0
}

/*----------------------------------------------------------------------*/

//inserts given card at specified index
func (d *Deck) addAt(c card, index int)  {
	d.cards = append (d.cards, card{-1,-1})
	copy(d.cards[index+1:], d.cards[index:(len(d.cards) - 1)])
	d.cards[index] = c
	
}

/*----------------------------------------------------------------------*/

//inserts given card at top of Deck (end of slice)
func (d *Deck) addTop(c card) {
	d.addAt(c, d.size())
}

/*----------------------------------------------------------------------*/

//inserts given card at bottom of Deck (beginning of slice)
func (d *Deck) addBottom(c card) {
	d.addAt(c, 0)
}

/*----------------------------------------------------------------------*/

//inserts random card at specified index
	func (d *Deck) addRandAt(index int)  {
	c := card{d.r.Intn(13) + 1, d.r.Intn(4) + 1}
	d.addAt(c,index)
}

/*----------------------------------------------------------------------*/

//inserts given card at random index
func (d *Deck) addAtRand(c card)  {
	if (d.size() != 0){
		d.addAt(c, d.r.Intn(d.size() + 1))
	} else {	
		d.addTop(c)
	}	
}

/*----------------------------------------------------------------------*/

//inserts random card at random index
func (d *Deck) addRandAtRand()  {
	c := card{d.r.Intn(13) + 1, d.r.Intn(4) + 1}
	if (d.size() == 0){
		d.addTop(c)
	} else {
		d.addAt(c, d.r.Intn(d.size() + 1))
	}
}

/*----------------------------------------------------------------------*/

//returns index of first occurrence of given card in Deck starting from top
//returns -1 if card is not in Deck
func (d *Deck) find(c card) int{
	for i := d.size() - 1; i >= 0; i-- {
		if(d.cards[i] == c){
			return i
		}
	}
	return -1
}

/*----------------------------------------------------------------------*/

//returns index of first occurrence of given card in Deck starting from bottom
//returns -1 if card is not in Deck
func (d *Deck) findBottom(c card) int{
	for i, currentCard := range d.cards {
		if(currentCard == c){
			return i
		}
	}
	return -1
}

/*----------------------------------------------------------------------*/

//removes first occurrence of given card (starting from top) from Deck
//returns card equal to the one removed
//returns card{-1,-1} if given card is not in Deck
func (d *Deck) remove(c card) card {
	index := d.find (c)
	removedCard := d.removeAt(index)	
	return removedCard
}

/*----------------------------------------------------------------------*/

//removes card at given index from Deck
//returns card equal to the one removed
//returns card{-1,-1} if index does not exist
func (d *Deck) removeAt(index int) card {
	if (d.checkIndex(index)){
		removedCard := d.cards[index]
		d.cards = append(d.cards[:index], d.cards[index+1:]...)
		return removedCard
	}
	return card{-1,-1}
}

/*----------------------------------------------------------------------*/

//removes first card from top of Deck
//returns card equal to the one removed
//returns card{-1,-1} if index does not exist
func (d *Deck) removeTop() card {
	return d.removeAt(d.size() - 1)
}

/*----------------------------------------------------------------------*/

//removes first card from bottom of Deck
//returns card equal to the one removed
//returns card{-1,-1} if index does not exist
func (d *Deck) removeBottom() card {
	return d.removeAt(0)
}

/*----------------------------------------------------------------------*/

//removes random card from Deck
//returns card equal to the one removed
//returns card{-1,-1} if Deck is empty
func (d *Deck) removeRand() card {
	if (d.size() != 0){
		randnum := d.r.Intn(d.size())
		return d.removeAt(randnum)
		}
	return card{-1,-1}
}

/*----------------------------------------------------------------------*/

//returns true if index is valid for Deck
//returns false if index is invalid for Deck
//returns false if Deck is empty
func (d *Deck) checkIndex(index int) bool {
	if (d.size() == 0){
		return false
	}
	return (index >= 0 && index < d.size())
}

/*----------------------------------------------------------------------*/

//sorts Deck into order using ad hoc radix sort algorithm
//ordered first by suit and then number
func (d *Deck) sort() {
	s := d.size()
	var cards [13]Deck
	for i := 0; i < s; i++ {
		c := d.removeTop ()
		cards[c.num - 1].addTop(c)
	}
	for i := 0; i < 13; i++ {
		cards[i].appendAtTop (d)
	}
	for i := 0; i < s; i++ {
		c := d.removeBottom ()
		cards[c.suit - 1].addTop(c)
	}
	for i := 0; i < 4; i++ {
		cards[i].appendAtTop (d)
	}
}

//shuffles Deck into new random order
func (d *Deck) shuffle() {
	s := d.size() * 3
	for i := 0; i < s; i++ {
		d.addAtRand (d.removeRand())
	}
}

/*----------------------------------------------------------------------*/

//fills empty Deck with standard Deck of 52 in order
func (d *Deck) fillStandardDeck() {
	if (d.size() == 0){
		for i := 1; i <= 4; i++ {
			for j := 1; j <= 13; j++{
				d.addTop (card{j,i})
			}
		}
	}
}

/*----------------------------------------------------------------------*/

//fills empty Deck with standard Deck of 52 in random order
func (d *Deck) fillStandardDeckShuffled() {
	if (d.size() == 0){
		d.fillStandardDeck()
		d.shuffle()
	}
}

/*----------------------------------------------------------------------*/

//empties Deck
func (d *Deck) clear() {
	if (d.size() > 0){
		d.cards = nil
	}
}

/*----------------------------------------------------------------------*/

//makes this Deck a deep copy of given Deck
func (d *Deck) copy(d2 *Deck) {
	if (d2 != d){
		d.clear()
		d.cards = append(d.cards[:], d2.cards[:]...)
	}
}

/*----------------------------------------------------------------------*/

//adds contents of this Deck to top of given Deck
func (d *Deck) appendAtTop(d2 *Deck) {
	d2.cards = append(d2.cards[:], d.cards[:]...)
	d.clear()
}

/*----------------------------------------------------------------------*/

//adds contents of this Deck to bottom of given Deck
func (d *Deck) appendAtBottom(d2 *Deck) {
	d2.cards = append(d.cards[:], d2.cards[:]...)
	d.clear()
}

/*----------------------------------------------------------------------*/

//adds contents of given Deck to random places in this Deck
func (d *Deck) mixIn(d2 *Deck) {
	for d2.size() != 0 {
		d.addAtRand (d2.removeTop())
	}
}

/*----------------------------------------------------------------------*/

//splits Deck into given number of smaller Decks returned as a slice 
//Number of cards in each returned Deck is equal
//any remainder cards stay in this Deck
//returns an empty slice if an invalid divisor is given
func (d *Deck) split(split int) []*Deck {
	decks := make([]*Deck, split)
	if (split > 0 && split <= d.size()) {
		cards := d.size() / split

		for i := 0; i < split; i++ {
			decks[i] = NewDeck()
			for j := 0; j < cards; j++ {
				decks[i].addBottom(d.removeTop())
			}
		}
	}
	return decks
}

/*----------------------------------------------------------------------*/

//splits Deck evenly into given decks
//Number of cards in each given Deck isequal
//any remainder cards stay in this Deck
//returns a boolean of whether the function had any effect or not
func (d *Deck) splitInto(decks ...*Deck) bool {
	split := len(decks)
	if (split > 0 && split <= d.size()) {
		cards := d.size() / split

		for i := 0; i < split; i++ {
			for j := 0; j < cards; j++ {
				decks[i].addBottom(d.removeTop())
			}
		}
		return true
	}
	return false
}

/*----------------------------------------------------------------------*/

//returns card at given index (starting from top)
//returns card{-1, -1} if index is invalid
func (d *Deck) cardAt(index int) card {
	if (d.checkIndex (index)) {
		return d.cards[d.size() - index]
	}
	return card {-1, -1}
}

/*----------------------------------------------------------------------*/

//converts Deck to a string
func (d *Deck) toString() string {
	if (d == nil) {
		return "";
	}
	j := d.size()
	var buffer bytes.Buffer
	for i := 0 ; i < j; i++ {
		buffer.WriteString(d.cards[i].toString())
		buffer.WriteString(" | ")
	}
	return buffer.String()
}

/*----------------------------------------------------------------------*/

//converts Card to string
func (c *card) toString() string {
	if (c == nil) {
		return "";
	}
	var buffer bytes.Buffer

	buffer.WriteString(strconv.Itoa(c.num))
	if (c.num < 10){
		buffer.WriteString(" ")
	}
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(c.suit))
	return buffer.String()
}

/*----------------------------------------------------------------------*/

//converts Deck to a string with words for card values
func (d *Deck) toStringWords() string {
	if (d == nil) {
		return "";
	}
	j := d.size()
	var buffer bytes.Buffer
	for i := 0 ; i < j; i++ {
		buffer.WriteString(d.cards[i].toStringWords())
		buffer.WriteString(" | ")
	}
	return buffer.String()
}

/*----------------------------------------------------------------------*/

//converts Card to string with words for its values
func (c *card) toStringWords() string {
	if (c == nil) {
		return "";
	}
	var buffer bytes.Buffer
	var num string
	var suit string
	switch c.num {
		case ACE   : num = " Ace "
		case KING  : num = " King"
		case QUEEN : num = "Queen"
		case JACK  : num = " Jack"
		case 10    : num = "  10 "
		default    : num = "  " + strconv.Itoa(c.num) + "  "
	}

	switch c.suit {
		case SPADES   : suit = " Spades "
		case HEARTS   : suit = " Hearts "
		case DIAMONDS : suit = "Diamonds"
		case CLUBS    : suit = " Clubs  "
	}

	buffer.WriteString(num + " of " + suit)
	return buffer.String()
}

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/