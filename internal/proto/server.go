package proto

import (
	"github.com/ProArash/go-movie-reseller/internal/proto/moviepb"
	"github.com/ProArash/go-movie-reseller/internal/proto/userpb"
	"google.golang.org/grpc"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

type MovieServer struct {
	moviepb.UnimplementedMovieServiceServer
}

func RegisterGrpcServers(gServer *grpc.Server) {
	userpb.RegisterUserServiceServer(gServer, &UserServer{})
	moviepb.RegisterMovieServiceServer(gServer, &MovieServer{})
}
