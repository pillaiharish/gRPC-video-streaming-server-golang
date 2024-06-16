package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "video-streaming-server-golang/videostream"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedVideoStreamServer
	videoFolder string
}

func (s *server) StreamVideo(req *pb.VideoRequest, stream pb.VideoStream_StreamVideoServer) error {
	filePath := filepath.Join(s.videoFolder, req.Filename)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v", filePath, err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		stream.Send(&pb.VideoResponse{Data: buffer[:n]})
	}

	return nil
}

func main() {
	videoFolder := os.Getenv("VIDEO_FOLDER")
	if videoFolder == "" {
		log.Fatal("VIDEO_FOLDER environment variable is not set")
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterVideoStreamServer(grpcServer, &server{videoFolder: videoFolder})
	log.Println("gRPC server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
