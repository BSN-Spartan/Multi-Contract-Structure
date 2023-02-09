package logic

import (
	"fmt"
	"math/big"

	"multiContractCallStructure/abi/Logic"
	"multiContractCallStructure/eth"
	"multiContractCallStructure/log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const NodeUrl = ""

// logic合约地址
var Address = common.HexToAddress("")

// data合约地址
var dataContractAddr = common.HexToAddress("")

const PrivateKey = ""

var x = big.NewInt(1)

func TestSetX(t *testing.T) {
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := Logic.NewLogic(Address, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	tx, err := instance.SetX(auth, dataContractAddr, x)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Println("tx Hash:", tx.Hash().String())
}

func TestSetXFromAddress(t *testing.T) {
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := Logic.NewLogic(Address, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	tx, err := instance.SetXFromAddress(auth, dataContractAddr, x)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Println("tx Hash:", tx.Hash().String())
}
