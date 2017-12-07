package main

import "github.com/clzhan/XRtspServer/stream_server"

func main() {

	ser := stream_server.NewStreamServer(":8554")
	ser.Run()

}