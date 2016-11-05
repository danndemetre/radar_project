package main
 
import (
    "fmt"
    "net"
    "time"
   // "strconv"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
 
    defer Conn.Close()
    //i := 0
    for {
    
        // string conversion from int to ascii
       // msg := strconv.Itoa(i)
       // i++
       
       id:= "$i04"
       alias:= "$aBedRoom"
       temp:= "$t34"
       
       msg := id+alias+temp
       
       
        buf := []byte(msg)
        _,err := Conn.Write(buf)
	
	
	
        if err != nil {
            fmt.Println(msg, err)
        }
        time.Sleep(time.Second * 1)
    }
}