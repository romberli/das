/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
		log.Error(fmt.Sprintf("%s\n%s", config.Messages[config.ErrRemovePidFile].Error(), err.Error()))
	}

	os.Exit(constant.DefaultNormalExitCode)
}
