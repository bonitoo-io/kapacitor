package diagnostic

import (
	"github.com/influxdata/kapacitor/services/victorops"
	"go.uber.org/zap"
)

type Service interface {
	NewVictorOpsHandler() victorops.Diagnostic
}

type service struct {
	logger *zap.Logger
}

func NewService() Service {
	// TODO: change
	l := zap.NewExample()
	return &service{
		logger: l,
	}
}

func (s *service) NewVictorOpsHandler() victorops.Diagnostic {
	return &VictorOpsHandler{
		logger: s.logger.With(zap.String("service", "victorops")),
	}
}
