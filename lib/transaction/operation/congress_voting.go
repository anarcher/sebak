package operation

import (
	"boscoin.io/sebak/lib/common"
	"boscoin.io/sebak/lib/errors"
)

type CongressVoting struct {
	Contract string `json:"contract"`
	Voting   struct {
		Start uint64 `json:"start"`
		End   uint64 `json:"end"`
	} `json:"voting"`
	FundingAddress string        `json:"funding_address"`
	Amount         common.Amount `json:"amount"`
}

func NewCongressVoting(contract string, start, end uint64, amount common.Amount, fundingAddress string) CongressVoting {

	return CongressVoting{
		Contract: contract,
		Voting: struct {
			Start uint64 `json:"start"`
			End   uint64 `json:"end"`
		}{
			Start: start,
			End:   end,
		},
		Amount:         amount,
		FundingAddress: fundingAddress,
	}
}

func (o CongressVoting) IsWellFormed(common.Config) (err error) {
	if len(o.Contract) == 0 {
		return errors.OperationBodyInsufficient
	}

	if o.Voting.End < o.Voting.Start {
		return errors.InvalidOperation
	}
	return
}

func (o CongressVoting) HasFee() bool {
	return true
}
