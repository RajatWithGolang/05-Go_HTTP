package main

import (
	"fmt"
	"net/http"
	"io/ioutil"

)
func process(res http.ResponseWriter,req *http.Request){
//   req.ParseMultipartForm(1024)
//   fileHeader := req.MultipartForm.File["uploaded"][0]
//   file,_ := fileHeader.Open()
//   data,_ := ioutil.ReadAll(file)
//   fmt.Fprintln(res,string(data))
//============If there is only one file to upload then===========
   file,_,_:= req.FormFile("uploaded")
    data,_ := ioutil.ReadAll(file)
   fmt.Fprintln(res,string(data))
}
func main(){
http.HandleFunc("/process",process)
http.ListenAndServe(":8080",nil)
}