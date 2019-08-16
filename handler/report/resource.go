package report

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pingcap/fn"
	"github.com/pingcap/tidb-foresight/model/report"
	"github.com/pingcap/tidb-foresight/utils"
	log "github.com/sirupsen/logrus"
)

type getResourceInfoHandler struct {
	m ResourceInfoGeter
}

func ResourceInfo(m ResourceInfoGeter) http.Handler {
	return &getResourceInfoHandler{m}
}

func (h *getResourceInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn.Wrap(h.getInspectionResourceInfo).ServeHTTP(w, r)
}

func (h *getResourceInfoHandler) getInspectionResourceInfo(r *http.Request) ([]*report.ResourceInfo, utils.StatusError) {
	inspectionId := mux.Vars(r)["id"]
	info, err := h.m.GetInspectionResourceInfo(inspectionId)
	if err != nil {
		log.Error("get inspection slow log:", err)
		return nil, utils.NewForesightError(http.StatusInternalServerError, "DB_QUERY_ERROR", "error on query data")
	}

	return info, nil
}
