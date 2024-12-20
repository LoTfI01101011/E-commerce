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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	registerRes, err := c.RegisterUser(ctx, &pb.RegisterRequest{Username: name, Email: email, Password: password})
	if err != nil {
		return "", err
	}
	return registerRes.Token, nil
}
func (u *User) LoginUser(email, password string) (string, error) {
	//init user client
	c := pb.NewUserServiceClient(u.conn)
	//init context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	loginRes, err := c.LoginUser(ctx, &pb.LoginRequest{Email: email, Password: password})
	if err != nil {
		return "", err
	}
	return loginRes.Token, nil
}
func (u *User) LogoutUser(token string) (string, error) {
	//init user client
	c := pb.NewUserServiceClient(u.conn)
	//init context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	logoutRes, err := c.LogoutUser(ctx, &pb.Token{Token: token})
	if err != nil {
		return "", err
	}
	return logoutRes.ResponseMessage, nil
}
func (u *User) CheckUserToken(token string) (bool, error) {
	c := pb.NewUserServiceClient(u.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	checkingRes, err := c.CheckUserToken(ctx, &pb.Token{Token: token})
	if err != nil {
		return false, err
	}
	return checkingRes.IsValid, nil
}
func (u *User) GetUserInfo(token string) (map[string]string, error) {
	c := pb.NewUserServiceClient(u.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	info, err := c.GetUserInfo(ctx, &pb.Token{Token: token})
	if err != nil {
		return nil, err
	}
	res := map[string]string{
		"userID":   info.UserId,
		"username": info.Username,
		"email":    info.Email,
	}
	return res, nil
}
