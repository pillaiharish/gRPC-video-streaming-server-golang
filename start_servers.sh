#!/bin/bash

# Set the video folder environment variable
export VIDEO_DIRECTORY="/Users/harishkumarpillai/stock_videos"

# Start the gRPC server
echo "Starting gRPC server..."
go run server.go &
grpc_pid=$!

# Wait for a few seconds to ensure the gRPC server starts
sleep 5

# Start the HTTP server
echo "Starting HTTP server..."
go run http_server.go &
http_pid=$!

# Wait for both servers to exit
wait $grpc_pid
wait $http_pid

