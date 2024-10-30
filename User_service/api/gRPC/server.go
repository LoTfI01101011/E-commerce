package gRPC

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/LoTfI01101011/E-commerce/User_service/api/gRPC/proto"
	"github.com/LoTfI01101011/E-commerce/User_service/internal"
	"github.com/LoTfI01101011/E-commerce/User_service/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
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
func addToBlackList(ctx context.Context, token string, time time.Duration, rdb *redis.Client) error {
	_, err := rdb.Set(ctx, token, "blacklisted", time).Result()
	return err
}
func CheckExparation(token string) time.Duration {
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Fatal("Failed")
		}
		return []byte(os.Getenv("Secret")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	claims, _ := tok.Claims.(jwt.MapClaims)
	expFloat := claims["exp"].(float64)
	expTime := time.Unix(int64(expFloat), 0)
	exp := expTime.Sub(time.Now())
	return exp
}
func ValidateToken(tokenString string, ctx context.Context) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Fatal("Unmatched method")
		}

		return []byte(os.Getenv("Secret")), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
func (s *Server) LoginUser(ctx context.Context, data *pb.LoginRequest) (*pb.Token, error) {
	//geting the user from the db
	var user models.User
	err := internal.DB.Where("email = ?", data.Email).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}
	if ctx.Err() != nil {
		// Handle timeout or cancellation
		return nil, fmt.Errorf("context error: Request timeout")
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
	if ctx.Err() != nil {
		// Handle timeout or cancellation
		return nil, fmt.Errorf("context error: Request timeout")
	}
	//generating jwt token
	token, err := GenerateToken(id)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}
	//insert in the db
	return &pb.Token{Token: token}, nil
}
func (s *Server) LogoutUser(ctx context.Context, tkn *pb.Token) (*pb.LogoutResponse, error) {
	//get the token
	tokenString := strings.TrimPrefix(tkn.Token, "Bearer ")
	//get the remaining duration still to expire the token
	exp := CheckExparation(tokenString)

	//chash the token to blacklist
	err := addToBlackList(ctx, tokenString, exp, internal.Redis)
	if err != nil {
		return nil, fmt.Errorf("There was problem with redis connection")
	}

	return &pb.LogoutResponse{ResponseMessage: "you are logged out successfuly"}, nil
}
func (s *Server) CheckUserToken(ctx context.Context, tkn *pb.Token) (*pb.CheckUserTokenResponse, error) {
	//trim the prefix of the token
	tokenString := strings.TrimPrefix(tkn.Token, "Bearer ")
	//checking if the tokne exict in the blacklist
	isblacklisted, err := internal.Redis.Get(ctx, tokenString).Result()
	if err == nil && isblacklisted == "blacklisted" {
		return &pb.CheckUserTokenResponse{IsValid: false}, fmt.Errorf("the user is unauthorized")
	}
	claims, err := ValidateToken(tokenString, ctx)
	if err != nil {
		return &pb.CheckUserTokenResponse{IsValid: false}, err
	}
	if claims["sub"].(string) == "" {
		return &pb.CheckUserTokenResponse{IsValid: false}, fmt.Errorf("the user is unauthorized")
	}
	return &pb.CheckUserTokenResponse{IsValid: true}, nil
}
func (s *Server) GetUserInfo(ctx context.Context, tkn *pb.Token) (*pb.GetUserInfoResponse, error) {
	//trim The prefixe of the token
	token := strings.TrimPrefix(tkn.Token, "Bearer ")
	//get the claims of the token
	claims, _ := ValidateToken(token, ctx)
	//get the user from the db
	var user models.User
	internal.DB.Where("id = ?", claims["sub"].(string)).Find(&user)
	return &pb.GetUserInfoResponse{UserId: user.ID.String(), Username: user.Name, Email: user.Email}, nil
}
