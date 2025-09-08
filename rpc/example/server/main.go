package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request
	return nil
}
func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	// chạy rpc server trên port 1234
	listener, err := net.Listen("tcp", ":1234")
	// nếu có lỗi thì in ra
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	// vòng lặp để phục vụ nhiều client
	for {
		// accept một connection đến
		conn, err := listener.Accept()
		// in ra lỗi nếu có
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// phục vụ client trên một goroutine khác
		// để giải phóng main thread tiếp tục vòng lặp
		go rpc.ServeConn(conn)
	}
}
