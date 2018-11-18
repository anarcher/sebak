package api

import (
	"encoding/json"
	"net/http"

	"boscoin.io/sebak/lib/block"
	"boscoin.io/sebak/lib/network/httputils"
	"boscoin.io/sebak/lib/node/runner/api/resource"
)

func (api NetworkHandlerAPI) GetBlockTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	hashReq, err := newBlockHashRequest(r)
	if err != nil {
		httputils.WriteJSONError(w, err)
		return
	}

	var b block.Block
	{
		if hashReq.isHash {
			b, err = block.GetBlock(api.storage, hashReq.hash)
		} else {
			b, err = block.GetBlockByHeight(api.storage, hashReq.height)
		}
	}
	if err != nil {
		httputils.WriteJSONError(w, err)
		return
	}

	p, err := NewPageQuery(r)
	if err != nil {
		httputils.WriteJSONError(w, err)
		return
	}

	var txs []resource.Resource
	var cursor []byte
	{
		option := p.WalkOption()
		prefix := block.GetBlockTransactionKeyPrefixBlock(b.Hash)

		api.storage.Walk(prefix, option, func(key, value []byte) (bool, error) {
			var hash string
			if err := json.Unmarshal(value, &hash); err != nil {
				return false, err
			}
			cursor = key
			tx, err := block.GetBlockTransaction(api.storage, hash)
			txs = append(txs, resource.NewTransaction(&tx))
			return true, err
		})
	}

	list := p.ResourceList(txs, cursor)
	httputils.MustWriteJSON(w, 200, list)
}
