package dao

import (
	"fmt"
	"github.com/everstake/cosmoscan-api/config"
	"github.com/everstake/cosmoscan-api/dao/cache"
	"github.com/everstake/cosmoscan-api/dao/clickhouse"
	"github.com/everstake/cosmoscan-api/dao/filters"
	"github.com/everstake/cosmoscan-api/dao/mysql"
	"github.com/everstake/cosmoscan-api/dmodels"
	"github.com/everstake/cosmoscan-api/smodels"
	"github.com/shopspring/decimal"
	"time"
)

type (
	DAO interface {
		Mysql
		Clickhouse
		Cache
	}
	Mysql interface {
		GetParsers() (parsers []dmodels.Parser, err error)
		GetParser(title string) (parser dmodels.Parser, err error)
		UpdateParser(parser dmodels.Parser) error
		CreateValidators(validators []dmodels.Validator) error
		UpdateValidators(validator dmodels.Validator) error
		CreateAccounts(accounts []dmodels.Account) error
		UpdateAccount(account dmodels.Account) error
		GetAccount(address string) (account dmodels.Account, err error)
		GetAccounts(filter filters.Accounts) (accounts []dmodels.Account, err error)
		GetAccountsTotal(filter filters.Accounts) (total uint64, err error)
		CreateProposals(proposals []dmodels.Proposal) error
		GetProposals(filter filters.Proposals) (proposals []dmodels.Proposal, err error)
		UpdateProposal(proposal dmodels.Proposal) error
	}
	Clickhouse interface {
		CreateBlocks(blocks []dmodels.Block) error
		GetBlocks(filter filters.Blocks) (blocks []dmodels.Block, err error)
		GetAggBlocksCount(filter filters.Agg) (items []smodels.AggItem, err error)
		GetAggBlocksDelay(filter filters.Agg) (items []smodels.AggItem, err error)
		GetAvgBlocksDelay(filter filters.TimeRange) (delay float64, err error)
		GetAggUniqBlockValidators(filter filters.Agg) (items []smodels.AggItem, err error)
		CreateTransactions(transactions []dmodels.Transaction) error
		GetAggOperationsCount(filter filters.Agg) (items []smodels.AggItem, err error)
		GetAggTransactionsFee(filter filters.Agg) (items []smodels.AggItem, err error)
		GetTransactionsFeeVolume(filter filters.TimeRange) (total decimal.Decimal, err error)
		GetTransactionsHighestFee(filter filters.TimeRange) (total decimal.Decimal, err error)
		GetAggTransfersVolume(filter filters.Agg) (items []smodels.AggItem, err error)
		CreateTransfers(transfers []dmodels.Transfer) error
		GetTransferVolume(filter filters.TimeRange) (total decimal.Decimal, err error)
		CreateDelegations(delegations []dmodels.Delegation) error
		GetAggDelegationsVolume(filter filters.Agg) (items []smodels.AggItem, err error)
		GetUndelegationsVolume(filter filters.TimeRange) (total decimal.Decimal, err error)
		GetDelegatorsTotal(filter filters.TimeRange) (total uint64, err error)
		GetMultiDelegatorsTotal(filter filters.TimeRange) (total uint64, err error)
		GetAggUndelegationsVolume(filter filters.Agg) (items []smodels.AggItem, err error)
		CreateDelegatorRewards(rewards []dmodels.DelegatorReward) error
		CreateValidatorRewards(rewards []dmodels.ValidatorReward) error
		CreateProposalDeposits(deposits []dmodels.ProposalDeposit) error
		CreateProposalVotes(votes []dmodels.ProposalVote) error
		GetProposalVotes(filter filters.ProposalVotes) (votes []dmodels.ProposalVote, err error)
		GetProposalVotesTotal(filter filters.ProposalVotes) (total uint64, err error)
		GetAggProposalVotes(filter filters.Agg, id []uint64) (items []smodels.AggItem, err error)
		CreateHistoricalStates(states []dmodels.HistoricalState) error
		GetHistoricalStates(state filters.HistoricalState) (states []dmodels.HistoricalState, err error)
		GetAggHistoricalStatesByField(filter filters.Agg, field string) (items []smodels.AggItem, err error)
		GetActiveAccounts(filter filters.ActiveAccounts) (addresses []string, err error)
		CreateBalanceUpdates(updates []dmodels.BalanceUpdate) error
		GetBalanceUpdate(filter filters.BalanceUpdates) (updates []dmodels.BalanceUpdate, err error)
		CreateJailers(jailers []dmodels.Jailer) error
		GetJailersTotal() (total uint64, err error)
		CreateStats(stats []dmodels.Stat) (err error)
		GetStats(filter filters.Stats) (stats []dmodels.Stat, err error)
	}

	Cache interface {
		CacheSet(key string, data interface{}, duration time.Duration)
		CacheGet(key string) (data interface{}, found bool)
	}

	daoImpl struct {
		Mysql
		Clickhouse
		Cache
	}
)

func NewDAO(cfg config.Config) (DAO, error) {
	mysqlDB, err := mysql.NewDB(cfg.Mysql)
	if err != nil {
		return nil, fmt.Errorf("mysql.NewDB: %s", err.Error())
	}
	ch, err := clickhouse.NewDB(cfg.Clickhouse)
	if err != nil {
		return nil, fmt.Errorf("clickhouse.NewDB: %s", err.Error())
	}
	return daoImpl{
		Mysql:      mysqlDB,
		Clickhouse: ch,
		Cache:      cache.New(),
	}, nil
}
