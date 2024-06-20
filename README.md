# video-streaming-server-golang
Play the video from a path to browser. This will stream the video non stop, play pause is supported but since grpc is streaming video there is a challenge to handle seeking forward and backward explicitly. Typical HTTP video streaming protocols like HLS or DASH support this feature inherently by segmenting the video, which is not supported in this code.


### Creating protobuf
```bash
protoc --go_out=. --go-grpc_out=. video_stream.proto
```


### Install protoc and Go Plugin
Ensure you have protoc installed. You can download it from the Protocol Buffers releases page.


### Next, install the protoc-gen-go and protoc-gen-go-grpc plugins:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```


### Generate the gRPC Code
Run the protoc command again with the updated .proto file:
```bash
protoc --go_out=. --go-grpc_out=. -I=. video_stream.proto
```

###  video_streaming.proto is the latest proto that is being used in the code
```bash
protoc --go_out=. --go-grpc_out=. -I=. video_streaming.proto
```

## Running the Servers

### On macOS and Linux

1. **Temporary Environment Variable** 
You can set an environment variable for the current terminal session by using the export command:
```bash
export VIDEO_FOLDER="/path/to/your/video/folder"
```
To set the environment variable permanently, you can add the export command to your shell configuration file (~/.bash_profile, ~/.zshrc, or ~/.bashrc depending on the shell you are using).

2. **Make the script executable:**
```bash
chmod +x start_servers.sh
```
3. **Run the script:**
After setting the variable, you can run your script:
```bash
./start_servers.sh
```

### On Windows

1. **Open PowerShell.**

2. **Temporary Environment Variable**
You can set an environment variable for the current PowerShell session by using the $env command:
```powershell
$env:VIDEO_FOLDER = "D:\19\uploads"
```
To set the environment variable permanently, you can use the System Properties dialog or set it via PowerShell. 
In Advanced system settings click on the Environment Variables button.
In the System variables section, click New to add a new environment variable.
Enter VIDEO_FOLDER as the variable name and ```D:\19\uploads``` as the variable value.
Click OK to save the changes.

or Using PowerShell
```
[System.Environment]::SetEnvironmentVariable("VIDEO_FOLDER", "D:\19\uploads", "Machine")
```

3. **Run the script:**
After setting the variable, you can run your script:
```powershell
.\start_servers.ps1
```

## Running grpcurl to get the encoded data from video
```bash
harish $ grpcurl -plaintext -d '{"video_name": "data1.mp4"}' localhost:50051 videostreaming.VideoStreamingService/StreamVideo
{
  "content": "AAAAIGZ0eXBpc29tAAACAGlzb21pc28yYXZjMW1wNDEAA9ULbW9vdgAAAGxtdmhkAAAAAAAAAAAAAAAAAAAD6AAD4KgAAQAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwABhaB0cmFrAAAAXHRraGQAAAADAAAAAAAAAAAAAAABAAAAAAAD4KgAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAABAAAAAA1VVVQHgAAAAAAAkZWR0cwAAABxlbHN0AAAAAAAAAAEAA+CoAAAEAAABAAAAAYUYbWRpYQAAACBtZGhkAAAAAAAAAAAAAAAAAAAyAAAxogBVxAAAAAAALWhkbHIAAAAAAAAAAHZpZGUAAAAAAAAAAAAAAABWaWRlb0hhbmRsZXIAAAGEw21pbmYAAAAUdm1oZAAAAAEAAAAAAAAAAAAAACRkaW5mAAAAHGRyZWYAAAAAAAAAAQAAAAx1cmwgAAAAAQABhINzdGJsAAAAq3N0c2QAAAAAAAAAAQAAAJthdmMxAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAA1QB4ABIAAAASAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGP//AAAANWF2Y0MBZAge/+EAHmdkCB6s2UDYPef/8CgAJ/EAAAMAAQAAAwAyDxYtlgEABGjq58sAAAAQcGFzcAAAAoAAAAJ/AAAAGHN0dHMAAAAAAAAAAQAAGNEAAAIAAAABoHN0c3MAAAAAAAAAZAAAAAEAAABWAAAA0wAAAOoAAAEcAAABmQAAAc4AAAJGAAACgAAAArMAAAL4AAADdQAAA/IAAARvAAAEtgAABTMAAAVgAAAFjQAABdMAAAYeAAAGRQAABqMAAAbEAAAG0wAABv4AAAc4AAAHWQAAB5UAAAgSAAAIjwAACNsAAAkZAAAJRgAACYYAAAmoAAAJ8QAACiQAAAqYAAAK2QAACuYAAAr4AAALKAAAC6UAAAwiAAAMiQAADMYAAA0oAAANlgAADfIAAA5vAAAO7AAAD0gAAA/FAAAP2QAAED8AABCHAAAQzAAAETAAABFNAAARuQAAEhIAABJLAAAShAAAErgAABLlAAATCwAAEywAABOVAAAT2gAAFBMAABRCAAAUdAAAFKcAABTQAAAVDgAAFVgAABVrAAAVeAAAFZEAABWgAAAVsAAAFcAAABXSAAAV4QAAFjcAABZwAAAWgAAAFowAABaxAAAWwAAAFswAABbxAAAXAA=="
}
{...}
```
Breakdown of the command:
```-plaintext```: Indicates that the server does not use TLS.

```-d '{"video_name": "data1.mp4"}'```: The JSON payload sent to the StreamVideo RPC, where "video_name" should match the field name defined in the .proto file.

localhost:50051: The address of the gRPC server.

videostreaming.VideoStreamingService/StreamVideo: The full name of the RPC method in the VideoStreamingService.


## To verify the available methods and services, we can also run:
```bash
harish $ grpcurl -plaintext localhost:50051 list
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
videostreaming.VideoStreamingService
```

## And to see the methods within a specific service:
```bash
harish $ grpcurl -plaintext localhost:50051 list videostreaming.VideoStreamingService
videostreaming.VideoStreamingService.ListVideos
videostreaming.VideoStreamingService.StreamVideo
```

## Screenshot for Videoplayer mobile browser
![Screeshot for mobile](https://github.com/pillaiharish/video-streaming-server-golang/blob/main/screen-capture-mobile.jpeg)