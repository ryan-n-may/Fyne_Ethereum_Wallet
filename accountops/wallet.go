package accountops;

import (
	fmt				"fmt"
	common 			"github.com/ethereum/go-ethereum/common"
	hexutil 		"github.com/ethereum/go-ethereum/common/hexutil"
	crypto 			"github.com/ethereum/go-ethereum/crypto"
	secp256k1		"github.com/ethereum/go-ethereum/crypto/secp256k1"
	ecdsa 			"crypto/ecdsa"
	math 			"math"
	big		 		"math/big"
	context 		"context"
	ethclient 		"github.com/ethereum/go-ethereum/ethclient"
);

type LocalAccount struct{
	Address_ 			common.Address;
	AddressHex_ 		string;
	// Private key
	PrivateKey_ 		ecdsa.PrivateKey;
	PrivateKeyBytes_ 	[]byte;
	PrivateKeyHex_ 		string;
	// Public key
	PublicKey_			ecdsa.PublicKey;
	PublicKeyBytes_ 	[]byte;
	PublicKeyHex_ 		string;
}

func ValueToWei(in big.Float) big.Int{
	WEI 	:= math.Pow10(18);
    WEI_BIG := big.NewFloat(WEI);
	out, _	:= new(big.Float).Mul(&in, WEI_BIG).Int(nil);
	return *out;
}

func WeiToValue(in big.Int) (out big.Float){
	in_f 	:= big.NewFloat(float64(in.Uint64()));
	WEI 	:= math.Pow10(18);
    WEI_BIG := big.NewFloat(WEI);
	out 	= *new(big.Float).Quo(in_f, WEI_BIG);
	return;
}

func (a LocalAccount) PrintWalletInfo(url string, block big.Int){
	balance, pending := a.GetBallance(url, &block);
	balance_f, _ := balance.Float64();
	pending_f, _ := pending.Float64();
	fmt.Println("Address: \t", a.Address_);
	fmt.Println("Ballance: \t", balance_f);
	fmt.Println("Pending: \t", pending_f);
	fmt.Println("Private key: \t", a.PrivateKeyHex_);
	fmt.Println("Public key: \t", a.PublicKeyHex_);
}

func GetPublic(pri ecdsa.PrivateKey) (pub ecdsa.PublicKey){
	pri.PublicKey.Curve = secp256k1.S256();
	pri.PublicKey.X, pri.PublicKey.Y = pri.PublicKey.Curve.ScalarBaseMult(pri.D.Bytes());
	pub = pri.PublicKey;
	return;
}

func TrimLeft(hex string) string{
	var j int;
	for i := 0; i < len(hex); i++{
		fmt.Println(string(hex[i]));
		if (string(hex[i]) != "0"){
			j = i;
			i = len(hex);
		}
	}
	fmt.Println("j = ", j);
	return hex[j:]; 
}

func (a LocalAccount) SetPrivateAndPublicKey(private_key_hex string) LocalAccount{
	//private_key_hex = "0x" + TrimLeft(private_key_hex[2:]);
	// Private key
	fmt.Println("Setting private and public key");
	a.PrivateKeyHex_ 				= private_key_hex;
	fmt.Println("Setting D2, private_key_hex = " + private_key_hex);
	privateKey, err					:= crypto.HexToECDSA(private_key_hex[2:]);
	if (err != nil){
		panic(err);
	}
	a.PrivateKey_ 					= *privateKey;
	fmt.Println("Setting private key bytes");
	fmt.Println(a.PrivateKey_.D);
	a.PrivateKeyBytes_ 				= crypto.FromECDSA(&a.PrivateKey_);
	fmt.Println("Private keys are set");
	fmt.Println("Setting public keys");
	a.PrivateKey_.PublicKey 		= GetPublic(a.PrivateKey_);
	// Public key
	a.PublicKey_ 					= a.PrivateKey_.PublicKey;
	a.PublicKeyBytes_ 				= crypto.FromECDSAPub(&a.PublicKey_);
	a.PublicKeyHex_					= hexutil.Encode(a.PublicKeyBytes_);
	return a;
}

func Open(hex string) LocalAccount{
	var account LocalAccount;
	fmt.Println("Converting hex to address, and saving hex");
	account.Address_ = common.HexToAddress(hex);
	account.AddressHex_ = hex;
	return account;
}

func Create() LocalAccount{
	privateKey, err := crypto.GenerateKey();
	if (err != nil){
		panic(err);
	}

	privateKeyBytes := crypto.FromECDSA(privateKey);
	privateKeyHex 	:= hexutil.Encode(privateKeyBytes);

	publicKey 		:= privateKey.Public();
	publicKeyECDSA 	:= publicKey.(*ecdsa.PublicKey);
	publicKeyBytes 	:= crypto.FromECDSAPub(publicKeyECDSA);
	publicKeyHex 	:= hexutil.Encode(publicKeyBytes);

	address 	:= crypto.PubkeyToAddress(*publicKeyECDSA);
	addressHex  := address.Hex();

	var account LocalAccount;
	account.AddressHex_ 		= addressHex;
	account.Address_ 			= address;

	account.PrivateKey_ 		= *privateKey;
	account.PrivateKeyBytes_ 	= privateKeyBytes;
	account.PrivateKeyHex_ 		= privateKeyHex;

	account.PublicKey_ 			= *publicKeyECDSA;
	account.PublicKeyBytes_ 	= publicKeyBytes;
	account.PublicKeyHex_ 		= publicKeyHex;

	return account;
}

func (a LocalAccount) GetBallance(url string, block *big.Int) (etherium_ballance, etherium_pending *big.Float){
	client, err 	:= ethclient.Dial(url);
	if (err != nil){
		panic(err);
	}
	balance, err 	:= client.BalanceAt(context.Background(), a.Address_, block)
	if (err != nil){
		fmt.Println("balance: " + (*balance).String());
		fmt.Println("Address: " + a.Address_.String());
		fmt.Println("block: " 	+ (*block).String());
		panic(err);
	}
	floating_ballance := new(big.Float);
	floating_ballance.SetString(balance.String());
	etherium_ballance = new(big.Float).Quo(
							floating_ballance,
							big.NewFloat(math.Pow10(18)));
	
	pending_balance, err := client.PendingBalanceAt(context.Background(), a.Address_);
    if (err != nil){
		panic(err);
	}
	floating_pending := new(big.Float);
	floating_pending.SetString(pending_balance.String());
	etherium_pending = new(big.Float).Quo(
							floating_pending,
							big.NewFloat(math.Pow10(18)));
	return;
}