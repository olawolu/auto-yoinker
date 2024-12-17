package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/olawolu/yoinker/yoink"
	"github.com/robfig/cron/v3"
)

const (
	yoinkContractAddress = "0x4bBFD120d9f352A0BEd7a014bd67913a2007a878"
)

func main() {
	godotenv.Load()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	client, err := newBlockchainClient(context.Background())
	if err != nil {
		err = fmt.Errorf("error creating blockchain client: %v", err)
		log.Println(err)
		return
	}
	defer client.Close()

	runner := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DefaultLogger),
	))
	defer runner.Stop()

	runPeriod := os.Getenv("RUN_PERIOD")
	wallet := os.Getenv("WALLET")
	interval := fmt.Sprintf("@every %v", runPeriod)
	_, err = runner.AddFunc(interval, func() {
		tx, err := yoinker(client, wallet)
		if err != nil {
			log.Printf("error yoinking: %v", err)
		}

		log.Printf("yoinked: %s", tx.Hash().Hex())
	})
	if err != nil {
		log.Fatalf("error adding function to cron: %v", err)
	}

	runner.Start()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		fmt.Println("\nshutting down yoinker!")
	}()
	wg.Wait()
}

func newBlockchainClient(ctx context.Context) (*ethclient.Client, error) {
	baseRPC := os.Getenv("BASE_RPC")

	client, err := ethclient.DialContext(ctx, baseRPC)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func yoinker(client *ethclient.Client, walletKey string) (*types.Transaction, error) {
	yoinkContract := common.HexToAddress(yoinkContractAddress)
	yoinkInstance, err := yoink.NewYoink(yoinkContract, client)
	if err != nil {
		err = fmt.Errorf("error creating contract instance: %v", err)
		return nil, err
	}

	pk, err := crypto.HexToECDSA(walletKey)
	if err != nil {
		err = fmt.Errorf("error converting private key: %v", err)
		return nil, err
	}
	// Call the contract
	transactionOpts, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(8453))
	if err != nil {
		err = fmt.Errorf("error creating transaction options: %v", err)
		return nil, err
	}

	fmt.Println("Yoinking...", transactionOpts.From.Hex())
	tx, err := yoinkInstance.Yoink(transactionOpts)
	if err != nil {
		err = fmt.Errorf("error calling contract: %v", err)
		return nil, err
	}

	txHash := tx.Hash()
	log.Printf("Transaction sent: %s", txHash.Hex())
	return tx, nil
}
