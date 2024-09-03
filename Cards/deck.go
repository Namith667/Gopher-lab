package main

import (
	"fmt"
	"os"
	"strings"
)


type deck []string

// arrstr :=[] string deck([])

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

//since deck contains many cards of string we need to 
//convert all that cards to one sigle string seperated by commas
//use strings.Join func
func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(fileName string) error {
	//converting to byteSlice
	return os.WriteFile(fileName, []byte(d.toString()), 0666)

}
func newDeckFromFile(fileName string) deck{
	bs, err  :=os.ReadFile(fileName) 
	if err != nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
	s := strings.Split(string(bs),",")
	return deck(s)
}

