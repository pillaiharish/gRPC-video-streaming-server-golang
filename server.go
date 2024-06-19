package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "video-streaming-server-golang/video-streaming-server-golang/videostreaming"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var videoDirectory string

type server struct {
	pb.UnimplementedVideoStreamingServiceServer
}

func (s *server) StreamVideo(req *pb.VideoRequest, stream pb.VideoStreamingService_StreamVideoServer) error {
	filePath := filepath.Join(videoDirectory, req.GetVideoName())
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("could not read file: %v", err)
		}
		if err := stream.Send(&pb.VideoChunk{Content: buf[:n]}); err != nil {
			return fmt.Errorf("could not send chunk: %v", err)
		}
	}
	return nil
}

func (s *server) ListVideos(ctx context.Context, empty *pb.Empty) (*pb.VideoList, error) {
	var videoNames []string
	err := filepath.Walk(videoDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			videoNames = append(videoNames, info.Name())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.VideoList{VideoNames: videoNames}, nil
}

func main() {
	videoDirectory = os.Getenv("VIDEO_DIRECTORY")
	if videoDirectory == "" {
		log.Fatal("VIDEO_DIRECTORY environment variable is not set")
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVideoStreamingServiceServer(s, &server{})
	reflection.Register(s)
	log.Println("gRPC server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
