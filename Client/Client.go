package main

import (
	"fmt"
	"github.com/Server_Streaming/stub/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
	"context"
)

const(
 address = "localhost:50051"
)

func main(){
      fmt.Println("Client started executing")
 conn,err := grpc.Dial(address,grpc.WithInsecure(),grpc.WithBlock())
	  fmt.Println("Client connection established")
 defer conn.Close()
 if err != nil{
    log.Fatalln(err)
 }
    c := pb.NewStreamServiceClient(conn)

    input := new(pb.Request)
    input.Id = 1
    resp,res_err := c.FetchResponse(context.Background(),input)

    if res_err != nil{
       log.Fatalln(res_err)
    }
    quit := make(chan bool)
    go func(){
       for i:=0;i<5;i++{
          res,final_err := resp.Recv()
		   if final_err == io.EOF || i == 4{
			   quit <- true
			   return
		   }
          if final_err != nil{
             log.Fatalln(final_err)
          }
          fmt.Println(res.GetResult())
          time.Sleep(time.Second*1)
                 }
    }()
    <- quit
	fmt.Println("Client ended executing")
}
