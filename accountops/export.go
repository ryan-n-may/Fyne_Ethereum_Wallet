package accountops;

import (
	"fmt"
	os 	  			"os"
	errors  		"errors"
	strings   		"strings"

);

func LoadFromCSV(path string) (LocalAccount, error){
    data, err := os.ReadFile(path);
    if err != nil {
        return LocalAccount{}, err;
    }
    wallet_info := strings.Split(string(data), ",");
    if (len(wallet_info) == 2 && IsAddressValid(wallet_info[0])){
    	fmt.Println("Address: " + wallet_info[0]);
    	fmt.Println("Key: " + wallet_info[1]);
    	a := Open(string(wallet_info[0]));
    	a = a.SetPrivateAndPublicKey(string(wallet_info[1]));
    	return a, nil;
	} else {
		err := errors.New("File is not a valid account.");
		return LocalAccount{}, err;
	}
}

func ExportToCSV(localaccount LocalAccount, path string) error{
	export := localaccount.AddressHex_ + "," + localaccount.PrivateKeyHex_;
	f, err := os.Create(path)
    if err != nil {
        return err;
    }
    defer f.Close()
    _, err = f.WriteString(export)
    if err != nil {
        return err;
    }
    return nil;
}