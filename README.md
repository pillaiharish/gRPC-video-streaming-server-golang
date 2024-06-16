# video-streaming-server-golang
Play the video from a path to browser


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
You can set an environment variable for the current PowerShell session by using the $env command:
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
