package accountops

import(
	regex 		"regexp"
	common 		"github.com/ethereum/go-ethereum/common"
	context 	"context"
	ethclient 	"github.com/ethereum/go-ethereum/ethclient"
);

func IsAddressValid(in string) bool{
	re := regex.MustCompile("^0x[0-9a-fA-F]{40}");
	valid := re.MatchString(in);
	return valid;
}

func IsSmartContract(url, hex string) bool{
	if IsAddressValid(hex) == false{
		return false;
	}
	address := common.HexToAddress(hex);
	client, err := ethclient.Dial(url);
	if (err != nil){
		panic(err);
	}
	bytecode, err := client.CodeAt(context.Background(),
						address,
						nil);
	if (err != nil){
		panic(err);
	}
	isContract := len(bytecode) > 0;
	return isContract;
}