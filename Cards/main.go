package main



//import "fmt"
func main(){
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.toString()
	cards.print()
	

	//a := [...]int{1, 3, 4, 5, 6, 6}
	
	//fmt.Println(a[rand.Intn(len(a))])

	// hand,remainingCards := deal(cards,5)
	// hand.print()
	// remainingCards.print()
	
	//conversion of string to byte slice
	// arrstr:=deck()
	// bslice:=[]byte(str)
	// fmt.Println(bslice)
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cards.newDeckFromFile("my_cards")

	

}
