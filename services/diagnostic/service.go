package diagnostic

import (
	"github.com/influxdata/kapacitor/services/reporting"
	"github.com/influxdata/kapacitor/services/slack"
	"github.com/influxdata/kapacitor/services/storage"
	"github.com/influxdata/kapacitor/services/task_store"
	"github.com/influxdata/kapacitor/services/victorops"
	"go.uber.org/zap"
)

type Service interface {
	NewVictorOpsHandler() victorops.Diagnostic
	NewSlackHandler() slack.Diagnostic
	NewStorageHandler() storage.Diagnostic
	NewTaskStoreHandler() task_store.Diagnostic
	NewReportingHandler() reporting.Diagnostic
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
		l: s.logger.With(zap.String("service", "victorops")),
	}
}

func (s *service) NewSlackHandler() slack.Diagnostic {
	return &SlackHandler{
		l: s.logger.With(zap.String("service", "slack")),
	}
}

func (s *service) NewTaskStoreHandler() task_store.Diagnostic {
	return &TaskStoreHandler{
		l: s.logger.With(zap.String("service", "task_store")),
	}
}

func (s *service) NewReportingHandler() reporting.Diagnostic {
	return &ReportingHandler{
		l: s.logger.With(zap.String("service", "reporting")),
	}
}

func (s *service) NewStorageHandler() storage.Diagnostic {
	return &StorageHandler{
		l: s.logger.With(zap.String("service", "storage")),
	}
}
