package main
import (
  "fmt"
  "net"
  "net/rpc"
  "./beat"
)

func main(){
  address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3001")
  if err != nil {
    fmt.Println(err.Error())
    panic(err)
  }
  conn, _ := net.DialTCP("tcp", nil, address)
  defer conn.Close()
  client := rpc.NewClient(conn)
  defer client.Close()

  var reply []beat.Log
  err = client.Call("Logger.All", 100, &reply)

  if err != nil {
    fmt.Println(err.Error())
  } else {
    for _,r := range reply{
      fmt.Println(r.Msg)
    }
  }
}
