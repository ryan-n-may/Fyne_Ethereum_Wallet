package accountops;

import (
	ioutil 		"io/ioutil"
	os 			"os"
	keystore  	"github.com/ethereum/go-ethereum/accounts/keystore"
	accounts    "github.com/ethereum/go-ethereum/accounts"
);

func (a LocalAccount) CreateKeyStore(password string) accounts.Account{
	keystore_ := keystore.NewKeyStore("./tmp", 
					keystore.StandardScryptN, 
					keystore.StandardScryptP);
	account_keystore, err := keystore_.NewAccount(password);
	if (err != nil){
		panic(err);
	}
	return account_keystore;
}

func (a LocalAccount) ImportKeyStore(password, file string) accounts.Account{
	keystore_ := keystore.NewKeyStore("./tmp", 
					keystore.StandardScryptN, 
					keystore.StandardScryptP);
	jsonBytes, err := ioutil.ReadFile(file);
	if (err != nil){
		panic(err);
	}
	account_keystore, err := keystore_.Import(jsonBytes, password, password)
	if (err != nil){
		panic(err);
	}
	err = os.Remove(file);
	if (err != nil){
		panic(err);
	}
	return account_keystore;
}
