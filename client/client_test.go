package client

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

const ConfirmDataStoreEventABI = "ConfirmDataStore(uint32,bytes32)"

var ConfirmDataStoreEventABIHash = crypto.Keccak256Hash([]byte(ConfirmDataStoreEventABI))

const DataLayrServiceManagerAddr = "0x5BD63a7ECc13b955C4F57e3F12A64c10263C14c1"

func TestEthClient_GetTxReceiptByHash(t *testing.T) {
	fmt.Println("start...........")
	clint, err := newEthClients("https://rpc.payload.de")
	if err != nil {
		fmt.Println("connect ethereum fail", "err", err)
		return
	}
	txReceipt, err := clint.GetTxReceiptByHash("0xbe755f52cf29c4fbdaae2ee236c9922dce6d7e365a39d5b666518662d4d01a3c")
	abiUint32, err := abi.NewType("uint32", "uint32", nil)
	if err != nil {
		fmt.Println("Abi new uint32 type error", "err", err)
		return
	}
	abiBytes32, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		fmt.Println("Abi new bytes32 type error", "err", err)
		return
	}
	confirmDataStoreArgs := abi.Arguments{
		{
			Name:    "dataStoreId",
			Type:    abiUint32,
			Indexed: false,
		}, {
			Name:    "headerHash",
			Type:    abiBytes32,
			Indexed: false,
		},
	}
	var dataStoreData = make(map[string]interface{})
	for _, rLog := range txReceipt.Logs {
		fmt.Println(rLog.Address.String())
		if strings.ToLower(rLog.Address.String()) != strings.ToLower(DataLayrServiceManagerAddr) {
			continue
		}
		if rLog.Topics[0] != ConfirmDataStoreEventABIHash {
			continue
		}
		if len(rLog.Data) > 0 {
			err := confirmDataStoreArgs.UnpackIntoMap(dataStoreData, rLog.Data)
			if err != nil {
				log.Error("Unpack data into map fail", "err", err)
				continue
			}
			if dataStoreData != nil {
				dataStoreId := dataStoreData["dataStoreId"].(uint32)
				headerHash := dataStoreData["headerHash"]
				fmt.Println(dataStoreId)
				fmt.Println(headerHash)
			}
			return
		}
	}
}

func TestEthClient_GetLogs(t *testing.T) {
	startBlock := big.NewInt(21333483)
	endBlock := big.NewInt(21333485)
	var contractList []common.Address
	addressCm := common.HexToAddress(DataLayrServiceManagerAddr)
	contractList = append(contractList, addressCm)
	clint, err := newEthClients("https://rpc.payload.de")
	if err != nil {
		fmt.Println("connect ethereum fail", "err", err)
		return
	}
	logList, err := clint.GetLogs(startBlock, endBlock, contractList)
	if err != nil {
		fmt.Println("get log fail")
		return
	}
	abiUint32, err := abi.NewType("uint32", "uint32", nil)
	if err != nil {
		fmt.Println("Abi new uint32 type error", "err", err)
		return
	}
	abiBytes32, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		fmt.Println("Abi new bytes32 type error", "err", err)
		return
	}
	confirmDataStoreArgs := abi.Arguments{
		{
			Name:    "dataStoreId",
			Type:    abiUint32,
			Indexed: false,
		}, {
			Name:    "headerHash",
			Type:    abiBytes32,
			Indexed: false,
		},
	}
	var dataStoreData = make(map[string]interface{})
	for _, rLog := range logList {
		fmt.Println(rLog.Address.String())
		if strings.ToLower(rLog.Address.String()) != strings.ToLower(DataLayrServiceManagerAddr) {
			continue
		}
		if rLog.Topics[0] != ConfirmDataStoreEventABIHash {
			continue
		}
		if len(rLog.Data) > 0 {
			err := confirmDataStoreArgs.UnpackIntoMap(dataStoreData, rLog.Data)
			if err != nil {
				log.Error("Unpack data into map fail", "err", err)
				continue
			}
			if dataStoreData != nil {
				dataStoreId := dataStoreData["dataStoreId"].(uint32)
				headerHash := dataStoreData["headerHash"]
				fmt.Println(dataStoreId)
				fmt.Println(headerHash)
			}
			return
		}
	}
}
