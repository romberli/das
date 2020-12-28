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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/romberli/das/config"
)

var (
	baseDir       string
	cfgFile       string
	daemon        bool
	daemonStr     string
	logFileName   string
	logLevel      string
	logFormat     string
	logMaxSize    int
	logMaxDays    int
	logMaxBackups int
	serverPort    int
	serverPidFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "das",
	Short: "das is short for database autonomy service",
	Long:  `das provides database autonomy service, which includes disk space usage prediction, sql tuning advice, database health check, and so on...`,
	Run: func(cmd *cobra.Command, args []string) {
		// if no subcommand is set, it will print help information.
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				fmt.Printf("got error when printing help information.\n%s", err.Error())
			}
			return
		}

		// init config
		err := initConfig()
		if err != nil {
			fmt.Println("init config failed.\n", err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(constant.DefaultAbnormalExitCode)
	}
}

func init() {
	// set usage template
	rootCmd.SetUsageTemplate(UsageTemplateWithoutDefault())

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// config
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", constant.DefaultRandomString, "config file path")
	// daemon
	rootCmd.PersistentFlags().StringVar(&daemonStr, "daemon", constant.DefaultRandomString, fmt.Sprintf("whether run in background as a daemon(default: %s)", constant.FalseString))
	// log
	rootCmd.PersistentFlags().StringVar(&logFileName, "log-file", constant.DefaultRandomString, fmt.Sprintf("specify the log file name(default: %s)", filepath.Join(config.DefaultLogDir, log.DefaultLogFileName)))
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", constant.DefaultRandomString, fmt.Sprintf("specify the log level(default: %s)", log.DefaultLogLevel))
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", constant.DefaultRandomString, fmt.Sprintf("specify the log format(default: %s)", log.DefaultLogFormat))
	rootCmd.PersistentFlags().IntVar(&logMaxSize, "log-max-size", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max size(default: %d)", log.DefaultLogMaxSize))
	rootCmd.PersistentFlags().IntVar(&logMaxDays, "log-max-days", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max days(default: %d)", log.DefaultLogMaxDays))
	rootCmd.PersistentFlags().IntVar(&logMaxBackups, "log-max-backups", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max size(default: %d)", log.DefaultLogMaxBackups))
	// server
	rootCmd.PersistentFlags().IntVar(&serverPort, "server-port", constant.DefaultRandomInt, fmt.Sprintf("specify the server port(default: %d)", config.DefaultServerPort))
	rootCmd.PersistentFlags().StringVar(&serverPidFile, "server-pid-file", constant.DefaultRandomString, fmt.Sprintf("specify the server pid file path(default: %s)", filepath.Join(config.DefaultBaseDir, fmt.Sprintf("%s.pid", config.DefaultCommandName))))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() error {
	var err error

	// init default config
	err = initDefaultConfig()
	if err != nil {
		return errors.New(fmt.Sprintf("init default configuration failed.\n%s", err.Error()))
	}

	// read config with config file
	err = ReadConfigFile()
	if err != nil {
		return errors.New(fmt.Sprintf("read config file failed.\n%s", err.Error()))
	}

	// override config with command line arguments
	err = OverrideConfig()
	if err != nil {
		return errors.New(fmt.Sprintf("override command line arguments failed.\n%s", err.Error()))
	}

	// init log
	fileName := viper.GetString("log.fileName")
	level := viper.GetString("log.level")
	format := viper.GetString("log.format")
	maxSize := viper.GetInt("log.maxSize")
	maxDays := viper.GetInt("log.maxDays")
	maxBackups := viper.GetInt("log.maxBackups")

	fileNameAbs := fileName
	isAbs := filepath.IsAbs(fileName)
	if !isAbs {
		fileNameAbs, err = filepath.Abs(fileName)
		if err != nil {
			return errors.New(fmt.Sprintf("get absolute path of log file failed. log file: %s\n%s", fileName, err.Error()))
		}
	}
	_, _, err = log.InitLogger(fileNameAbs, level, format, maxSize, maxDays, maxBackups)
	if err != nil {
		return errors.New(fmt.Sprintf("initialize logger failed.\n%s", err.Error()))
	}

	return nil
}

// initDefaultConfig initiate default configuration
func initDefaultConfig() (err error) {
	// get base dir
	baseDir, err = filepath.Abs(config.DefaultBaseDir)
	if err != nil {
		return errors.New(fmt.Sprintf("get base dir of %s failed.\n%s", config.DefaultCommandName, err.Error()))
	}
	// set default config value
	config.SetDefaultConfig(baseDir)
	err = config.ValidateConfig()
	if err != nil {
		return err
	}

	return nil
}

// ReadConfigFile read configuration from config file, it will override the init configuration
func ReadConfigFile() (err error) {
	if cfgFile != constant.EmptyString && cfgFile != constant.DefaultRandomString {
		viper.SetConfigFile(cfgFile)
		err = config.ValidateConfig()
		if err != nil {
			return errors.New(fmt.Sprintf("validate config file configuration failed.\n%s", err.Error()))
		}
	}

	return nil
}

// OverrideConfig read configuration from command line interface, it will override the config file configuration
func OverrideConfig() (err error) {
	// override config
	if cfgFile != constant.EmptyString && cfgFile != constant.DefaultRandomString {
		viper.Set("config", cfgFile)
	}

	// override daemon
	if daemonStr != constant.DefaultRandomString {
		daemon, err = cast.ToBoolE(daemonStr)
		if err != nil {
			return errors.New(fmt.Sprintf("override daemon failed.\n%s", fmt.Sprintf(config.ErrNotValidDaemon, daemonStr)))
		}
		viper.Set("daemon", daemon)
	}

	// override log
	if logFileName != constant.DefaultRandomString {
		viper.Set("log.fileName", logFileName)
	}
	if logLevel != constant.DefaultRandomString {
		logLevel = strings.ToLower(logLevel)
		viper.Set("log.level", logLevel)
	}
	if logFormat != constant.DefaultRandomString {
		logLevel = strings.ToLower(logFormat)
		viper.Set("log.format", logFormat)
	}
	if logMaxSize != constant.DefaultRandomInt {
		viper.Set("log.maxSize", logMaxSize)
	}
	if logMaxDays != constant.DefaultRandomInt {
		viper.Set("log.maxDays", logMaxDays)
	}
	if logMaxBackups != constant.DefaultRandomInt {
		viper.Set("log.maxBackups", logMaxBackups)
	}

	// override server
	if serverPort != constant.DefaultRandomInt {
		viper.Set("server.port", serverPort)
	}
	if serverPidFile != constant.DefaultRandomString {
		viper.Set("server.pidFile", serverPidFile)
	}

	// validate configuration
	err = config.ValidateConfig()
	if err != nil {
		return errors.New(fmt.Sprintf("validate command line arguments failed.\n%s", err.Error()))
	}

	return err
}

// UsageTemplateWithoutDefault returns a usage template which does not contain default part
func UsageTemplateWithoutDefault() string {
	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsagesWithoutDefault | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsagesWithoutDefault | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}
