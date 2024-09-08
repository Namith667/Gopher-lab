package main

import (
	"fmt"
)
func main(){
	colors:= make(map[string]string)
	colors["white"]="#ff011"

	colors["red"]="ff1122"
	colors["yello"]="444ff"
	
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string){
	for color,hex:=range c{
		fmt.Printf("Key: %v 	Value: %v",color,hex)
		fmt.Println()
	}
	}  