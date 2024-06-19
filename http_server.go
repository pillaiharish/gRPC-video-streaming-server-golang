package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	pb "video-streaming-server-golang/video-streaming-server-golang/videostreaming"

	"google.golang.org/grpc"
)

const (
	grpcServerAddress = "localhost:50051"
	staticDir         = "./static"
)

var videoDirectory string

func main() {
	videoDirectory = os.Getenv("VIDEO_DIRECTORY")
	if videoDirectory == "" {
		log.Fatal("VIDEO_DIRECTORY environment variable is not set")
	}

	http.HandleFunc("/videos", listVideosHandler)
	http.HandleFunc("/stream", streamVideoHandler)
	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	log.Println("HTTP server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func listVideosHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := pb.NewVideoStreamingServiceClient(conn)
	resp, err := client.ListVideos(context.Background(), &pb.Empty{})
	if err != nil {
		http.Error(w, "Failed to list videos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	for _, videoName := range resp.VideoNames {
		fmt.Fprintf(w, "<a href=\"/stream?video=%s\">%s</a><br>", videoName, videoName)
	}
}

func streamVideoHandler(w http.ResponseWriter, r *http.Request) {
	videoName := r.URL.Query().Get("video")
	if videoName == "" {
		http.Error(w, "Missing video parameter", http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := pb.NewVideoStreamingServiceClient(conn)
	stream, err := client.StreamVideo(context.Background(), &pb.VideoRequest{VideoName: videoName})
	if err != nil {
		http.Error(w, "Failed to stream video", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Failed to receive chunk", http.StatusInternalServerError)
			return
		}
		w.Write(chunk.Content)
	}
}
