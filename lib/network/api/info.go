package api

import (
	"net/http"

	"boscoin.io/sebak/lib/network/httputils"
)

func (api NetworkHandlerAPI) GetInfoHandler(w http.ResponseWriter, r *http.Request) {

	type Info struct {
		Version     string `json:"version"`
		BlockHeight uint64 `json:"blockHeight"`
		BlockHash   string `json:"blockHash"`
		Connected   int    `json:"connected"`
	}

	info := &Info{}

	if err := httputils.WriteJSON(w, 200, info); err != nil {
		httputils.WriteJSONError(w, err)
	}
}
