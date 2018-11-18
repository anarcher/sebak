package api

import (
	"net/http"
	"strconv"

	"boscoin.io/sebak/lib/errors"
	"github.com/gorilla/mux"
)

type blockHashRequest struct {
	hash   string
	height uint64
	isHash bool
}

func newBlockHashRequest(r *http.Request) (*blockHashRequest, error) {
	vars := mux.Vars(r)
	hash := vars["hashOrHeight"]
	if hash == "" {
		return nil, errors.BadRequestParameter
	}

	req := &blockHashRequest{}

	height, err := strconv.ParseUint(hash, 10, 64)
	if err != nil {
		req.hash = hash
		req.isHash = true
	} else {
		req.height = height
	}

	return req, nil
}
