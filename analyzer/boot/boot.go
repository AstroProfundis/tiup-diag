package boot

import (
	"github.com/pingcap/tidb-foresight/bootstrap"
	"github.com/pingcap/tidb-foresight/model"
	"github.com/pingcap/tidb-foresight/wrapper/db"
)

const (
	SQLITE = "sqlite3"
)

type bootstrapTask struct {
	inspectionId string
	home         string
	db           db.DB
}

// Generate config and connect database
func Bootstrap(inspectionId, home string) *bootstrapTask {
	_, db := bootstrap.MustInit(home)

	return &bootstrapTask{inspectionId, home, db}
}

// `Run()` in `bootstrapTask` create config of current collected files' paths and db model.
func (t *bootstrapTask) Run() (*Config, *Model) {
	return NewConfig(t.inspectionId, t.home), NewModel(t.inspectionId, model.New(t.db))
}
