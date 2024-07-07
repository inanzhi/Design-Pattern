package structure_and_constructor

import "fmt"

type ServerConfig struct {
	Address  string
	Port     int
	Protocol string
	Timeout  int
}

func NewServerConfig(address string, port int) *ServerConfig {
	return &ServerConfig{
		Address:  address,
		Port:     port,
		Protocol: "http", // 默认值
		Timeout:  30,     // 默认值
	}
}

func main() {
	config := NewServerConfig("localhost", 8080)
	fmt.Printf("Address: %s, Port: %d, Protocol: %s, Timeout: %d\n",
		config.Address, config.Port, config.Protocol, config.Timeout)
}
