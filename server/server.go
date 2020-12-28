package server

import (
	"fmt"
	"time"

	"github.com/romberli/log"
)

type Server struct {
	Port    int
	PidFile string
}

func NewServer(port int, pidFile string) *Server {
	return &Server{
		port,
		pidFile,
	}
}

func (s *Server) Run() {
	fmt.Println(fmt.Sprintf("server started. port: %d, pid file: %s", s.Port, s.PidFile))
	for i := 0; i < 100; i++ {
		log.Infof("%d time", i)
		time.Sleep(1 * time.Second)
	}
}
