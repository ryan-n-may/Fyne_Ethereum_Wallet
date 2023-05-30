package saveops;

import (
	sha256   		"crypto/sha256"
	bytes   		"bytes"
	aes   			"crypto/aes"
	fmt 			"fmt"
	ioutil 			"io/ioutil"
	os 				"os"
	accountops 		"wallet/accountops"
	gob		   		"encoding/gob"
	common 			"github.com/ethereum/go-ethereum/common"
	ecdsa 			"crypto/ecdsa"
	crypto 			"github.com/ethereum/go-ethereum/crypto"
)

func Segment(in []byte, size int) (out [][]byte){
	for size < len(in) {
		in, out = in[size:], append(out, in[0:size:size]);
	}
	if len(in) == size {
		return append(out, in);
	} else {
		fmt.Println("Filling last slice");
		filler := make([]byte, size-len(in));
		in = append(in, filler...);
		return append(out, in);
	}
}

func DeSegment(in [][]byte) (out []byte){
	for i := 0; i < len(in); i++{
		out = append(out, in[i]...);
	}
	return out;
}

func SaveEncryptedStruct(user, pass string, in accountops.LocalAccount) error{
	key := sha256.Sum256([]byte(pass));	// Key should be 256 bits, so 32 bytes. 
	// Creating directory to save file
	wd, err := os.Getwd();
	dirPath := wd + "/accounts/" + user;
	filePath := dirPath + "/wallet_" + in.AddressHex_ + ".gob";
	err = os.MkdirAll(dirPath, os.ModePerm);
	if err != nil{
		return err;
	}
	structBytes, err := SerialiseStruct(in);
	if err != nil{
		return err;
	}

	structBytesSegmented := Segment(structBytes, 16);

	encryptedBytesSegmented := structBytesSegmented;

	for i := 0; i < len(structBytesSegmented); i++{
		encryptedBytesSegmented[i], err = EncryptBinary(key[:], structBytesSegmented[i]);
	}

	encryptedBytes := DeSegment(encryptedBytesSegmented[:][:]);
	if err != nil{
		return err;
	}
	err = WriteBytes(filePath, encryptedBytes);
	if err != nil{
		return err;
	}
	return nil;
}

func ReadEncryptedAccountDirectory(user, pass string) ([]accountops.LocalAccount, error){
	// Get key from password
	key := sha256.Sum256([]byte(pass));	// Key should be 256 bits, so 32 bytes. 
	// read all files in account directory
	wd, err := os.Getwd();
	dirPath := wd + "/accounts/" + user;
	if err != nil{
		return []accountops.LocalAccount{{}}, err;
	}
	files, err := ioutil.ReadDir(dirPath);
	if err != nil{
		return []accountops.LocalAccount{{}}, err;
	}
	var a_array [](accountops.LocalAccount);

	for _, f := range files {
		var accountDecode (accountops.LocalAccount);

		err, fileBytes := ReadBytes(dirPath + "/" + f.Name());
		if err != nil {
			return []accountops.LocalAccount{{}}, err;
		}

		segmentFileBytes := Segment(fileBytes, 16);
		decryptedBytesSegment := segmentFileBytes;

		for i := 0; i < len(segmentFileBytes); i++{
			decryptedBytesSegment[i], err = DecryptBinary(key[:], segmentFileBytes[i]);
			if err != nil{
				return []accountops.LocalAccount{{}}, err;
			}
		}
		
		decryptedBytes := DeSegment(decryptedBytesSegment);

		accountDecode, err = DeSerialiseStruct(decryptedBytes, key[:]);
		if err != nil{
			return []accountops.LocalAccount{{}}, err;
		}

		a_array = append(a_array, accountDecode);
	}
	return a_array, nil;
}

func EncryptBinary(key []byte, data []byte) ([]byte, error){
	c, err := aes.NewCipher(key);
	if err != nil{
		return nil, err;
	}

	out := make([]byte, len(data));
	c.Encrypt(out, data);
	return out, nil;
}

func DecryptBinary(key []byte, cipher []byte) ([]byte, error){
	c, err := aes.NewCipher(key);
	if err != nil{
		return nil, err;
	}

	out := make([]byte, len(cipher));
	c.Decrypt(out, cipher);
	return out, nil;
}

func GatherUsernames(rel_dir string) ([]string, error){
	wd, err := os.Getwd();
	if err != nil{
		return nil, err;
	}
	entries, err := os.ReadDir(wd + rel_dir);
	if err != nil{
		return nil, err;
	}
	var usernames []string;
	for _, e := range entries {
		usernames = append(usernames, e.Name());
	}
	return usernames, nil;
}

func DeSerialiseStruct(data []byte, key []byte) (accountops.LocalAccount, error){	
	gob.Register(accountops.LocalAccount{});
	gob.Register(new(common.Address));
	gob.Register(new(ecdsa.PublicKey));
	gob.Register(crypto.S256());

	var accountDecode accountops.LocalAccount;

	b := bytes.NewBuffer(data);

	decoder := gob.NewDecoder(b);
	err := decoder.Decode(&accountDecode);
	if err != nil{
		return accountops.LocalAccount{}, err;
	}
	
	return accountDecode, nil;
}

func SerialiseStruct(in accountops.LocalAccount) ([]byte, error){
	gob.Register(accountops.LocalAccount{});
	gob.Register(new(common.Address));
	gob.Register(new(ecdsa.PublicKey));
	gob.Register(crypto.S256());
	// Creating the correct directories and file names

	var b bytes.Buffer;

	// Encoding struct
	enc := gob.NewEncoder(&b);
	err := enc.Encode(in);
	if err != nil{
		return nil, err;
	}

	fmt.Println("encoded struct ", b.Bytes());

	return b.Bytes(), nil;
}

func ReadBytes(path string) (error, []byte){
	// Writing string to file
	dat, err := os.ReadFile(path);
	if err != nil{
		return err, nil;
	}
	return nil, dat;
}

func WriteBytes(filePath string, data []byte) error{
	// Creating file 
	f, err := os.Create(filePath);
	defer f.Close();
	if err != nil{
		return err;
	}
	// Writing string to file
	_, err = f.Write(data);
	if err != nil{
		return err;
	}

	return nil;
}