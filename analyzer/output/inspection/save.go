package inspection

import (
	"strings"
	"time"

	"github.com/pingcap/tidb-foresight/analyzer/boot"
	"github.com/pingcap/tidb-foresight/analyzer/input/args"
	"github.com/pingcap/tidb-foresight/analyzer/input/envs"
	"github.com/pingcap/tidb-foresight/analyzer/input/meta"
	"github.com/pingcap/tidb-foresight/analyzer/input/topology"
	"github.com/pingcap/tidb-foresight/model"
	"github.com/pingcap/tidb-foresight/utils"
	log "github.com/sirupsen/logrus"
)

type saveInspectionTask struct{}

func SaveInspection() *saveInspectionTask {
	return &saveInspectionTask{}
}

// Save inspection main record to database (then the frontend can see it)
func (t *saveInspectionTask) Run(m *boot.Model, c *boot.Config, args *args.Args, topo *topology.Topology, meta *meta.Meta, e *envs.Env) {
	components := map[string][]string{}

	for _, h := range topo.Hosts {
		for _, c := range h.Components {
			components[c.Name] = append(components[c.Name], h.Ip+":"+c.Port)
		}
	}

	if err := m.SetInspection(&model.Inspection{
		Uuid:         c.InspectionId,
		InstanceId:   args.InstanceId,
		InstanceName: topo.ClusterName,
		User:         e.Get("FORESIGHT_USER"),
		Status:       "running",
		Type:         e.Get("INSPECTION_TYPE"),
		Tidb:         strings.Join(components["tidb"], ","),
		Tikv:         strings.Join(components["tikv"], ","),
		Pd:           strings.Join(components["pd"], ","),
		Grafana:      strings.Join(components["grafana"], ","),
		Prometheus:   strings.Join(components["prometheus"], ","),
		CreateTime:   utils.NullTime{time.Unix(int64(meta.CreateTime), 0), true},
		FinishTime:   utils.NullTime{time.Unix(int64(meta.EndTime), 0), true},
		ScrapeBegin:  utils.NullTime{args.ScrapeBegin, true},
		ScrapeEnd:    utils.NullTime{args.ScrapeEnd, true},
	}); err != nil {
		log.Error("insert inspection:", err)
		return
	}
}