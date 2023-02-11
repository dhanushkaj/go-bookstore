package main

import (
	"context"
	"log"
	"net"

	"github.com/dhanushkaj/go-bookstore/pkg/grpc/pb"
	"github.com/dhanushkaj/go-bookstore/pkg/repository"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

const (
	port = ":50051"
)

var db *gorm.DB

type BookManagementServer struct {
	pb.UnimplementedBookManagementServer
}

func NewBookManagementServer() *BookManagementServer {

	return &BookManagementServer{}
}

func (server *BookManagementServer) Run() error {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("failed to listen:%", err)

	}
	s := grpc.NewServer()
	pb.RegisterBookManagementServer(s, server)
	log.Printf("server listneinig at %v", lis.Addr())
	return s.Serve(lis)
}

func (server *BookManagementServer) CreateBook(ctx context.Context, in *pb.NewBook) (*pb.Book, error) {
	log.Println("book recived ", in.Name)
	created_book := &pb.Book{Name: in.Name, Author: in.Author, Publication: in.Publication}
	db.Create(&created_book)
	return created_book, nil
}

func main() {
	r, _ := repository.InitDB()
	db = r.DB
	var bookService *BookManagementServer = NewBookManagementServer()
	if err := bookService.Run(); err != nil {
		log.Fatal("failed to server :%v", err)
	}
}
