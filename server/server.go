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
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/log"
	"go.uber.org/zap/zapcore"

	"github.com/romberli/das/pkg/message"
	"github.com/romberli/das/router"
)

type Server interface {
	// Addr returns listen address
	Addr() string
	// PidFile returns pid file path
	PidFile() string
	// Router returns router
	Router() router.Router
	// Register registers router path
	Register()
	// Run runs server
	Run()
	// Stop stops server
	Stop()
}

var _ Server = (*server)(nil)

type server struct {
	*http.Server
	addr    string
	pidFile string
	router  router.Router
}

// NewServer returns new *server
func NewServer(addr string, pidFile string, readTimeout, writeTimeout int, router router.Router) *server {
	return &server{
		Server: &http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
		},
		addr:    addr,
		pidFile: pidFile,
		router:  router,
	}
}

// NewServerWithDefaultRouter returns new *server with default gin router
func NewServerWithDefaultRouter(addr string, pidFile string, readTimeout, writeTimeout int) *server {
	if log.GetLevel() != zapcore.DebugLevel {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.NewGinRouter()

	return NewServer(addr, pidFile, readTimeout, writeTimeout, r)
}

// Addr returns listen address
func (s *server) Addr() string {
	return s.addr
}

// PidFile returns pid file path
func (s *server) PidFile() string {
	return s.pidFile
}

// Router returns router
func (s *server) Router() router.Router {
	return s.router
}

// Register registers router path
func (s *server) Register() {
	s.router.Register()
}

// Run runs server
func (s *server) Run() {
	fmt.Println(fmt.Sprintf("server started. addr: %s, pid file: %s", s.addr, s.pidFile))

	err := s.router.Run(s.addr)
	if err != nil {
		log.Errorf("server run failed.\n%s", err.Error())
	}
}

// Stop stops server
func (s *server) Stop() {
	err := linux.RemovePidFile(s.pidFile)
	if err != nil {
		log.Error(fmt.Sprintf("%s\n%s", message.Messages[message.ErrRemovePidFile].Error(), err.Error()))
	}

	os.Exit(constant.DefaultNormalExitCode)
}
