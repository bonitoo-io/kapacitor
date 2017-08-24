package diagnostic

import (
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

func (h *VictorOpsHandler) WithContext(ctx map[string]string) victorops.Diagnostic {
	fields := []zapcore.Field{}
	for k, v := range ctx {
		fields = append(fields, zap.String(k, v))
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

func (h *SlackHandler) WithContext(ctx map[string]string) slack.Diagnostic {
	fields := []zapcore.Field{}
	for k, v := range ctx {
		fields = append(fields, zap.String(k, v))
	}

	return &SlackHandler{
		l: h.l.With(fields...),
	}
}
