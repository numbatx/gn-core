package indexer

import (
	"time"

	"github.com/numbatx/gn-core/data"
)

// ArgsSaveBlockData will contains all information that are needed to save block data
type ArgsSaveBlockData struct {
	HeaderHash             []byte
	Body                   data.BodyHandler
	Header                 data.HeaderHandler
	SignersIndexes         []uint64
	NotarizedHeadersHashes []string
	HeaderGasConsumption   HeaderGasConsumption
	TransactionsPool       *Pool
}

// HeaderGasConsumption holds the data needed to save the gas consumption of a header
type HeaderGasConsumption struct {
	GasProvided    uint64
	GasRefunded    uint64
	GasPenalized   uint64
	MaxGasPerBlock uint64
}

// Pool will holds all types of transaction
type Pool struct {
	Txs      map[string]data.TransactionHandler
	Scrs     map[string]data.TransactionHandler
	Rewards  map[string]data.TransactionHandler
	Invalid  map[string]data.TransactionHandler
	Receipts map[string]data.TransactionHandler
	Logs     map[string]data.LogHandler
}

// ValidatorRatingInfo is a structure containing validator rating information
type ValidatorRatingInfo struct {
	PublicKey string
	Rating    float32
}

// RoundInfo is a structure containing block signers and shard id
type RoundInfo struct {
	Index            uint64
	SignersIndexes   []uint64
	BlockWasProposed bool
	ShardId          uint32
	Epoch            uint32
	Timestamp        time.Duration
}
