package state

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/line/ostracon/crypto"

	tmstate "github.com/line/ostracon/proto/ostracon/state"
	tmproto "github.com/line/ostracon/proto/ostracon/types"
	tmversion "github.com/line/ostracon/proto/ostracon/version"
	"github.com/line/ostracon/types"
	tmtime "github.com/line/ostracon/types/time"
	"github.com/line/ostracon/version"
)

var (
	// database keys
	stateKey = []byte("stateKey")
)

//-----------------------------------------------------------------------------

// InitStateVersion sets the Consensus.Block and Software versions,
// but leaves the Consensus.App version blank.
// The Consensus.App version will be set during the Handshake, once
// we hear from the app what protocol version it is running.
var InitStateVersion = tmstate.Version{
	Consensus: tmversion.Consensus{
		Block: version.BlockProtocol,
		App:   0,
	},
	Software: version.TMCoreSemVer,
}

//-----------------------------------------------------------------------------

// State is a short description of the latest committed block of the Tendermint consensus.
// It keeps all information necessary to validate new blocks,
// including the last validator set and the consensus params.
// All fields are exposed so the struct can be easily serialized,
// but none of them should be mutated directly.
// Instead, use state.Copy() or state.NextState(...).
// NOTE: not goroutine-safe.
type State struct {
	Version tmstate.Version

	// immutable
	ChainID       string
	InitialHeight int64 // should be 1, not 0, when starting from height 1
	VoterParams   *types.VoterParams

	// LastBlockHeight=0 at genesis (ie. block(H=0) does not exist)
	LastBlockHeight int64
	LastBlockID     types.BlockID
	LastBlockTime   time.Time

	// vrf hash from proof
	LastProofHash []byte

	// LastVoters is used to validate block.LastCommit.
	// Validators are persisted to the database separately every time they change,
	// so we can query for historical validator sets.
	// Note that if s.LastBlockHeight causes a valset change,
	// we set s.LastHeightValidatorsChanged = s.LastBlockHeight + 1 + 1
	// Extra +1 due to nextValSet delay.
	NextValidators              *types.ValidatorSet
	Validators                  *types.ValidatorSet
	Voters                      *types.VoterSet
	LastVoters                  *types.VoterSet
	LastHeightValidatorsChanged int64

	// Consensus parameters used for validating blocks.
	// Changes returned by EndBlock and updated after Commit.
	ConsensusParams                  tmproto.ConsensusParams
	LastHeightConsensusParamsChanged int64

	// Merkle root of the results from executing prev block
	LastResultsHash []byte

	// the latest AppHash we've received from calling abci.Commit()
	AppHash []byte
}

func (state State) MakeHashMessage(round int32) []byte {
	return types.MakeRoundHash(state.LastProofHash, state.LastBlockHeight, round)
}

// Copy makes a copy of the State for mutating.
func (state State) Copy() State {

	return State{
		Version:       state.Version,
		ChainID:       state.ChainID,
		InitialHeight: state.InitialHeight,
		VoterParams:   state.VoterParams,

		LastBlockHeight: state.LastBlockHeight,
		LastBlockID:     state.LastBlockID,
		LastBlockTime:   state.LastBlockTime,

		LastProofHash: state.LastProofHash,

		NextValidators:              state.NextValidators.Copy(),
		Validators:                  state.Validators.Copy(),
		Voters:                      state.Voters.Copy(),
		LastVoters:                  state.LastVoters.Copy(),
		LastHeightValidatorsChanged: state.LastHeightValidatorsChanged,

		ConsensusParams:                  state.ConsensusParams,
		LastHeightConsensusParamsChanged: state.LastHeightConsensusParamsChanged,

		AppHash: state.AppHash,

		LastResultsHash: state.LastResultsHash,
	}
}

// Equals returns true if the States are identical.
func (state State) Equals(state2 State) bool {
	sbz, s2bz := state.Bytes(), state2.Bytes()
	return bytes.Equal(sbz, s2bz)
}

// Bytes serializes the State using protobuf.
// It panics if either casting to protobuf or serialization fails.
func (state State) Bytes() []byte {
	sm, err := state.ToProto()
	if err != nil {
		panic(err)
	}
	bz, err := sm.Marshal()
	if err != nil {
		panic(err)
	}
	return bz
}

// IsEmpty returns true if the State is equal to the empty State.
func (state State) IsEmpty() bool {
	return state.Validators == nil // XXX can't compare to Empty
}

// ToProto takes the local state type and returns the equivalent proto type
func (state *State) ToProto() (*tmstate.State, error) {
	if state == nil {
		return nil, errors.New("state is nil")
	}

	sm := new(tmstate.State)

	sm.Version = state.Version
	sm.ChainID = state.ChainID
	sm.InitialHeight = state.InitialHeight
	sm.LastBlockHeight = state.LastBlockHeight
	sm.VoterParams = state.VoterParams.ToProto()

	sm.LastBlockID = state.LastBlockID.ToProto()
	sm.LastBlockTime = state.LastBlockTime
	vals, err := state.Validators.ToProto()
	if err != nil {
		return nil, err
	}
	sm.Validators = vals

	voters, err := state.Voters.ToProto()
	if err != nil {
		return nil, err
	}
	sm.Voters = voters

	nVals, err := state.NextValidators.ToProto()
	if err != nil {
		return nil, err
	}
	sm.NextValidators = nVals

	if state.LastBlockHeight >= 1 { // At Block 1 LastVoters is nil
		lVoters, err := state.LastVoters.ToProto()
		if err != nil {
			return nil, err
		}
		sm.LastVoters = lVoters
	} else {
		sm.LastVoters = nil
	}

	sm.LastHeightValidatorsChanged = state.LastHeightValidatorsChanged
	sm.ConsensusParams = state.ConsensusParams
	sm.LastHeightConsensusParamsChanged = state.LastHeightConsensusParamsChanged
	sm.LastResultsHash = state.LastResultsHash
	sm.AppHash = state.AppHash

	sm.LastProofHash = state.LastProofHash

	return sm, nil
}

// StateFromProto takes a state proto message & returns the local state type
func StateFromProto(pb *tmstate.State) (*State, error) { //nolint:golint
	if pb == nil {
		return nil, errors.New("nil State")
	}

	state := new(State)

	state.Version = pb.Version
	state.ChainID = pb.ChainID
	state.InitialHeight = pb.InitialHeight
	state.VoterParams = types.VoterParamsFromProto(pb.VoterParams)

	bi, err := types.BlockIDFromProto(&pb.LastBlockID)
	if err != nil {
		return nil, err
	}
	state.LastBlockID = *bi
	state.LastBlockHeight = pb.LastBlockHeight
	state.LastBlockTime = pb.LastBlockTime

	vals, err := types.ValidatorSetFromProto(pb.Validators)
	if err != nil {
		return nil, err
	}
	state.Validators = vals

	voters, err := types.VoterSetFromProto(pb.Voters)
	if err != nil {
		return nil, err
	}
	state.Voters = voters

	nVals, err := types.ValidatorSetFromProto(pb.NextValidators)
	if err != nil {
		return nil, err
	}
	state.NextValidators = nVals

	if state.LastBlockHeight >= 1 { // At Block 1 LastVoters is nil
		lVoters, err := types.VoterSetFromProto(pb.LastVoters)
		if err != nil {
			return nil, err
		}
		state.LastVoters = lVoters
	} else {
		state.LastVoters = &types.VoterSet{Voters: []*types.Validator{}} // XXX Need to be the same
	}

	state.LastHeightValidatorsChanged = pb.LastHeightValidatorsChanged
	state.ConsensusParams = pb.ConsensusParams
	state.LastHeightConsensusParamsChanged = pb.LastHeightConsensusParamsChanged
	state.LastResultsHash = pb.LastResultsHash
	state.AppHash = pb.AppHash

	state.LastProofHash = pb.LastProofHash

	return state, nil
}

//------------------------------------------------------------------------
// Create a block from the latest state

// MakeBlock builds a block from the current state with the given txs, commit,
// and evidence. Note it also takes a proposerAddress because the state does not
// track rounds, and hence does not know the correct proposer. TODO: fix this!
func (state State) MakeBlock(
	height int64,
	txs []types.Tx,
	commit *types.Commit,
	evidence []types.Evidence,
	proposerAddress []byte,
	round int32,
	proof crypto.Proof,
) (*types.Block, *types.PartSet) {

	// Build base block with block data.
	block := types.MakeBlock(height, txs, commit, evidence)

	// Set time.
	var timestamp time.Time
	if height == state.InitialHeight {
		timestamp = state.LastBlockTime // genesis time
	} else {
		timestamp = MedianTime(commit, state.LastVoters)
	}

	// Fill rest of header with state data.
	block.Header.Populate(
		state.Version.Consensus, state.ChainID,
		timestamp, state.LastBlockID,
		state.Voters.Hash(), state.Validators.Hash(), state.NextValidators.Hash(),
		types.HashConsensusParams(state.ConsensusParams), state.AppHash, state.LastResultsHash,
		proposerAddress,
		round,
		proof,
	)

	return block, block.MakePartSet(types.BlockPartSizeBytes)
}

// MedianTime computes a median time for a given Commit (based on Timestamp field of votes messages) and the
// corresponding validator set. The computed time is always between timestamps of
// the votes sent by honest processes, i.e., a faulty processes can not arbitrarily increase or decrease the
// computed value.
func MedianTime(commit *types.Commit, voters *types.VoterSet) time.Time {
	weightedTimes := make([]*tmtime.WeightedTime, len(commit.Signatures))
	totalVotingPower := int64(0)

	for i, commitSig := range commit.Signatures {
		if commitSig.Absent() {
			continue
		}
		_, validator := voters.GetByAddress(commitSig.ValidatorAddress)
		// If there's no condition, TestValidateBlockCommit panics; not needed normally.
		if validator != nil {
			totalVotingPower += validator.VotingPower
			weightedTimes[i] = tmtime.NewWeightedTime(commitSig.Timestamp, validator.VotingPower)
		}
	}

	return tmtime.WeightedMedian(weightedTimes, totalVotingPower)
}

//------------------------------------------------------------------------
// Genesis

// MakeGenesisStateFromFile reads and unmarshals state from the given
// file.
//
// Used during replay and in tests.
func MakeGenesisStateFromFile(genDocFile string) (State, error) {
	genDoc, err := MakeGenesisDocFromFile(genDocFile)
	if err != nil {
		return State{}, err
	}
	return MakeGenesisState(genDoc)
}

// MakeGenesisDocFromFile reads and unmarshals genesis doc from the given file.
func MakeGenesisDocFromFile(genDocFile string) (*types.GenesisDoc, error) {
	genDocJSON, err := ioutil.ReadFile(genDocFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't read GenesisDoc file: %v", err)
	}
	genDoc, err := types.GenesisDocFromJSON(genDocJSON)
	if err != nil {
		return nil, fmt.Errorf("error reading GenesisDoc: %v", err)
	}
	return genDoc, nil
}

// MakeGenesisState creates state from types.GenesisDoc.
func MakeGenesisState(genDoc *types.GenesisDoc) (State, error) {
	err := genDoc.ValidateAndComplete()
	if err != nil {
		return State{}, fmt.Errorf("error in genesis file: %v", err)
	}

	var validatorSet, nextValidatorSet *types.ValidatorSet
	if genDoc.Validators == nil {
		validatorSet = types.NewValidatorSet(nil)
		nextValidatorSet = types.NewValidatorSet(nil)
	} else {
		validators := make([]*types.Validator, len(genDoc.Validators))
		for i, val := range genDoc.Validators {
			validators[i] = types.NewValidator(val.PubKey, val.Power)
		}
		validatorSet = types.NewValidatorSet(validators)
		nextValidatorSet = types.NewValidatorSet(validators)
	}

	return State{
		Version:       InitStateVersion,
		ChainID:       genDoc.ChainID,
		InitialHeight: genDoc.InitialHeight,
		VoterParams:   genDoc.VoterParams,

		LastBlockHeight: 0,
		LastBlockID:     types.BlockID{},
		LastBlockTime:   genDoc.GenesisTime,

		// genesis block use the hash of GenesisDoc instead for the `LastProofHash`
		LastProofHash: genDoc.Hash(),

		NextValidators:              nextValidatorSet,
		Validators:                  validatorSet,
		Voters:                      types.SelectVoter(validatorSet, genDoc.Hash(), genDoc.VoterParams),
		LastVoters:                  &types.VoterSet{Voters: []*types.Validator{}}, // LastVoters don't exist if LastBlockHeight==0; XXX Need to be the same
		LastHeightValidatorsChanged: 1,

		ConsensusParams:                  *genDoc.ConsensusParams,
		LastHeightConsensusParamsChanged: genDoc.InitialHeight,

		AppHash: genDoc.AppHash,
	}, nil
}
