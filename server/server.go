package server

import (
	"fmt"
	"os"
	"time"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/log"

	"github.com/romberli/das/config"
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

	for i := 0; i < 30; i++ {
		log.Infof("%d time", i)
		time.Sleep(1 * time.Second)
	}

	s.Stop()
}

func (s *Server) Stop() {
	err := linux.RemovePidFile(s.PidFile)
	if err != nil {
		log.Errorf(config.ErrRemovePidFile, err.Error())
	}

	os.Exit(constant.DefaultNormalExitCode)
}
