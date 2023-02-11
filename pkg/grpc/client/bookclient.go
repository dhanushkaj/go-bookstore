package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dhanushkaj/go-bookstore/pkg/grpc/pb"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("error connecting to server")
	}
	defer conn.Close()

	c := pb.NewBookManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	book1 := &pb.Book{
		Name:        "Go in Action",
		Author:      "Erik St",
		Publication: "Manning",
	}

	book2 := &pb.Book{

		Name:        "Mastegin Go",
		Author:      "Mihalis",
		Publication: "Amazon",
	}

	books := []*pb.Book{}
	books = append(books, book1)
	books = append(books, book2)

	for _, b := range books {

		r, err := c.CreateBook(ctx, &pb.NewBook{Name: b.Name, Author: b.Author, Publication: b.Publication})

		if err != nil {
			log.Fatal("could not create usr", err)
		} else {

			log.Printf(`book created succesfully : 
		
		NAME:%s
		AGE:%s
		ID:%s`, r.GetName(), r.GetAuthor(), r.GetPublication())
		}
	}

}
