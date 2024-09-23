package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main(){


	resp, err :=http.Get("http://google.com")
	if err!= nil{
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	fmt.Println(resp.Status,resp.Proto)
	io.Copy(os.Stdout,resp.Body)
	

}