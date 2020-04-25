package config

const DEFAULT_LOG_LEVEL = "ERROR"
const DEFAULT_LOG_PATH = "logs/"
const DEFAULT_LOG_NAME = "logpipe.log"

type LogConf struct {
	Level string
	Path  string
}
