package diagnostic

import (
	"github.com/influxdata/kapacitor/services/victorops"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type VictorOpsHandler struct {
	logger *zap.Logger
}

func (h *VictorOpsHandler) Error(msg string, err error) {
	h.logger.Error(msg, zap.Error(err))
}

func (h *VictorOpsHandler) WithContext(ctx map[string]string) victorops.Diagnostic {
	fields := []zapcore.Field{}
	for k, v := range ctx {
		fields = append(fields, zap.String(k, v))
	}

	return &VictorOpsHandler{
		logger: h.logger.With(fields...),
	}
}
