package gRPC

import (
	"context"
	"log"
	"time"

	pb "github.com/LoTfI01101011/E-commerce/User_service/api/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	address string
	conn    *grpc.ClientConn
}

func (u *User) Start(address string) {
	//init new grpc client
	u.address = address
	conn, err := grpc.NewClient(u.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error")
	}
	u.conn = conn

}
func (u *User) Stop() {
	u.conn.Close()
}

func (u *User) RegisterUser(name, email, password string) (string, error) {
	//init user client
	c := pb.NewUserServiceClient(u.conn)
	//init context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	registerRes, err := c.RegisterUser(ctx, &pb.RegisterRequest{Username: name, Email: email, Password: password})
	if err != nil {
		return "", err
	}
	return registerRes.Token, nil
}
