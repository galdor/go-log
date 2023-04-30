package log

type BackendType string

const (
	BackendTypeTerminal BackendType = "terminal"
	BackendTypeJSON     BackendType = "json"
)

type Backend interface {
	Log(Message)
}
