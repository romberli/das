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
package cmd

import (
	"fmt"
	"os"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/romberli/das/config"
	"github.com/romberli/das/server"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start command",
	Long:  `start the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err       error
			isRunning bool
			command   string
			pid       int
			pidBytes  []byte
		)

		// init config
		err = initConfig()
		if err != nil {
			fmt.Println("init config failed.\n", err.Error())
		}

		// check pid file
		serverPidFile = viper.GetString("server.pidFile")
		isRunning, err = linux.IsRunningWithPidFile(serverPidFile)
		if err != nil {
			log.Errorf("check if server is running failed when starting the server.\n%s", err.Error())
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		if isRunning {
			log.Errorf("pid file exists or pid is still running, pid file: %s.", serverPidFile)
			os.Exit(constant.DefaultAbnormalExitCode)
		}

		// check if runs in daemon mode
		daemon = viper.GetBool("daemon")
		if daemon {
			args = os.Args[1:]
			for i, arg := range os.Args[1:] {
				if config.TrimSpaceOfArg(arg) == config.DaemonArgTrue {
					args[i] = config.DaemonArgFalse
				}
			}

			for _, arg := range args {
				command = fmt.Sprintf("%s %s", command, arg)
			}
			_, err = linux.ExecuteCommand(command)
			if err != nil {
				log.Errorf("got error when starting das as foreground.\n%s", err.Error())
				os.Exit(constant.DefaultAbnormalExitCode)
				// fmt.Printf("got error when starting agent as a daemon.\n%s", err.Error())
			}
			os.Exit(constant.DefaultNormalExitCode)
		} else {
			if pid == 0 {
				pid = os.Getpid()
			}

			err = linux.SavePid(pid, serverPidFile, constant.DefaultFileMode)
			if err != nil {
				log.Errorf("got error when writing pid to file.\n%s", err.Error())
				os.Exit(constant.DefaultAbnormalExitCode)
			}

			log.Infof("%s started successfully with pid %s, pid file: %s.",
				config.DefaultCommandName, string(pidBytes), serverPidFile)

			// handle signal

			// start server
			serverPort = viper.GetInt("server.port")
			serverPidFile = viper.GetString("server.pidFile")
			s := server.NewServer(serverPort, serverPidFile)
			go s.Run()
			linux.HandleSignalsWithPidFile(serverPidFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
