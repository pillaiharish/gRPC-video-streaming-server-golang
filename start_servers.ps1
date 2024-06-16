# Set the video folder environment variable
$env:VIDEO_FOLDER = "D:\19\uploads"

# Start the gRPC server
Write-Output "Starting gRPC server..."
Start-Process -NoNewWindow -FilePath "go" -ArgumentList "run server.go" -PassThru
$grpc_proc = Get-Process | Where-Object { $_.Path -like "*server.exe" }

# Wait for a few seconds to ensure the gRPC server starts
Start-Sleep -Seconds 5

# Start the HTTP server
Write-Output "Starting HTTP server..."
Start-Process -NoNewWindow -FilePath "go" -ArgumentList "run http_server.go" -PassThru
$http_proc = Get-Process | Where-Object { $_.Path -like "*http_server.exe" }

# Wait for both servers to exit
$grpc_proc.WaitForExit()
$http_proc.WaitForExit()
