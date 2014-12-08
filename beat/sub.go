package main
import (
  "fmt"
  "time"
  "net"
  "net/rpc"
  "./beat"
)

// 日志管理
type Logger struct{
  // 路径
  Path    string
}
// 获取全部日志
func (l *Logger) All(size int, reply *[]beat.Log) error {
  //TODO 读取path下的日志
  *reply = append(*reply, beat.Log{"你好",time.Now()})
  return nil
}
func main(){
  fmt.Println("run")
  newServer := rpc.NewServer()
  newServer.Register(new(Logger))
  l, err := net.Listen("tcp", "127.0.0.1:3001")
  if err != nil {
    fmt.Println(err.Error())
  }
  newServer.Accept(l)
}
