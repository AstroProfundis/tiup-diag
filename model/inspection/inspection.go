package inspection

import (
	"github.com/pingcap/tidb-foresight/utils"
)

type Inspection struct {
	Uuid         string         `json:"uuid" gorm:"PRIMARY_KEY"`
	InstanceId   string         `json:"instance_id"`
	InstanceName string         `json:"instance_name"`
	User         string         `json:"user"`
	Status       string         `json:"status"`
	Message      string         `json:"message"`
	Type         string         `json:"type"`
	CreateTime   utils.NullTime `json:"create_time,omitempty"`
	FinishTime   utils.NullTime `json:"finish_time,omitempty"`
	ScrapeBegin  utils.NullTime `json:"scrape_begin,omitempty"`
	ScrapeEnd    utils.NullTime `json:"scrape_end,omitempty"`
	Tidb         string         `json:"tidb"`
	Tikv         string         `json:"tikv"`
	Pd           string         `json:"pd"`
	Grafana      string         `json:"grafana"`
	Prometheus   string         `json:"prometheus"`
}

const DIAG_FILTER = "type in ('auto', 'manual')"

func (m *inspection) ListAllInspections(page, size int64) ([]*Inspection, int, error) {
	insps := []*Inspection{}
	count := 0

	if err := m.db.Where(DIAG_FILTER).Offset((page - 1) * size).Limit(size).Find(&insps).Error(); err != nil {
		return nil, 0, err
	}

	if err := m.db.Where(DIAG_FILTER).Count(&count).Error(); err != nil {
		return nil, 0, err
	}

	return insps, count, nil
}

func (m *inspection) ListInspections(instId string, page, size int64) ([]*Inspection, int, error) {
	insps := []*Inspection{}
	count := 0
	filter := &Inspection{InstanceId: instId}
	query := m.db.Model(&Inspection{}).Where(DIAG_FILTER).Where(filter)

	if err := query.Offset((page - 1) * size).Limit(size).Find(&insps).Error(); err != nil {
		return nil, 0, err
	}

	if err := query.Count(&count).Error(); err != nil {
		return nil, 0, err
	}

	return insps, count, nil
}

func (m *inspection) SetInspection(insp *Inspection) error {
	return m.db.Save(insp).Error()
}

func (m *inspection) GetInspection(inspId string) (*Inspection, error) {
	insp := Inspection{}

	if err := m.db.Where(&Inspection{Uuid: inspId}).Take(&insp).Error(); err != nil {
		return nil, err
	}

	return &insp, nil
}

func (m *inspection) DeleteInspection(inspId string) error {
	return m.db.Delete(&Inspection{Uuid: inspId}).Error()
}

func (m inspection) UpdateInspectionStatus(inspId, status string) error {
	return m.db.Model(&Inspection{}).Where(&Inspection{Uuid: inspId}).Update("status", status).Error()
}