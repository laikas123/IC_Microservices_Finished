
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	//again we provide an alias for our directory	
	pb "github.com/laikas123/IC_Microservices_Final/ProtoFiles"
	

	//these are just test credentials found online
	"google.golang.org/grpc/testdata"
)


var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

//small thing to point out the Client name is uppercase beginning for 
//an interface and is lowercase for a struct 
func getLocationData(client pb.ICCalculatorServiceClient, status *pb.LocationStatus) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer cancel()
	
	
	statusReturned, err := client.QueryLocations(ctx, status)
	if err != nil {
		log.Fatalf("%v.QueryLocations(_) = _, %v: ", client, err)
	}

	log.Println(statusReturned)
}



func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	//this is where we create our inital RouteGuideClient 
	//which in turn calls gets passed as RouteGuideClient to all methods
	client := pb.NewICCalculatorServiceClient(conn)

	pointLo := pb.Point{
		X: 0,
		Y: 0,
	}
	pointHi := pb.Point{
		X: 10,
		Y: 10,
	}

	newRectangle := pb.Rectangle{
		Lo: &pointLo,
		Hi: &pointHi,
	}

	// Looking for a valid feature
	getLocationData(client, &pb.LocationStatus{Usersonline: 10, Locationtoserve: &newRectangle})

	
}
