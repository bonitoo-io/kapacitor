package diagnostic

import (
	"github.com/influxdata/kapacitor/keyvalue"
	"github.com/influxdata/kapacitor/services/slack"
	"github.com/influxdata/kapacitor/services/victorops"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type VictorOpsHandler struct {
	l *zap.Logger
}

func (h *VictorOpsHandler) Error(msg string, err error) {
	h.l.Error(msg, zap.Error(err))
}

func (h *VictorOpsHandler) WithContext(ctx ...keyvalue.T) victorops.Diagnostic {
	fields := []zapcore.Field{}
	for _, kv := range ctx {
		fields = append(fields, zap.String(kv.Key, kv.Value))
	}

	return &VictorOpsHandler{
		l: h.l.With(fields...),
	}
}

type SlackHandler struct {
	l *zap.Logger
}

func (h *SlackHandler) InsecureSkipVerify() {
	h.l.Warn("service is configured to skip ssl verification")
}

func (h *SlackHandler) Error(msg string, err error) {
	h.l.Error(msg, zap.Error(err))
}

func (h *SlackHandler) WithContext(ctx ...keyvalue.T) slack.Diagnostic {
	fields := []zapcore.Field{}
	for _, kv := range ctx {
		fields = append(fields, zap.String(kv.Key, kv.Value))
	}

	return &SlackHandler{
		l: h.l.With(fields...),
	}
}
