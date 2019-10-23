package main

import (
	 "fmt"
	 "net/http"
)
//Handler Function
func indexHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,"Hi")
}
func helloHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,"Hello Rajat")
}
func main(){
	http.HandleFunc("/",indexHandler)// 
	http.HandleFunc("/hello",helloHandler)	
	http.ListenAndServe(":8080",nil)
}