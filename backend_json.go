package log

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	DefaultTimestampKey    = "time"
	DefaultTimestampLayout = time.RFC3339Nano
	DefaultDomainKey       = "domain"
	DefaultLevelKey        = "level"
	DefaultMessageKey      = "msg"
	DefaultDataKey         = "data"
)

type JSONBackendCfg struct {
	TimestampKey    string
	TimestampLayout string
	DomainKey       string
	LevelKey        string
	MessageKey      string
	DataKey         string
}

type JSONBackend struct {
	Cfg JSONBackendCfg
}

func NewJSONBackend(cfg JSONBackendCfg) *JSONBackend {
	if cfg.TimestampKey == "" {
		cfg.TimestampKey = DefaultTimestampKey
	}

	if cfg.TimestampLayout == "" {
		cfg.TimestampLayout = DefaultTimestampLayout
	}

	if cfg.DomainKey == "" {
		cfg.DomainKey = DefaultDomainKey
	}

	if cfg.LevelKey == "" {
		cfg.LevelKey = DefaultLevelKey
	}

	if cfg.MessageKey == "" {
		cfg.MessageKey = DefaultMessageKey
	}

	if cfg.DataKey == "" {
		cfg.DataKey = DefaultDataKey
	}

	return &JSONBackend{
		Cfg: cfg,
	}
}

func (b *JSONBackend) Log(msg Message) {
	obj := map[string]interface{}{
		b.Cfg.TimestampKey: msg.Time.Format(b.Cfg.TimestampLayout),
		b.Cfg.DomainKey:    msg.domain,
		b.Cfg.LevelKey:     msg.FullLevel(),
		b.Cfg.MessageKey:   msg.Message,
		b.Cfg.DataKey:      msg.Data,
	}

	data, err := json.Marshal(obj)
	if err != nil {
		// Not much we can do here
		fmt.Fprintf(os.Stderr, "cannot encode log message %#v: %w\n", obj, err)
		return
	}

	fmt.Fprintf(os.Stderr, "%s\n", data)
}
