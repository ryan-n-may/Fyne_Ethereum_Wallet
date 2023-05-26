package transactions;

import (
	big			"math/big"
	common 		"github.com/ethereum/go-ethereum/common"
	types 		"github.com/ethereum/go-ethereum/core/types"
	context 	"context"
	ethclient 	"github.com/ethereum/go-ethereum/ethclient"
);

type TransactionStruct struct{
	HashHex_		string;
	Value_			string;
	Gas_			uint64;
	GasPrice_		uint64;
	Nonce_			uint64;
	Data_			[]byte;
	ToHex_			string;
	FromHex_		string;
	Receipt_		types.Receipt;

}

type BlockStruct struct{
	Block_                  types.Block;
	BlockNumber_ 			big.Int;
	BlockTime_				uint64;
	BlockDifficulty_		big.Int;
	BlockHash_				string;
	BlockNumTransactions_ 	int;
}

// input NIL to get latest block header
func FetchBlockHeaderByNumber(url string, block_num big.Int) types.Header{
	client, err := ethclient.Dial(url);
	if (err != nil){
		panic(err);
	}
	header, err := client.HeaderByNumber(context.Background(), &block_num);
	if (err != nil){
		panic(err);
	}
	return *header;
}

func FetchBlockByNumber(url string, block_num big.Int) BlockStruct{
	client, err := ethclient.Dial(url);
	if (err != nil){
		panic(err);
	}
	block, err := client.BlockByNumber(context.Background(), &block_num)
	if (err != nil){
		panic(err);
	}
	block_struct := BlockStruct{*block,
						*block.Number(),
						block.Time(),
						*block.Difficulty(),
						block.Hash().Hex(),
						len(block.Transactions())};
	return block_struct;
}

func QueryTransactions(url string, block_num big.Int) [100]TransactionStruct{
	client, err := ethclient.Dial(url);
	if (err != nil){
		panic(err);
	}
	block, err := client.BlockByNumber(context.Background(), &block_num)
	if (err != nil){
		panic(err);
	}

	var transactions [100]TransactionStruct;
	for i := 0; i < 100; i++{
		tx := block.Transactions()[i];

		chainID, err := client.NetworkID(context.Background());
		if (err != nil){
			panic(err);
		}
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash());
		if (err != nil){
			panic(err);
		}
		from, err := types.Sender(types.NewEIP155Signer(chainID), tx); 
		if (err != nil){
			panic(err);
		}
		transaction := TransactionStruct{tx.Hash().Hex(),
							tx.Value().String(),
							tx.Gas(),
							tx.GasPrice().Uint64(),
							tx.Nonce(),
							tx.Data(),
							tx.To().Hex(),
							from.Hex(),
							*receipt};
		
		transactions[i] = transaction;
	}
	return transactions;
}

func TransactionByHash(url string, hashhex string) (transaction TransactionStruct, pending bool){
	client, err := ethclient.Dial(url);
	txHash := common.HexToHash(hashhex);
	tx, pending, err := client.TransactionByHash(context.Background(), txHash)
	if (err != nil){
		panic(err);
	}
	chainID, err := client.NetworkID(context.Background());
	if (err != nil){
		panic(err);
	}
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash());
	if (err != nil){
		panic(err);
	}
	from, err := types.Sender(types.NewEIP155Signer(chainID), tx); 
	if (err != nil){
		panic(err);
	}
	transaction = TransactionStruct{tx.Hash().Hex(),
						tx.Value().String(),
						tx.Gas(),
						tx.GasPrice().Uint64(),
						tx.Nonce(),
						tx.Data(),
						tx.To().Hex(),
						from.Hex(),
						*receipt};
	
	return;
}