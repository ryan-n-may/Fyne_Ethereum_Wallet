package main;

import (
	//fmt				"fmt"
	//big		 		"math/big"
	//wallet 			"wallet/accountops"
	//transactions 	"wallet/transactions"
	gui				"wallet/gui"
);

func main(){
	ui := gui.InitialiseUI();
	ui = ui.UpdateMainWindow();
	ui = ui.Show();

	/*
	fmt.Println("My Exodus wallet:");

	hex_address := "0x6e3369F1B809cfaED18bB0079B0179685D540a46";
	private_key_hex := "3475a9e0b22294648aecb8afbbdd4c58687cdd9ed0ca896c05225b52b121e8d1";

	url := "https://cloudflare-eth.com";
	
	var block big.Int;
	block.SetUint64(5532993);


	a := wallet.Open(hex_address);
	a = a.SetPrivateAndPublicKey(private_key_hex);

	a.PrintWalletInfo(url, block);

	

	b := wallet.Create();
	fmt.Println("Creating another wallet:");
	b.PrintWalletInfo(url, block);

	/*
	//coinspot_recieve_address 	:= "0x5e6d52dcaf7128dd6afe564fad68bc842400c787";
	eth_value 					:= *big.NewFloat(0.0001);
	successful := transactions.Transfer(a, url, eth_value, 0, *big.NewInt(0), b.AddressHex_);

	if (successful == true){
		fmt.Print("Successful transfer");
	}else{
		fmt.Print("Unsuccessful transfer");
	}

	*/

	//account := a.CreateKeyStore("password");
	//fmt.Println(account.Address.Hex());*/
	

	/* in Eth, Gas price is the free required to successfully conduct
	a transaction or contract on blockchain. Based on computation power, and 
	based upon gwei (10^-9 ETH). 
	Gas price is normally determined by the market. 
	*/

	/* in Eth, Nonce is the number of transactions sent from a given address.
	in cryptography, a nonce is a one-time code selected in a random manner.
	Each time a transaction is sent, nonce is increased by 1. 
	Transactions must be in order, and no skipping is allowed. 
	Nonce prevents double spending. */


}	