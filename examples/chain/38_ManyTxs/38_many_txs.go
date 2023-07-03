package main

import (
	"fmt"
	"os"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/shopspring/decimal"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

const (
	maxGas                = 3_000_000
	txToMake              = 200
	spotOrdersPerTx       = 10
	derivativeOrdersPerTx = 10
	sleepBetweenTxs       = 100 * time.Millisecond
)

func makeLimitOrderDerivative(
	chainClient chainclient.ChainClient,
	network common.Network,
	senderAddress cosmtypes.AccAddress,
	priceRaw string, amountRaw float64,
) *exchangetypes.DerivativeOrder {
	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	// INJ/USDT PERP
	marketId := "0x9b9980167ecc3645ff1a5517886652d94a0825e54a77d2057cbbe3ebee015963"
	amount := decimal.NewFromFloat(amountRaw)
	price := cosmtypes.MustNewDecFromStr(priceRaw)
	leverage := cosmtypes.MustNewDecFromStr("1")

	order := chainClient.DerivativeOrder(defaultSubaccountID, network, &chainclient.DerivativeOrderData{
		OrderType:    exchangetypes.OrderType_BUY_PO,
		Quantity:     amount,
		Price:        price,
		Leverage:     leverage,
		FeeRecipient: senderAddress.String(),
		MarketId:     marketId,
		IsReduceOnly: false,
	})

	return order
}

func makeLimitOrderSpot(
	chainClient chainclient.ChainClient,
	network common.Network,
	senderAddress cosmtypes.AccAddress,
	priceRaw float64, amountRaw float64,
) *exchangetypes.SpotOrder {
	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	// INJ/USDT
	marketId := "0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"
	amount := decimal.NewFromFloat(amountRaw)
	price := decimal.NewFromFloat(priceRaw)

	order := chainClient.SpotOrder(defaultSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_BUY_PO,
		Quantity:     amount,
		Price:        price,
		FeeRecipient: senderAddress.String(),
		MarketId:     marketId,
	})

	return order
}

func main() {
	//network := common.LoadNetwork("mainnet", "lb")

	// local hosted node
	network := common.Network{
		LcdEndpoint:          "http://localhost:10337",
		TmEndpoint:           "http://localhost:26657",
		ChainGrpcEndpoint:    "tcp://localhost:9900",
		ExchangeGrpcEndpoint: "tcp://localhost:9910",
		ExplorerGrpcEndpoint: "tcp://k8s.global.mainnet.explorer.grpc.injective.network:443",
		ChainId:              "injective-1",
		Fee_denom:            "inj",
		Name:                 "mainnet",
	}

	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	injAddress := os.Getenv("INJ_ADDRESS")
	privKey := os.Getenv("PRIV_KEY")

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		"",
		"",
		"",
		injAddress,
		"",
		privKey,
		false,
	)
	if err != nil {
		panic(err)
	}

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)
	if err != nil {
		panic(err)
	}

	clientCtx.Simulate = false

	clientCtx.WithNodeURI(network.TmEndpoint)
	clientCtx = clientCtx.WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)
	if err != nil {
		panic(err)
	}

	msg := new(exchangetypes.MsgBatchUpdateOrders)
	msg.Sender = senderAddress.String()
	msg.SubaccountId = chainClient.DefaultSubaccount(senderAddress).Hex()

	smarketIds := []string{
		"0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0",
	}
	dmarketIds := []string{
		"0x9b9980167ecc3645ff1a5517886652d94a0825e54a77d2057cbbe3ebee015963",
	}

	msg.SpotMarketIdsToCancelAll = smarketIds
	msg.DerivativeMarketIdsToCancelAll = dmarketIds

	for i := 0; i < spotOrdersPerTx; i++ {
		msg.SpotOrdersToCreate = append(
			msg.SpotOrdersToCreate,
			makeLimitOrderSpot(chainClient, network, senderAddress, 1.5+float64(i)*0.1, 0.007),
		)
	}

	for i := 0; i < derivativeOrdersPerTx; i++ {
		basePrice := 1_500_000
		multiplier := 100_000
		price := basePrice + i*multiplier

		msg.DerivativeOrdersToCreate = append(msg.DerivativeOrdersToCreate, makeLimitOrderDerivative(
			chainClient, network, senderAddress, fmt.Sprintf("%d", price), 0.007,
		))
	}

	start := time.Now()
	txCounter := 0

	for i := 0; i < txToMake; i++ {
		chainClient.SetGasWanted(maxGas)
		r, err := chainClient.AsyncBroadcastMsg(msg)
		if err != nil {
			panic(err)
		}
		if r.TxResponse.RawLog != "[]" {
			panic(r.TxResponse.RawLog)
		}

		time.Sleep(sleepBetweenTxs)

		txCounter++
	}

	duration := time.Since(start)
	fmt.Println("The total duration was", duration)
	fmt.Println("txCounter:", txCounter)
}
