package main

import(
    "io"
    "net/http"
    "html/template"
)

func hello(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "Hello Asshole")
}

func main(){
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8000", nil)
}

