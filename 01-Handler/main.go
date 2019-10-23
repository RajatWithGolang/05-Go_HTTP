package main

import (
	 "fmt"
	 "net/http"
)
type customType string

func (c customType) ServeHTTP(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,"Hello Rajat")
}
func main(){
	var c customType 
	http.ListenAndServe(":8080",c)
}