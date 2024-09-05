package main
import "testing"
func TestNewDeck(t *testing.T){
	d := newDeck()
	if len(d) != 16{
		t.Errorf("Expacted deck length of 16 but got %v",len(d))
	}//%v - value passed(16)

	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card od Ace of Spades, but got %v",d[0])
	}

	if d[len(d)-1] != "Four of Clubs"{
		t.Errorf("expexted final card Four of clubs , but got %v",d[len(d)-1])
	}
}

