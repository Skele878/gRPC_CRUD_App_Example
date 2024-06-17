package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/Skele878/gRPC_CRUD_App_Example/db"
	pb "github.com/Skele878/gRPC_CRUD_App_Example/proto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Creating gRPC server

type Server struct {
	pb.UnimplementedMovieServiceServer
}

var (
	DB    *gorm.DB
	movie db.Movie
)

// Declaring Methods CreateMovie for server
// Define a Movie struct to generate table in Postgres db

func (*Server) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	fmt.Println("Create Movie")
	movie := req.GetMovie()
	movie.Id = uuid.New().String()

	data := db.Movie{
		ID:    movie.GetId(),
		Title: movie.GetTitle(),
		Genre: movie.GetGenre(),
	}

	res := DB.Create(&data)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie creation unsuccessful")
	}

	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.GetId(),
			Title: movie.GetTitle(),
			Genre: movie.GetGenre(),
		},
	}, nil
}

func (*Server) GetMovie(ctx context.Context, req *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {
	fmt.Println("Read Movie", req.GetId())
	res := DB.Find(&movie, "id = ?", req.GetId())
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}
	return &pb.ReadMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil
}

func (*Server) GetMovies(ctx context.Context, req *pb.ReadMoviesRequest) (*pb.ReadMoviesResponse, error) {
	fmt.Println("Read Movies")
	movies := []*pb.Movie{}
	res := DB.Find(&movies)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}
	return &pb.ReadMoviesResponse{
		Movies: movies,
	}, nil
}

func (*Server) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {
	fmt.Println("Update Movie")
	reqMovie := req.GetMovie()

	res := DB.Model(&movie).Where("id=?", reqMovie.Id).Updates(db.Movie{Title: reqMovie.Title, Genre: reqMovie.Genre})

	if res.RowsAffected == 0 {
		return nil, errors.New("movies not found")
	}

	return &pb.UpdateMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil
}

func (*Server) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	fmt.Println("Delete Movie")
	res := DB.Where("id = ?", req.GetId()).Delete(&movie)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}

	return &pb.DeleteMovieResponse{
		Success: true,
	}, nil
}
