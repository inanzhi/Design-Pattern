package option_pattern

import "fmt"

// 选项模式 可以创建一个带有默认值的struct变量 并选择性的修改其中一些参数的值。
type Server struct {
	Address  string
	Port     int
	Protocol string
	Timeout  int
}
type Option func(*Server)

func NewServer(address string, port int, opts ...func(*Server)) *Server {
	server := &Server{
		Address:  address,
		Port:     port,
		Protocol: "http", // 默认值
		Timeout:  30,     // 默认值
	}
	for _, opt := range opts {
		opt(server)
		//func(s *Server) {
		//	s.Protocol = protocol
		//}
	}
	return server
}

// 等同于
//
//	func(s *Server) {
//		s.Protocol = "https"
//	}
//
// 但大型项目应该使用 type Option func(*Server) 清晰
func WithProtocol(protocol string) func(*Server) {
	return func(s *Server) {
		s.Protocol = protocol
	}
}
func WithTimeout(timeout int) func(*Server) {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func main() {
	//WithProtocol("https")
	//等同于
	//func(s *Server) {
	//	s.Protocol = "https"
	//}

	server := NewServer("localhost", 8080, WithProtocol("https"), WithTimeout(60))
	fmt.Printf("Address: %s, Port: %d, Protocol: %s, Timeout: %d\n",
		server.Address, server.Port, server.Protocol, server.Timeout)
}
