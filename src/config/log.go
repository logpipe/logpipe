package config

const DEFAULT_LOG_LEVEL = "ERROR"
const DEFAULT_LOG_PATH = "logs"

type LogConf struct {
	Level string
	Path  string
}
