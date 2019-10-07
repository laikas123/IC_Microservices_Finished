
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"



    //we have given our routeguide directory an alias pb
	pb "github.com/laikas123/IC_Microservices_Final/ProtoFiles"
)


var (
    tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
    certFile   = flag.String("cert_file", "", "The TLS cert file")
    keyFile    = flag.String("key_file", "", "The TLS key file")
    jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
    port       = flag.Int("port", 10000, "The server port")
)

type routeGuideServer struct {

	mu         sync.Mutex // protects routeNotes
	
}


func (s *routeGuideServer) QueryLocations(ctx context.Context, req *pb.LocationStatus) (*pb.Number, error) {
    return nil, nil
}


func (s *routeGuideServer) CalculateDistance(ctx context.Context, req *pb.TwoPoints) (*pb.Number, error) {
    return nil, nil
}


func (s *routeGuideServer) CalculateGasLoss(ctx context.Context, req *pb.Number) (*pb.Number, error) {
    return nil, nil
}


func (s *routeGuideServer) CalculateLocationProfit(ctx context.Context, req *pb.Number) (*pb.Number, error) {
    return nil, nil
}


func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	
    
	return s
}

func main() {
    //this is where the entire thing starts so parse the flags
	flag.Parse()
    //listen via tcp on the port specified
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
    //A ServerOption sets options such as credentials, codec and keepalive parameters, etc.
    //here we have a variable for a slice of all our ServerOptions
	var opts []grpc.ServerOption
    //note we are referencing the variable tls which is 
    //defined by our command line flags
	if *tls {
		if *certFile == "" {
            //Path just gets the filepath
			*certFile = testdata.Path("server1.pem")
		}
       
		if *keyFile == "" {
            //path just gets the filepath
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {

			log.Fatalf("Failed to generate credentials %v", err)
		}
     
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	grpcServer := grpc.NewServer(opts...)
 
	pb.RegisterICCalculatorServiceServer(grpcServer, newServer())
  
	grpcServer.Serve(lis)
}

