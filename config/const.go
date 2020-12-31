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
package config

import (
	"github.com/romberli/go-util/constant"
)

// global constant
const (
	DefaultCommandName    = "das"
	DefaultDaemon         = false
	DefaultMode           = "agent"
	DefaultConfigFileName = "das.yaml"
	DefaultBaseDir        = constant.CurrentDir
	DefaultConfigDir      = "./config"
	DefaultLogDir         = "./log"
	MinLogMaxSize         = 1
	MaxLogMaxSize         = constant.MaxInt
	MinLogMaxDays         = 1
	MaxLogMaxDays         = constant.MaxInt
	MinLogMaxBackups      = 1
	MaxLogMaxBackups      = constant.MaxInt
	DefaultServerPort     = 6090
	DaemonArgTrue         = "--daemon=true"
	DaemonArgFalse        = "--daemon=false"
)

// configuration constant
const (
	ConfKey          = "config"
	DaemonKey        = "daemon"
	LogFileNameKey   = "log.fileName"
	LogLevelKey      = "log.level"
	LogFormatKey     = "log.format"
	LogMaxSizeKey    = "log.maxSize"
	LogMaxDaysKey    = "log.maxDays"
	LogMaxBackupsKey = "log.maxBackups"
	ServerPortKey    = "server.port"
	ServerPidFileKey = "server.pidFile"
)
