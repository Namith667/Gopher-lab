package main

import "fmt"

//import "fmt"
func main(){
	cards := newDeck()
	// hand,remainingCards := deal(cards,5)
	// hand.print()
	// remainingCards.print()
	
	//conversion of string to byte slice
	// arrstr:=deck()
	// bslice:=[]byte(str)
	// fmt.Println(bslice)
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	

}
