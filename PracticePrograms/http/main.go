package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
func main(){


	resp, err :=http.Get("http://google.com")
	if err!= nil{
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	fmt.Println(resp.Status,resp.Proto)

	//to read html doc
	//Reader interface helps to handle different types of data 
	//instead of implementing various diff types in funcs to handle those data 
	
	
	//Reader converts to byte[]-- output data can be used by anyone
	
	// bs := make([]byte, 9999) // Provide the length of the byte slice-- we assign length to byte slice because Read function does not configure the length of response. So we give a assumed value of 9999
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	
	//alternate way
	io.Copy(os.Stdout,resp.Body)
	

}