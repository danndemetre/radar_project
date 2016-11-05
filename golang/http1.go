package main

//mport "fmt"
import  "net/http"
import "os"
import "io"
import "encoding/json"
import "bytes"
    

/*
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
*/

type User struct{
    Id      string
    Balance uint64
}

func main() {
    u := User{Id: "US123", Balance: 8}
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(u)
    res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", b)
    io.Copy(os.Stdout, res.Body)
}


