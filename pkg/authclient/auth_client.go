package authclient

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/seggga/approve-auth/pkg/proto"
	"google.golang.org/grpc"
)

// TokenPair represents data object auth-client works with
type TokenPair struct {
	Access    string
	Refresh   string
	Login     string
	Refreshed bool
}

// Client ...
type Client struct {
	Conn *grpc.ClientConn
	pb.AuthAPIClient
}

// NewClient creates an instance of AuthClient
func NewClient() (*Client, error) {
	// create connection
	path := "mail:" + os.Getenv("AUTH_PORT_4000_TCP_PORT")
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(cwt, path, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("error creating connection %s: %v", path, err)
	}
	// create client
	client := pb.NewAuthAPIClient(conn)
	return &Client{
		Conn:          conn,
		AuthAPIClient: client,
	}, nil
}

// Authenticate executes jwt check and renew.
func (c *Client) Authenticate(ctx context.Context, tokenPair *TokenPair) (*TokenPair, error) {

	// create request with access token
	tokenReq := &pb.CheckTokenRequest{
		Token: tokenPair.Access,
	}
	// check access token
	tokenResp, err := c.CheckToken(ctx, tokenReq)
	if err != nil {
		return nil, fmt.Errorf("cannot check access token: %v", err)
	}

	// got a bad error while checking access token
	if tokenResp.GetError() != pb.TokenNotValid && tokenResp.Error != pb.NoError {
		return nil, fmt.Errorf("check token returned error: %v", tokenResp.GetError())
	}
	// access token is valid
	if tokenResp.Error == pb.NoError {
		return &TokenPair{
			Login:     tokenResp.GetLogin(),
			Refreshed: false,
		}, nil
	}

	// access token is not valid, we should check refresh token
	tokenReq = &pb.CheckTokenRequest{
		Token: tokenPair.Refresh,
	}
	tokenResp, err = c.CheckToken(ctx, tokenReq)
	if err != nil {
		return nil, fmt.Errorf("cannot check refresh token: %v", err)
	}

	// got a bad error while checking refresh token
	if tokenResp.GetError() != pb.NoError {
		return nil, fmt.Errorf("Unauthorized: %v", tokenResp.GetError())
	}

	refreshReq := &pb.RefreshTokensRequest{
		Token: tokenPair.Refresh,
	}
	refreshResp, err := c.RefreshTokens(ctx, refreshReq)
	if err != nil {
		return nil, fmt.Errorf("cannot check refresh token: %v", err)
	}
	// got a bad error while checking refresh token
	if refreshResp.GetError() != pb.NoError {
		return nil, fmt.Errorf("Unauthorized: %v", refreshResp.GetError())
	}

	return &TokenPair{
		Access:    refreshResp.GetAccessToken(),
		Refresh:   refreshResp.GetRefreshToken(),
		Login:     tokenResp.GetLogin(),
		Refreshed: true,
	}, nil
}

// Stop closes grpc connection
func (c *Client) Stop() {
	_ = c.Conn.Close()
}
