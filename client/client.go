package client

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	client *ethclient.Client
}

func newEthClients(rpcUrl string) (*EthClient, error) {
	client, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		fmt.Println("dail eth client fail")
		return nil, err
	}
	return &EthClient{
		client: client,
	}, nil
}

func (ec EthClient) GetTxReceiptByHash(txHash string) (*types.Receipt, error) {
	return ec.client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
}

func (ec EthClient) GetLogs(starkBlock, endBlock *big.Int, contractAddressList []common.Address) ([]types.Log, error) {
	filterQuery := ethereum.FilterQuery{FromBlock: starkBlock, ToBlock: endBlock, Addresses: contractAddressList}
	return ec.client.FilterLogs(context.Background(), filterQuery)
}
