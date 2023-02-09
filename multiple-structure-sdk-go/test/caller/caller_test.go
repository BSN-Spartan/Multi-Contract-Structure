package caller

import (
	"fmt"
	"math/big"

	"multiContractCallStructure/abi/Caller"
	"multiContractCallStructure/eth"
	"multiContractCallStructure/log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const NodeUrl = ""
const NodeWsUrl = ""

// caller合约地址
var callerAddress = common.HexToAddress("")

// receiver合约地址
var receiverAddr = common.HexToAddress("")

const PrivateKey = ""

var x = big.NewInt(1)

func TestCallSetX(t *testing.T) {
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := Caller.NewCaller(callerAddress, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	tx, err := instance.CallSetX(auth, receiverAddr, x)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Println("tx Hash:", tx.Hash().String())
}

func TestDelegatecallSetX(t *testing.T) {
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := Caller.NewCaller(callerAddress, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	tx, err := instance.DelegatecallSetX(auth, receiverAddr, x)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Println("tx Hash:", tx.Hash().String())
}

// 该方法查询X值，只有调用了DelegatecallSetX才会有返回值
func TestGetX(t *testing.T) {
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := Caller.NewCaller(callerAddress, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	x, err := instance.GetX(nil)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Println("x value:", x)
}

func TestEvent(t *testing.T) {
	var callerResponse = make(chan *Caller.CallerResponse, 100)
	wsCli, err := ethclient.Dial(NodeWsUrl)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	wsInstance, err := Caller.NewCaller(callerAddress, wsCli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	go func() {
	hear:
		transfer, err := wsInstance.CallerFilterer.WatchResponse(nil, callerResponse)
		if err != nil {
			log.Errorf(err.Error())
			goto hear
		}
		defer transfer.Unsubscribe()
		select {
		case errs := <-transfer.Err():
			log.Warn(errs)
			goto hear
		}
	}()
	for {
		select {
		case event := <-callerResponse:
			fmt.Println("success:", event.Success)
			fmt.Println("data:", event.Data)
		}
	}
}
