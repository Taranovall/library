package utils

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"library-service/jwt"
	"time"
)

func GetJwtToken(username string) string {
	conn, _ := grpc.Dial("jwt-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer conn.Close()

	c := jwt.NewJwtServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, _ := c.GenerateToken(ctx, &jwt.JwtRequest{
		Username: username,
	})

	return response.Token
}
