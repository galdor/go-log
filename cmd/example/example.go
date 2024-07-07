package main

import (
	"time"

	"go.n16f.net/log"
	"go.n16f.net/program"
)

func main() {
	p := program.NewProgram("example", "example for the go-log library")

	p.AddOption("", "backend", "backend", "terminal",
		"the logging backend to use")

	p.SetMain(mainCmd)

	p.ParseCommandLine()
	p.Run()
}

func mainCmd(p *program.Program) {
	loggerCfg := log.LoggerCfg{
		DebugLevel: 1,
	}

	backendTypeString := p.OptionValue("backend")
	switch backendTypeString {
	case "terminal":
		loggerCfg.BackendType = log.BackendTypeTerminal
		loggerCfg.TerminalBackend = &log.TerminalBackendCfg{
			Color: true,
		}

	case "json":
		loggerCfg.BackendType = log.BackendTypeJSON
		loggerCfg.JSONBackend = &log.JSONBackendCfg{
			TimestampLayout: time.RFC3339Nano,
		}

	default:
		p.Fatal("unknown backend type %q", backendTypeString)
	}

	logger, err := log.NewLogger("example", loggerCfg)
	if err != nil {
		p.Fatal("cannot create logger: %v", err)
	}

	logger.Debug(1, "a level 1 debug message")
	logger.Debug(2, "a level 2 debug message")
	logger.Info("an info message")
	logger.InfoData(log.Data{"a": 42, "b": "hello"},
		"an info message with additional data")
	logger.Error("an error message")

	childLogger := logger.Child("child", log.Data{"x": "y"})
	childLogger.Info("an info message by the child")
}
