package main

import (
	"fmt"
	"net/http"

)
func readHeader(res http.ResponseWriter,req *http.Request){
   length := req.ContentLength
   body := make([]byte,length)
   req.Body.Read(body)   
   fmt.Fprintln(res,string(body))
}
func main(){
http.HandleFunc("/header",readHeader)
http.ListenAndServe(":8080",nil)
}