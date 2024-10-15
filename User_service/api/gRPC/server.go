package gRPC

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/LoTfI01101011/E-commerce/User_service/api/gRPC/proto"
	"github.com/LoTfI01101011/E-commerce/User_service/internal"
	"github.com/LoTfI01101011/E-commerce/User_service/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func GenerateToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("Secret")))
}

func (s *Server) LoginUser(ctx context.Context, data *pb.LoginRequest) (*pb.Token, error) {
	//geting the user from the db
	var user models.User
	err := internal.DB.Where("email = ?", data.Email).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}
	//chenking the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}
	//generating the token
	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate token")
	}
	//return the token
	return &pb.Token{Token: token}, nil
}
func (s *Server) RegisterUser(ctx context.Context, register *pb.RegisterRequest) (*pb.Token, error) {
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(register.Password), 10)
	if err != nil {
		return nil, fmt.Errorf("failed hashing the password")
	}
	//inseting the data on the db
	id, _ := uuid.NewV7()
	user := models.User{ID: id, Name: register.Username, Password: string(hash), Email: register.Email}
	if internal.DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	internal.DB.Create(user)
	//generating jwt token
	token, err := GenerateToken(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate token")
	}
	//insert in the db
	return &pb.Token{Token: token}, nil
}
func (s *Server) LogoutUser(context.Context, *pb.Token) (*pb.LogoutResponse, error) {
	return nil, nil
}
func (s *Server) CheckUserToken(context.Context, *pb.Token) (*pb.CheckUserTokenResponse, error) {
	return nil, nil
}
func (s *Server) GetUserInfo(context.Context, *pb.Token) (*pb.GetUserInfoResponse, error) {
	return nil, nil
}
