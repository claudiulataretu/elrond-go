package presenter

import "github.com/ElrondNetwork/elrond-go/core"

// GetAppVersion will return application version
func (psh *PresenterStatusHandler) GetAppVersion() string {
	return psh.getFromCacheAsString(core.MetricAppVersion)
}

// GetNodeType will return type of node
func (psh *PresenterStatusHandler) GetNodeType() string {
	return psh.getFromCacheAsString(core.MetricNodeType)
}

// GetPublicKeyTxSign will return node public key for sign transaction
func (psh *PresenterStatusHandler) GetPublicKeyTxSign() string {
	return psh.getFromCacheAsString(core.MetricPublicKeyTxSign)
}

// GetPublicKeyBlockSign will return node public key for sign blocks
func (psh *PresenterStatusHandler) GetPublicKeyBlockSign() string {
	return psh.getFromCacheAsString(core.MetricPublicKeyBlockSign)
}

// GetShardId will return shard Id of node
func (psh *PresenterStatusHandler) GetShardId() uint64 {
	return psh.getFromCacheAsUint64(core.MetricShardId)
}

// GetCountConsensus will return count of how many times node was in consensus group
func (psh *PresenterStatusHandler) GetCountConsensus() uint64 {
	return psh.getFromCacheAsUint64(core.MetricCountConsensus)
}

// GetCountLeader will return count of how many time node was leader in consensus group
func (psh *PresenterStatusHandler) GetCountLeader() uint64 {
	return psh.getFromCacheAsUint64(core.MetricCountLeader)
}

// GetCountAcceptedBlocks will return count of how many accepted blocks was proposed by the node
func (psh *PresenterStatusHandler) GetCountAcceptedBlocks() uint64 {
	return psh.getFromCacheAsUint64(core.MetricCountAcceptedBlocks)
}