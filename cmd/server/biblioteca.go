package main

import (
	"context"

	pb "algogrit.com/grpc-biblioteca/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BibliotecaServerImpl struct {
	books      map[int64]pb.Book
	idSequence int64
	pb.UnimplementedBibliotecaServer
}

func (bs *BibliotecaServerImpl) nextID() int64 {
	bs.idSequence++
	return bs.idSequence
}

func (bs *BibliotecaServerImpl) Show(ctx context.Context, br *pb.BookRequest) (*pb.Book, error) {
	book := bs.books[br.ID]

	return &book, nil
}

func (bs *BibliotecaServerImpl) Index(empty *emptypb.Empty, index pb.Biblioteca_IndexServer) error {
	for _, value := range bs.books {
		index.Send(&value)
	}
	return nil
}

func NewBibliotecaServer() pb.BibliotecaServer {
	server := &BibliotecaServerImpl{}

	server.books = map[int64]pb.Book{
		1: {ID: server.nextID(), Title: "The C Book", Author: "Dennis Ritchie"},
		2: {ID: server.nextID(), Title: "C++", Author: "Bjarne Stroustrop"},
	}

	return server
}
