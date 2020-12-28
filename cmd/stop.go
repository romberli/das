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
)

var pid int

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop command",
	Long:  `stop the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err       error
			isRunning bool
		)

		// init config
		err = initConfig()
		if err != nil {
			fmt.Println("init config failed.\n", err.Error())
		}

		// get pid
		if pid != constant.DefaultRandomInt {
			_, _ = linux.ExecuteCommand(fmt.Sprintf("kill %d", pid))
			return
		}
		serverPidFile = viper.GetString("server.pidFile")
		isRunning, err = linux.IsRunningWithPidFile(serverPidFile)
		if err != nil {
			log.Errorf("check if server is running failed when stopping the server.\n%s", err.Error())
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		if !isRunning {
			log.Errorf("pid file does NOT exist or pid is NOT running, pid file: %s.", serverPidFile)
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		pid, err = linux.GetPidFromPidFile(serverPidFile)
		if err != nil {
			log.Errorf("get pid from pid file failed.pid file: %s\n%s", serverPidFile, err.Error())
			os.Exit(constant.DefaultAbnormalExitCode)
		}

		// kill server pid
		_, _ = linux.ExecuteCommand(fmt.Sprintf("kill %d", pid))
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")
	stopCmd.PersistentFlags().IntVar(&pid, "server-pid", constant.DefaultRandomInt, fmt.Sprintf("specify the server pid"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
