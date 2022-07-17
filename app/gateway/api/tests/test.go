package tests

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

func ExpectError(t *testing.T, err error, code codes.Code, msg string) {
	if err == nil {
		t.Errorf("test failed because it was expected to receive error")
	}

	st, _ := status.FromError(err)
	if st.Code() != code && st.Message() != msg {
		t.Errorf("test failed because it was expected to receive %v error, but received: %v ", st.Code().String(), err)
	}
}

func listener() (net.Listener, string, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, "", err
	}

	return l, fmt.Sprint(l.Addr().(*net.TCPAddr).Port), nil

}
