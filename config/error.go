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

var (
	ErrNotValidDaemon        = "Daemon must be either true or false, %s is not valid."
	ErrEmptyLogFileName      = "Log file name could NOT be an empty string."
	ErrNotValidLogFileName   = "Log file name must be either unix or windows path format, %s is not valid."
	ErrNotValidLogLevel      = "Log level must be one of [debug, info, warn, error, fatal], %s is not valid."
	ErrNotValidLogFormat     = "Log level must be either text or json, %s is not valid."
	ErrNotValidLogMaxSize    = "Log max size must be between %d and %d, %d is not valid."
	ErrNotValidLogMaxDays    = "Log max days must be between %d and %d, %d is not valid."
	ErrNotValidLogMaxBackups = "Log max backups must be between %d and %d, %d is not valid."
	ErrNotValidServerPort    = "Server port must be between %d and %d, %d is not valid."
	ErrNotValidPidFile       = "pid file name must be either unix or windows path format, %s is not valid."
)
