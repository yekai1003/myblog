package main

import (
	"fmt"
	"net/http"
	//"gopkg.in/mgo.v2/bson"
)

func main() {
	fmt.Println("blog begin  ...")
	MgSess = &MongoSessin{}
	MgSess.Connect("localhost:27017")
	http.HandleFunc("/ping", Pong)
	http.HandleFunc("/upload", UploadFile)
	http.HandleFunc("/lists", Lists)
	http.Handle("/", http.FileServer(http.Dir("../file/")))
	http.ListenAndServe(":8086", nil)
}
