package api

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	"boscoin.io/sebak/lib/block"
	"github.com/stretchr/testify/require"
)

func TestBlockTransactionsHandler(t *testing.T) {
	ts, st := prepareAPIServer()
	defer st.Close()
	defer ts.Close()

	_, txs := prepareTxs(st, 100)
	b := block.GetLatestBlock(st)

	reqFunc := func(url string) ([]interface{}, map[string]interface{}) {
		respBody := request(ts, url, false)
		defer respBody.Close()
		reader := bufio.NewReader(respBody)

		bs, err := ioutil.ReadAll(reader)
		require.NoError(t, err)

		result := make(map[string]interface{})
		json.Unmarshal(bs, &result)
		records := result["_embedded"].(map[string]interface{})["records"].([]interface{})
		links := result["_links"].(map[string]interface{})
		return records, links
	}
	testFunc := func(hash, query string) ([]interface{}, map[string]interface{}) {
		url := strings.Replace(GetBlockTransactionsHandlerPattern, "{hashOrHeight}", hash, -1)
		url = url + "?" + query
		return reqFunc(url)
	}

	{
		q := "limit=100"
		records, _ := testFunc(b.Hash, q)
		require.Equal(t, len(records), 100)

		for i, a := range txs {
			b := records[i].(map[string]interface{})
			require.Equal(t, a.Hash, b["hash"])
		}
	}

}
