package main

import (
	"context"
	"encoding/hex"
	"os"
	"sync"
	"time"

	vapi "go.vocdoni.io/dvote/api"
	"go.vocdoni.io/dvote/apiclient"
	"go.vocdoni.io/dvote/crypto/ethereum"
	"go.vocdoni.io/dvote/log"
)

func init() {
	ops["plaintextelection"] = operation{
		test:        &E2EPlaintextElection{},
		description: "Publishes a census and a non-anonymous, non-secret election, emits N votes and verifies the results",
		example:     os.Args[0] + " --operation=plaintextelection --votes=1000",
	}
}

var _ VochainTest = (*E2EPlaintextElection)(nil)

type E2EPlaintextElection struct {
	e2eElection
}

func (t *E2EPlaintextElection) Setup(api *apiclient.HTTPclient, c *config) error {
	t.api = api
	t.config = c

	ed := newTestElectionDescription()
	ed.ElectionType = vapi.ElectionType{
		Autostart:     true,
		Interruptible: true,
	}
	ed.VoteType = vapi.VoteType{MaxVoteOverwrites: 1}
	ed.Census = vapi.CensusTypeDescription{Type: vapi.CensusTypeWeighted}

	if err := t.setupElection(ed); err != nil {
		return err
	}

	log.Debugf("election details: %+v", *t.election)
	return nil
}

func (t *E2EPlaintextElection) Teardown() error {
	// nothing to do here
	return nil
}

func (t *E2EPlaintextElection) Run() error {
	c := t.config
	api := t.api

	// Send the votes (parallelized)
	startTime := time.Now()
	wg := sync.WaitGroup{}
	voteAccounts := func(accounts []*ethereum.SignKeys, wg *sync.WaitGroup) {
		defer wg.Done()
		log.Infof("sending %d votes", len(accounts))

		votesSent := 0
		contextDeadlines := 0
		for i, acc := range accounts {
			ctxDeadline, err := t.sendVote(acc, []int{i % 2}, nil)
			if err != nil {
				log.Error(err)
				break
			}
			contextDeadlines += ctxDeadline
			votesSent++
		}
		log.Infof("successfully sent %d votes... got %d HTTP errors", votesSent, contextDeadlines)
		time.Sleep(time.Second * 4)
	}

	pcount := c.nvotes / c.parallelCount
	for i := 0; i < len(t.voterAccounts); i += pcount {
		end := i + pcount
		if end > len(t.voterAccounts) {
			end = len(t.voterAccounts)
		}
		wg.Add(1)
		go voteAccounts(t.voterAccounts[i:end], &wg)
	}

	wg.Wait()
	log.Infof("%d votes submitted successfully, took %s (%d votes/second)",
		c.nvotes, time.Since(startTime), int(float64(c.nvotes)/time.Since(startTime).Seconds()))

	// Wait for all the votes to be verified
	log.Infof("waiting for all the votes to be registered...")
	for {
		count, err := api.ElectionVoteCount(t.election.ElectionID)
		if err != nil {
			log.Warn(err)
		}
		if count == uint32(c.nvotes) {
			break
		}
		time.Sleep(time.Second * 5)
		log.Infof("verified %d/%d votes", count, c.nvotes)
		if time.Since(startTime) > c.timeout {
			log.Fatalf("timeout waiting for votes to be registered")
		}
	}

	log.Infof("%d votes registered successfully, took %s (%d votes/second)",
		c.nvotes, time.Since(startTime), int(float64(c.nvotes)/time.Since(startTime).Seconds()))

	// Set the account back to the organization account
	if err := api.SetAccount(hex.EncodeToString(c.accountKeys[0].PrivateKey())); err != nil {
		return err
	}

	// End the election by setting the status to ENDED
	log.Infof("ending election...")
	hash, err := api.SetElectionStatus(t.election.ElectionID, "ENDED")
	if err != nil {
		return err
	}

	// Check the election status is actually ENDED
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*40)
	defer cancel()
	if _, err := api.WaitUntilTxIsMined(ctx, hash); err != nil {
		log.Fatalf("gave up waiting for tx %s to be mined: %s", hash, err)
	}

	t.election, err = api.Election(t.election.ElectionID)
	if err != nil {
		return err
	}
	if t.election.Status != "ENDED" {
		log.Fatal("election status is not ENDED")
	}
	log.Infof("election %s status is ENDED", t.election.ElectionID.String())

	// Wait for the election to be in RESULTS state
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*300)
	defer cancel()
	t.election, err = api.WaitUntilElectionStatus(ctx, t.election.ElectionID, "RESULTS")
	if err != nil {
		return err
	}
	log.Infof("election %s status is RESULTS", t.election.ElectionID.String())
	log.Infof("election results: %v", t.election.Results)

	return nil
}