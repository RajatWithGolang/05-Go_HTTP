package main

import(
	"fmt"
	"net/http"
)

func indexHandler(res http.ResponseWriter,req *http.Request){
   fmt.Fprintf(res,"Hello Index")
}
func dogHandler(res http.ResponseWriter,req *http.Request){
	fmt.Fprintf(res,"Dog Dog Dog")
}

func meHandler(res http.ResponseWriter,req *http.Request){
	fmt.Fprintf(res,"Rajat Singh Rawat")
}

func main(){
	http.HandleFunc("/",indexHandler)
	http.Handle("/dog",http.HandlerFunc(dogHandler))
	http.HandleFunc("/me",meHandler)
	http.ListenAndServe(":8080",nil)
}