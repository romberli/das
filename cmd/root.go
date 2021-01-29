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
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/romberli/das/config"
	"github.com/romberli/das/pkg/message"
)

var (
	baseDir                  string
	cfgFile                  string
	daemon                   bool
	daemonStr                string
	logFileName              string
	logLevel                 string
	logFormat                string
	logMaxSize               int
	logMaxDays               int
	logMaxBackups            int
	serverAddr               string
	serverPid                int
	serverPidFile            string
	serverReadTimeout        int
	serverWriteTimeout       int
	dbMySQLAddr              string
	dbMySQLName              string
	dbMySQLUser              string
	dbMySQLPass              string
	dbPoolMaxConnections     int
	dbPoolInitConnections    int
	dbPoolMaxIdleConnections int
	dbPoolMaxIdleTime        int
	dbPoolKeepAliveInterval  int
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
				fmt.Println(fmt.Sprintf("%s\n%s", message.NewMessage(message.ErrPrintHelpInfo).Error(), err.Error()))
				os.Exit(constant.DefaultAbnormalExitCode)
			}

			os.Exit(constant.DefaultNormalExitCode)
		}

		// init config
		err := initConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf("%s\n%s", message.NewMessage(message.ErrInitConfig).Error(), err.Error()))
			os.Exit(constant.DefaultAbnormalExitCode)
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
	rootCmd.PersistentFlags().IntVar(&logMaxSize, "log-max-size", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max size(default: %d, unit: MB)", log.DefaultLogMaxSize))
	rootCmd.PersistentFlags().IntVar(&logMaxDays, "log-max-days", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max days(default: %d)", log.DefaultLogMaxDays))
	rootCmd.PersistentFlags().IntVar(&logMaxBackups, "log-max-backups", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max size(default: %d)", log.DefaultLogMaxBackups))
	// server
	rootCmd.PersistentFlags().StringVar(&serverAddr, "server-port", constant.DefaultRandomString, fmt.Sprintf("specify the server port(default: %s)", config.DefaultServerAddr))
	rootCmd.PersistentFlags().StringVar(&serverPidFile, "server-pid-file", constant.DefaultRandomString, fmt.Sprintf("specify the server pid file path(default: %s)", filepath.Join(config.DefaultBaseDir, fmt.Sprintf("%s.pid", config.DefaultCommandName))))
	rootCmd.PersistentFlags().IntVar(&serverReadTimeout, "server-read-timeout", constant.DefaultRandomInt, fmt.Sprintf("specify the read timeout in seconds of http request(default: %d)", config.DefaultServerReadTimeout))
	rootCmd.PersistentFlags().IntVar(&serverWriteTimeout, "server-write-timeout", constant.DefaultRandomInt, fmt.Sprintf("specify the read timeout in seconds of http request(default: %d)", config.DefaultServerReadTimeout))
	// database
	rootCmd.PersistentFlags().StringVar(&dbMySQLAddr, "db-mysql-addr", constant.DefaultRandomString, fmt.Sprintf("specify database address(format: host:port)(default: %s)", fmt.Sprintf("%s:%d", constant.DefaultLocalHostIP, constant.DefaultMySQLPort)))
	rootCmd.PersistentFlags().StringVar(&dbMySQLName, "db-mysql-name", constant.DefaultRandomString, fmt.Sprintf("specify database name(default: %s)", config.DefaultDBName))
	rootCmd.PersistentFlags().StringVar(&dbMySQLUser, "db-mysql-user", constant.DefaultRandomString, fmt.Sprintf("specify database user name(default: %s)", config.DefaultDBUser))
	rootCmd.PersistentFlags().StringVar(&dbMySQLPass, "db-mysql-pass", constant.DefaultRandomString, fmt.Sprintf("specify database user password(default: %s)", config.DefaultDBPass))
	rootCmd.PersistentFlags().IntVar(&dbPoolMaxConnections, "db-pool-max-connections", constant.DefaultRandomInt, fmt.Sprintf("specify max connections of the connection pool(default: %d)", mysql.DefaultMaxConnections))
	rootCmd.PersistentFlags().IntVar(&dbPoolInitConnections, "db-pool-init-connections", constant.DefaultRandomInt, fmt.Sprintf("specify initial connections of the connection pool(default: %d)", mysql.DefaultMaxIdleConnections))
	rootCmd.PersistentFlags().IntVar(&dbPoolMaxIdleConnections, "db-pool-max-idle-connections", constant.DefaultRandomInt, fmt.Sprintf("specify max idle connections of the connection pool(default: %d)", mysql.DefaultMaxIdleConnections))
	rootCmd.PersistentFlags().IntVar(&dbPoolMaxIdleTime, "db-pool-max-idle-time", constant.DefaultRandomInt, fmt.Sprintf("specify max idle time of connections of the connection pool, (default: %d, unit: seconds)", mysql.DefaultMaxIdleTime))
	rootCmd.PersistentFlags().IntVar(&dbPoolKeepAliveInterval, "db-pool-keep-alive-interval", constant.DefaultRandomInt, fmt.Sprintf("specify keep alive interval of connections of the connection pool(default: %d, unit: seconds)", mysql.DefaultKeepAliveInterval))

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
		return multierror.Append(message.NewMessage(message.ErrInitDefaultConfig), err)
	}

	// read config with config file
	err = ReadConfigFile()
	if err != nil {
		return multierror.Append(message.NewMessage(message.ErrReadConfigFile), err)
	}

	// override config with command line arguments
	err = OverrideConfig()
	if err != nil {
		return multierror.Append(message.NewMessage(message.ErrOverrideCommandLineArgs), err)
	}

	// init log
	fileName := viper.GetString(config.LogFileKey)
	level := viper.GetString(config.LogLevelKey)
	format := viper.GetString(config.LogFormatKey)
	maxSize := viper.GetInt(config.LogMaxSizeKey)
	maxDays := viper.GetInt(config.LogMaxDaysKey)
	maxBackups := viper.GetInt(config.LogMaxBackupsKey)

	fileNameAbs := fileName
	isAbs := filepath.IsAbs(fileName)
	if !isAbs {
		fileNameAbs, err = filepath.Abs(fileName)
		if err != nil {
			return multierror.Append(message.NewMessage(message.ErrAbsoluteLogFilePath), err)
		}
	}
	_, _, err = log.InitLogger(fileNameAbs, level, format, maxSize, maxDays, maxBackups)
	if err != nil {
		return multierror.Append(message.NewMessage(message.ErrInitLogger), err)
	}

	log.SetDisableDoubleQuotes(true)
	log.SetDisableEscape(true)

	return nil
}

// initDefaultConfig initiate default configuration
func initDefaultConfig() (err error) {
	// get base dir
	baseDir, err = filepath.Abs(config.DefaultBaseDir)
	if err != nil {
		return multierror.Append(message.NewMessage(message.ErrBaseDir, config.DefaultCommandName), err)
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
		viper.SetConfigType("yaml")
		err = viper.ReadInConfig()
		if err != nil {
			return err
		}
		err = config.ValidateConfig()
		if err != nil {
			return multierror.Append(message.NewMessage(message.ErrValidateConfig), err)
		}
	}

	return nil
}

// OverrideConfig read configuration from command line interface, it will override the config file configuration
func OverrideConfig() (err error) {
	// override config
	if cfgFile != constant.EmptyString && cfgFile != constant.DefaultRandomString {
		viper.Set(config.ConfKey, cfgFile)
	}

	// override daemon
	if daemonStr != constant.DefaultRandomString {
		viper.Set(config.DaemonKey, daemonStr)
	}

	// override log
	if logFileName != constant.DefaultRandomString {
		viper.Set(config.LogFileKey, logFileName)
	}
	if logLevel != constant.DefaultRandomString {
		logLevel = strings.ToLower(logLevel)
		viper.Set(config.LogLevelKey, logLevel)
	}
	if logFormat != constant.DefaultRandomString {
		logLevel = strings.ToLower(logFormat)
		viper.Set(config.LogFormatKey, logFormat)
	}
	if logMaxSize != constant.DefaultRandomInt {
		viper.Set(config.LogMaxSizeKey, logMaxSize)
	}
	if logMaxDays != constant.DefaultRandomInt {
		viper.Set(config.LogMaxDaysKey, logMaxDays)
	}
	if logMaxBackups != constant.DefaultRandomInt {
		viper.Set(config.LogMaxBackupsKey, logMaxBackups)
	}

	// override server
	if serverAddr != constant.DefaultRandomString {
		viper.Set(config.ServerAddrKey, serverAddr)
	}
	if serverPidFile != constant.DefaultRandomString {
		viper.Set(config.ServerPidFileKey, serverPidFile)
	}
	if serverReadTimeout != constant.DefaultRandomInt {
		viper.Set(config.ServerReadTimeoutKey, serverReadTimeout)
	}
	if serverWriteTimeout != constant.DefaultRandomInt {
		viper.Set(config.ServerWriteTimeoutKey, serverWriteTimeout)
	}

	// override database
	if dbMySQLAddr != constant.DefaultRandomString {
		viper.Set(config.DBMySQLAddrKey, dbMySQLAddr)
	}
	if dbMySQLName != constant.DefaultRandomString {
		viper.Set(config.DBMySQLNameKey, dbMySQLName)
	}
	if dbMySQLUser != constant.DefaultRandomString {
		viper.Set(config.DBMySQLUserKey, dbMySQLUser)
	}
	if dbMySQLPass != constant.DefaultRandomString {
		viper.Set(config.DBMySQLPassKey, dbMySQLPass)
	}
	if dbPoolMaxConnections != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxConnectionsKey, dbPoolMaxConnections)
	}
	if dbPoolInitConnections != constant.DefaultRandomInt {
		viper.Set(config.DBPoolInitConnectionsKey, dbPoolInitConnections)
	}
	if dbPoolMaxIdleConnections != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxIdleConnectionsKey, dbPoolMaxIdleConnections)
	}
	if dbPoolMaxIdleTime != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxIdleTimeKey, dbPoolMaxIdleTime)
	}
	if dbPoolKeepAliveInterval != constant.DefaultRandomInt {
		viper.Set(config.DBPoolKeepAliveIntervalKey, dbPoolKeepAliveInterval)
	}

	// validate configuration
	err = config.ValidateConfig()
	if err != nil {
		return multierror.Append(message.NewMessage(message.ErrValidateConfig), err)
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
