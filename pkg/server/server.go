package cmd

import (
	"context"
	"flag"
	"fmt"

	as "github.com/aerospike/aerospike-client-go"

	"github.com/petrpan26/go-grpc-http-rest-microservice-tutorial/pkg/protocol/grpc"
	"github.com/petrpan26/go-grpc-http-rest-microservice-tutorial/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBIp string
	// DatastoreDBUser is username to connect to database
	DatastoreDBPort int
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBIp, "db-ip", "", "Database ip")
	flag.StringVar(&cfg.DatastoreDBPort, "db-port", "", "Database port")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}
		param)
	db, err := as.NewClient(cfg.DatastoreDBIp, cfg.DatastoreDBPort)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
