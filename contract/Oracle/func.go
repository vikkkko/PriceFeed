package Oracle

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"ism/contract/client"
	"math/big"
	"strings"
)

func SetDirectPrice(hexKey,ethApi string,contractAddr common.Address,assets []common.Address, prices []*big.Int) (error)  {
	parsed, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return err
	}
	opts,client,err := client.GetOpts(hexKey,ethApi)
	if err != nil {
		return err
	}
	fmt.Println("contractAddr: %s", contractAddr)
	ins := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	tx,err := ins.Transact(opts,"setDirectPrice",assets,prices)
	fmt.Println("tx sent: %s", tx.Hash().Hex())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}