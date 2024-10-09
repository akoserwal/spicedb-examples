package main

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
)

func main() {
	spicedbClient, err := NewSpicedbClient("localhost:50051", "foobar")

	clientStream, err := spicedbClient.ImportBulkRelationships(context.Background())
	if err != nil {
		fmt.Errorf("failed to intialize import bulk relationship: %w", err)
	}

	batch := []*v1.Relationship{
		{
			Subject: &v1.SubjectReference{
				Object: &v1.ObjectReference{
					ObjectType: "user",
					ObjectId:   fmt.Sprintf("%s%d", "bob", rand.Int()),
				},
			},
			Relation: "member",
			Resource: &v1.ObjectReference{
				ObjectType: "group",
				ObjectId:   fmt.Sprintf("%s%d", "admin", rand.Int()),
			},
		},
		{
			Subject: &v1.SubjectReference{
				Object: &v1.ObjectReference{
					ObjectType: "user",
					ObjectId:   fmt.Sprintf("%s%d", "tim", rand.Int()),
				},
			},
			Relation: "member",
			Resource: &v1.ObjectReference{
				ObjectType: "group",
				ObjectId:   fmt.Sprintf("%s%d", "grp", rand.Int()),
			},
		},
	}

	if err = clientStream.Send((*v1.ImportBulkRelationshipsRequest)(&v1.BulkImportRelationshipsRequest{
		Relationships: batch,
	})); err != nil {
		if !errors.Is(err, io.EOF) {
			fmt.Println(err)
		}
		fmt.Println(err)
	}

	for {
		res, err := clientStream.CloseAndRecv()
		if err != nil {
			if err == io.EOF {
				log.Print("EOF")
			}
			fmt.Errorf("failed to get response: %w", err)
		}
		if res != nil {
			// Print the number of imported tuples received in the response
			fmt.Printf("Relations Numloaded in spicedb: %d\n", res.NumLoaded)
			break
		}
	}

}

func NewSpicedbClient(endpoint string, token string) (*authzed.Client, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.EmptyDialOption{})
	opts = append(opts, grpcutil.WithInsecureBearerToken(token))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	spclient, err := authzed.NewClient(
		endpoint,
		opts...,
	)
	if err != nil {
		fmt.Errorf("failed to intialize spicedb client %w", err)
	}
	return spclient, err
}
