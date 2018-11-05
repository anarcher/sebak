package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"boscoin.io/sebak/lib/block"
	"boscoin.io/sebak/lib/common"
	"boscoin.io/sebak/lib/network/httputils"
	"boscoin.io/sebak/lib/node/runner"
	"boscoin.io/sebak/lib/node/runner/api/resource"
)

type TeeReadCloser struct {
	io.ReadCloser
	teeReader io.Reader
}

func (tee TeeReadCloser) Read(p []byte) (n int, err error) {
	return tee.teeReader.Read(p)
}

func NewTeeReadCloser(origin io.ReadCloser, w io.Writer) io.ReadCloser {
	return &TeeReadCloser{
		ReadCloser: origin,
		teeReader:  io.TeeReader(origin, w),
	}
}

func (api NetworkHandlerAPI) PostTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.WriteJSONError(w, err)
		return
	}

	message := common.NetworkMessage{Type: common.TransactionMessage, Data: body}
	checker := &runner.MessageChecker{
		DefaultChecker:  common.DefaultChecker{Funcs: api.txChecks},
		Consensus:       api.consensus,
		TransactionPool: api.transactionPool,
		Storage:         api.storage,
		LocalNode:       api.localNode,
		NetworkID:       api.consensus.NetworkID,
		Message:         message,
		Log:             log,
		Conf:            api.conf,
	}

	if err = common.RunChecker(checker, common.DefaultDeferFunc); err != nil {
		if len(checker.Transaction.H.Hash) > 0 {
			block.SaveTransactionHistory(api.storage, checker.Transaction, block.TransactionHistoryStatusRejected)
		}
		httputils.WriteJSONError(err)
		return
	}

	tx := checker.Transaction
	json.Unmarshal(bufferRequest.Bytes(), &tx)
	if err := httputils.WriteJSON(w, 200, resource.NewTransactionPost(tx)); err != nil {
		httputils.WriteJSONError(w, err)
	}
}
