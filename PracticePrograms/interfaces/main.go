package main
import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main(){
	eb:=englishBot{}
	sb:=spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
} 

func (englishBot)getGreeting() string{
	//VERY CUSTOM LOGIC 
	return "Hello there"
}
func (spanishBot)getGreeting()string{
	//VERY CUSTOM LOGIC 
	return "Hola"
}
