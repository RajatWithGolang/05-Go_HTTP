package main

import (
	"fmt"
	"net/http"

)
func readHeader(res http.ResponseWriter,req *http.Request){
   h := req.Header
   // or h:= req.Header["Accept-Encoding"]
   // or h := req.Header.Get("Accept-Encoding")
   fmt.Fprintln(res,h)
}
func main(){
http.HandleFunc("/header",readHeader)
http.ListenAndServe(":8080",nil)
}