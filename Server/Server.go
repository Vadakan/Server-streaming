package main

import (
    "fmt"
    "github.com/Server_Streaming/stub/pb"
    "google.golang.org/grpc"
    "log"
    "net"
    "sync"
)

const(
 port = ":50051"
)
var wg sync.WaitGroup
type server struct{

  pb.UnimplementedStreamServiceServer
}

func (s server) FetchResponse(in *pb.Request,stream pb.StreamService_FetchResponseServer)  error {
   fmt.Println("input recieved :",*in)

   for i:=0;i<5;i++{
      wg.Add(i)
      go func(a int32){
          fmt.Println("Server stream count :",a)
          resp := new(pb.Response)
          resp.Result = a
          err := stream.Send(resp)
          if err != nil{
             log.Fatalln(err)
          }
          wg.Done()
      }(int32(i))

   }
   wg.Wait()
   return nil
}
func main(){

  lis,err := net.Listen("tcp",port)

  if err != nil{
     log.Fatalln(err)
  }

   s :=  grpc.NewServer()

    pb.RegisterStreamServiceServer(s,&server{})

    if err := s.Serve(lis);err != nil{
      log.Fatalln(err)
    }

}
