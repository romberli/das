# copernicus configuration file description
If set an argument both in command line and configuration file, the command line argument will take effect.
Note: arguments are case-sensitive.

## config
- description: config file path.
- command-line-argument: --config
- type: string
- range: 
- default: 

## daemon
- description: if true, the program will run in background as a daemon, otherwise, it will run in the foreground.
- command-line-argument: --daemon
- type: bool
- range: true, false
- default: false

## log
- description: log section.
- command-line-argument: 
- type: 
- range: 
- default: 

## log.fileName
- description: log file name, use the absolute path.
- command-line-argument: --log-file
- type: string
- range: 
- default: "run.log"

## log.level
- description: specify log level
- command-line-argument: --log-level
- type: string
- range: debug, info, warn, error, fatal
- default: info

## log.format
- description: specify log format
- command-line-argument: --log-format
- type: string
- range: text, json
- default: text

## log.MaxSize
- description: Max log file size in MB, when reached the ***maxSize***, copernicus will rotate the log file.
- command-line-argument: --log-maxsize
- type: int
- range: 1 to 2^64 - 1, it's also affected by file system limits.
- default: 100

## log.MaxDays
- description: Max log keep days, log files older than ***maxDays*** will be deleted.
- command-line-argument: --log-max-days
- type: int
- range: 1 to 2^64 - 1
- default: 7

## log.MaxBackups
- description: max log backups, the oldest log files which are more than ***maxBackups*** will be deleted.
- command-line-argument: --log-max-backups
- type: int
- range: 1 to 2^64 - 1
- default: 5

## server
- description: server section of copernicus, if copernicus start with server subcommand, only server section will take effect. 
- command-line-argument: 
- type: 
- range: 
- default: 

## server.port
- description: port number to listen on
- command-line-argument: --server-port
- type: int
- range: 1 to 65535
- default: 6090

## server.pidFile
- description: pid file of the server process
- command-line-argument: --server-pid-file
- type: int
- range: 1 to 65535
- default: 6090