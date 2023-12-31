package main

import (
	"context"
	"fmt"
	"github.com/Noah-Wilderom/go-logger-service/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

const (
	webPort      = 80
	mongoURL     = "mongodb://mongo:27017"
	mongoTimeout = 15 * time.Second
	rpcPort      = 5001
	grpcPort     = 50001
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect to mongodb
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// Register the RPC server
	err = rpc.Register(new(RPCServer))
	go app.rpcListen()

	go app.gRPCListen()

	log.Println("Starting logger service on port", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func (app *Config) rpcListen() {
	log.Println("Starting RPC server on port", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", rpcPort))
	if err != nil {
		log.Println("RPC error:", err)
	}
	defer listen.Close()

	for {
		rpcCon, err := listen.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(rpcCon)
	}
}

func connectToMongo() (*mongo.Client, error) {
	// create the connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})

	// connect
	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connection:", err)
		return nil, err
	}

	log.Println("Connected to mongodb")

	return conn, nil
}
