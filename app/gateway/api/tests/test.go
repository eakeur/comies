package tests

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"testing"
	"time"
)

func NewTestServer(t *testing.T, server *grpc.Server) *grpc.ClientConn {
	lis, p, err := listener()
	if err != nil {
		t.Fatalf("Failed listening in an available port: %v", err)
	}

	go func() {
		log.Printf("Test server listening on port %v", p)
		if err := server.Serve(lis); err != nil {
			t.Fatalf("Test server stopped listening on port %v: %v", p, err)
		}
	}()

	time.Sleep(time.Second * 1)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect with server: %v", err)
	}

	return conn

}

func listener() (net.Listener, string, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, "", err
	}

	return l, fmt.Sprint(l.Addr().(*net.TCPAddr).Port), nil

}
