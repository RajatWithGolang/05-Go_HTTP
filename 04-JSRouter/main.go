package main

import(
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func hello(res http.ResponseWriter,req *http.Request,p httprouter.Params){
   fmt.Fprintf(res,"Hello, %s!\n",p.ByName("name"))
}
  

func main(){
	mux := httprouter.New()
	mux.GET("/hello/:name",hello)
	http.ListenAndServe(":8080",mux)
}