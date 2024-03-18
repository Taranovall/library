package grpc

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jwt-service/auth"
	"jwt-service/jwt"
	"net"
)

const (
	gRpcPort = "50001"
)

type JwtServer struct {
	jwt.UnimplementedJwtServiceServer
}

func (j *JwtServer) GenerateToken(ctx context.Context, req *jwt.JwtRequest) (*jwt.JwtResponse, error) {
	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		return nil, err
	}

	result := &jwt.JwtResponse{Token: token}
	logrus.Infof("JWT: %s", token)
	return result, nil
}

func GRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		logrus.Fatalf("Failed to listen to gRPC: %s", err)
	}

	s := grpc.NewServer()

	jwt.RegisterJwtServiceServer(s, &JwtServer{})
	logrus.Infof("gRPC server started on port: %d", gRpcPort)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("Failed to listen to gRPC: %s", err)
	}
}
