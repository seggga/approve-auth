package grpc

import (
	"context"
	"net"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/seggga/approve-auth/internal/storage"
	"github.com/seggga/approve-auth/internal/tokens"

	pb "github.com/seggga/approve-auth/pkg/proto"
)

var _ pb.AuthAPIServer = &Server{}

// Server ...
type Server struct {
	addr      string
	srv       *grpc.Server
	stor      storage.UserStorage
	logger    *zap.SugaredLogger
	JWTSecret string
	pb.UnimplementedAuthAPIServer
}

// NewServer ..
func NewServer(stor storage.UserStorage, JWTSecret string, slog *zap.SugaredLogger) *Server {
	return &Server{
		addr:      ":" + os.Getenv("AUTH_PORT_4000_TCP_PORT"),
		stor:      stor,
		srv:       grpc.NewServer(),
		logger:    slog,
		JWTSecret: JWTSecret,
	}
}

// Start ...
func (s *Server) Start() {
	if s.addr == ":" {
		s.logger.Fatal("variable AUTH_PORT_4000_TCP_PORT not set")
	}
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		s.logger.Fatalf("error creating lestener on %s: %v", s.addr, err)
	}

	pb.RegisterAuthAPIServer(s.srv, s)
	go s.srv.Serve(lis)
	s.logger.Infof("gRPC service started on %s", s.addr)

}

// Stop ...
func (s *Server) Stop() {
	s.srv.GracefulStop()
	s.logger.Info("gRPC server stopped")
}

// CheckToken ...
func (s *Server) CheckToken(ctx context.Context, r *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {
	// check token
	s.logger.Debug("check token request")
	ok, login := tokens.CheckToken(r.Token, s.JWTSecret)
	if !ok {
		s.logger.Debugw("given access token is not valid",
			"login", login,
			"token", r.Token,
		)

		return &pb.CheckTokenResponse{
				Login: "",
				Error: pb.TokenNotValid},
			nil
	}

	// check existence of user with given login
	s.logger.Debugw("token is valid",
		"login", login,
	)
	_, err := s.stor.ReadUser(login)
	if err != nil {

		s.logger.Debugw("user not found: "+err.Error(),
			"login", login,
		)

		return &pb.CheckTokenResponse{
				Login: login,
				Error: pb.UserNotFound},
			err
	}

	s.logger.Debugw("request authenticated",
		"login", login,
	)

	return &pb.CheckTokenResponse{
			Login: login,
			Error: pb.NoError},
		nil
}

// RefreshTokens ...
func (s *Server) RefreshTokens(ctx context.Context, r *pb.RefreshTokensRequest) (*pb.RefreshTokensResponse, error) {
	s.logger.Debug("refresh token request")
	// check token
	ok, login := tokens.CheckToken(r.Token, s.JWTSecret)
	if !ok {

		s.logger.Debugw("given refresh token is not valid",
			"login", login,
		)

		return &pb.RefreshTokensResponse{
				AccessToken:  "",
				RefreshToken: "",
				Error:        pb.TokenNotValid},
			nil
	}
	// create tokens
	access, refresh, err := tokens.CreateTokenPair(login, os.Getenv("JWT_SECRET"))
	if err != nil {

		s.logger.Errorw("error creating token pair: "+err.Error(),
			"login", login,
		)

		return &pb.RefreshTokensResponse{
				AccessToken:  "",
				RefreshToken: "",
				Error:        pb.InternalServerError},
			err
	}

	s.logger.Debugw("token pair generated successfully",
		"login", login,
	)

	return &pb.RefreshTokensResponse{
			AccessToken:  access,
			RefreshToken: refresh,
			Error:        pb.NoError},
		nil
}
