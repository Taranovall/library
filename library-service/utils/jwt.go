package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"library-service/jwt"
	"strings"
	"time"
)

func GetJwtToken(username string) string {
	conn, _ := grpc.Dial("jwt-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer conn.Close()

	c := jwt.NewJwtServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, _ := c.GenerateToken(ctx, &jwt.JwtUsername{
		Username: username,
	})

	return response.Token
}

func VerifyTokenHeader(ctx *gin.Context) (string, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.Trim(strings.SplitAfter(tokenHeader, "Bearer")[1], " ")
	result, err := parseToken(accessToken)

	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}

	return result, nil
}

func parseToken(token string) (string, error) {
	conn, _ := grpc.Dial("jwt-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer conn.Close()

	c := jwt.NewJwtServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.ParseToken(ctx, &jwt.JwtString{
		Token: token,
	})

	if err != nil {
		logrus.Warnf("Cannot parse token: %s", err)
		return "", err
	}

	return response.Username, nil
}
