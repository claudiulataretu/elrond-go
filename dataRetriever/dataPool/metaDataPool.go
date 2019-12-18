package dataPool

import (
	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/storage"
)

type metaDataPool struct {
	miniBlocks           storage.Cacher
	shardHeaders         dataRetriever.HeadersPool
	transactions         dataRetriever.ShardedDataCacherNotifier
	unsignedTransactions dataRetriever.ShardedDataCacherNotifier
	currBlockTxs         dataRetriever.TransactionCacher
}

// NewMetaDataPool creates a data pools holder object
func NewMetaDataPool(
	miniBlocks storage.Cacher,
	headersCacher dataRetriever.HeadersPool,
	transactions dataRetriever.ShardedDataCacherNotifier,
	unsignedTransactions dataRetriever.ShardedDataCacherNotifier,
	currBlockTxs dataRetriever.TransactionCacher,
) (*metaDataPool, error) {
	if miniBlocks == nil || miniBlocks.IsInterfaceNil() {
		return nil, dataRetriever.ErrNilMiniBlockHashesPool
	}
	if headersCacher == nil || headersCacher.IsInterfaceNil() {
		return nil, dataRetriever.ErrNilShardHeaderPool
	}
	if transactions == nil || transactions.IsInterfaceNil() {
		return nil, dataRetriever.ErrNilTxDataPool
	}
	if unsignedTransactions == nil || unsignedTransactions.IsInterfaceNil() {
		return nil, dataRetriever.ErrNilUnsignedTransactionPool
	}
	if currBlockTxs == nil || currBlockTxs.IsInterfaceNil() {
		return nil, dataRetriever.ErrNilCurrBlockTxs
	}

	return &metaDataPool{
		miniBlocks:           miniBlocks,
		shardHeaders:         headersCacher,
		transactions:         transactions,
		unsignedTransactions: unsignedTransactions,
		currBlockTxs:         currBlockTxs,
	}, nil
}

// CurrentBlockTxs returns the holder for current block transactions
func (mdp *metaDataPool) CurrentBlockTxs() dataRetriever.TransactionCacher {
	return mdp.currBlockTxs
}

// MiniBlocks returns the holder for meta mini block hashes
func (mdp *metaDataPool) MiniBlocks() storage.Cacher {
	return mdp.miniBlocks
}

// Headers returns the holder for shard headers
func (mdp *metaDataPool) Headers() dataRetriever.HeadersPool {
	return mdp.shardHeaders
}

// Transactions returns the holder for transactions which interact with the metachain
func (mdp *metaDataPool) Transactions() dataRetriever.ShardedDataCacherNotifier {
	return mdp.transactions
}

// UnsignedTransactions returns the holder for unsigned transactions which are generated by the metachain
func (mdp *metaDataPool) UnsignedTransactions() dataRetriever.ShardedDataCacherNotifier {
	return mdp.unsignedTransactions
}

// IsInterfaceNil returns true if there is no value under the interface
func (mdp *metaDataPool) IsInterfaceNil() bool {
	if mdp == nil {
		return true
	}
	return false
}
