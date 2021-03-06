package consensus

import (
	"sync"

	"boscoin.io/sebak/lib/ballot"
	"boscoin.io/sebak/lib/errors"
	"boscoin.io/sebak/lib/voting"
)

type RunningRound struct {
	sync.RWMutex

	VotingBasis  voting.Basis
	Proposer     string                              // LocalNode's `Proposer`
	Transactions map[ /* Proposer */ string][]string /* Transaction.Hash */
	Voted        map[ /* Proposer */ string]*RoundVote
	Ballots      []ballot.Ballot
}

func NewRunningRound(proposer string, blt ballot.Ballot) (*RunningRound, error) {
	transactions := map[string][]string{
		blt.Proposer(): blt.Transactions(),
	}

	roundVote := NewRoundVote(blt)
	voted := map[string]*RoundVote{
		blt.Proposer(): roundVote,
	}

	return &RunningRound{
		VotingBasis:  blt.VotingBasis(),
		Proposer:     proposer,
		Transactions: transactions,
		Voted:        voted,
		Ballots:      []ballot.Ballot{blt},
	}, nil
}

func (rr *RunningRound) RoundVote(proposer string) (rv *RoundVote, err error) {
	var found bool
	rv, found = rr.Voted[proposer]
	if !found {
		err = errors.RoundVoteNotFound
		return
	}
	return
}

func (rr *RunningRound) IsVoted(ballot ballot.Ballot) bool {
	rr.RLock()
	defer rr.RUnlock()
	if roundVote, found := rr.Voted[ballot.Proposer()]; !found {
		return false
	} else {
		return roundVote.IsVoted(ballot)
	}
}

func (rr *RunningRound) Vote(ballot ballot.Ballot) {
	rr.Lock()
	defer rr.Unlock()

	if _, found := rr.Voted[ballot.Proposer()]; !found {
		rr.Voted[ballot.Proposer()] = NewRoundVote(ballot)
	} else {
		rr.Voted[ballot.Proposer()].Vote(ballot)
	}

	rr.Ballots = append(rr.Ballots, ballot)
}
