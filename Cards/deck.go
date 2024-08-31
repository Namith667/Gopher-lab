package main
import "fmt"

type deck []string

func newDeck() deck{
	cards := deck{}
	cardSuites  :=[]string {"Spades","Diamonds","Hearts","Clubs"}
	cardValues  :=[]string {"Ace","Two","Three","Four"}
	for _, suites:= range cardSuites{
		for _,values := range cardValues{
			cards=append(cards,suites+" of "+values)
		}
	}
	return cards	
}


func (d deck) print(){
	for i,card := range d{
		fmt.Println(i,card)
	}
}
func deal(d deck, handSize int ) (deck,deck){
	return d[:handSize],d[handSize:]
}