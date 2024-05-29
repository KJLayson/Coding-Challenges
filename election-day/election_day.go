package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if counter == nil {
        return 0
    } else {
        var vote int = *counter
    	return vote
    }

}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	result := &ElectionResult{Name: candidateName, Votes: votes}

    return result
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	var name string = result.Name
    var votes int = result.Votes

    return fmt.Sprintf("%s (%d)", name, votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	_, exists := results[candidate]
    if exists {
        results[candidate] -= 1
    }
}
