package api

import (
	"net/http"

	"boscoin.io/sebak/lib/block"
	"boscoin.io/sebak/lib/network/httputils"
	"boscoin.io/sebak/lib/node/runner/api/resource"
)

func (api NetworkHandlerAPI) GetBlockHandler(w http.ResponseWriter, r *http.Request) {
	hashReq, err := newBlockHashRequest(r)
	if err != nil {
		httputils.WriteJSONError(w, err)
		return
	}

	var res resource.Resource
	{
		var b block.Block
		var err error
		if hashReq.isHash {
			b, err = block.GetBlock(api.storage, hashReq.hash)
		} else {
			b, err = block.GetBlockByHeight(api.storage, hashReq.height)
		}
		if err != nil {
			httputils.WriteJSONError(w, err)
			return
		}
		res = resource.NewBlock(&b)
	}
	httputils.MustWriteJSON(w, 200, res)
}
