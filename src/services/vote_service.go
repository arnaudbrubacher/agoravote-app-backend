package services

import (
	"errors"
	"sync"
)

type Vote struct {
	ID     string
	UserID string
	PostID string
	Value  int // 1 for upvote, -1 for downvote
}

type VoteService struct {
	votes map[string]Vote
	mu    sync.Mutex
}

func NewVoteService() *VoteService {
	return &VoteService{
		votes: make(map[string]Vote),
	}
}

func (vs *VoteService) CastVote(vote Vote) error {
	vs.mu.Lock()
	defer vs.mu.Unlock()

	if _, exists := vs.votes[vote.ID]; exists {
		return errors.New("vote already exists")
	}

	vs.votes[vote.ID] = vote
	return nil
}

func (vs *VoteService) FetchVotes(postID string) []Vote {
	vs.mu.Lock()
	defer vs.mu.Unlock()

	var postVotes []Vote
	for _, vote := range vs.votes {
		if vote.PostID == postID {
			postVotes = append(postVotes, vote)
		}
	}
	return postVotes
}