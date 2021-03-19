package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

func GetCall(ethApi, abiJson, methodName string, contractAddr common.Address, blockNumber *big.Int,
	retValue interface{}, args ...interface{}) error {
	erc20Abi, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return err
	}
	packDataBts, err := erc20Abi.Pack(methodName, args...)
	if err != nil {
		return err
	}
	client, err := ethclient.Dial(ethApi)
	if err != nil {
		return err
	}
	ethCallMsg := ethereum.CallMsg{
		From: contractAddr,
		To:   &contractAddr,
		Data: packDataBts}

	resBts, err := client.CallContract(context.Background(), ethCallMsg, blockNumber)
	if err != nil {
		return fmt.Errorf("client.CallContract err %s", err)
	}
	if len(resBts) == 0 {
		return errors.New("client.CallContract return 0 bytes")
	}
	return erc20Abi.UnpackIntoInterface(retValue, methodName, resBts)
}

func GetOpts(hexKey,ethApi string) (*bind.TransactOpts,*ethclient.Client,error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return  nil,nil,err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil,nil,fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	client, err := ethclient.Dial(ethApi)
	if err != nil {
		return nil,nil,err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil,nil,err
	}

	opts,err := bind.NewKeyedTransactorWithChainID(privateKey,chainID)
	if err != nil {
		return nil,nil,err
	}


	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil,nil,err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil,nil,err
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)     // in wei
	opts.GasLimit = uint64(5500000) // in units
	opts.GasPrice = gasPrice
	return  opts,client,nil
}
