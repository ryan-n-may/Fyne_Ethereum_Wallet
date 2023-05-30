# Fyne Ethereum Wallet
![Ethereum](https://img.shields.io/badge/Ethereum-3C3C3D?style=for-the-badge&logo=Ethereum&logoColor=white) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Sublime Text](https://img.shields.io/badge/sublime_text-%23575757.svg?style=for-the-badge&logo=sublime-text&logoColor=important)

The Fyne Ethereum Wallet is a simple Ethereum wallet for storing and transacting Etherium. The wallet is build in Go using the Fyne GUI. 

This was built over a weekend while I had the flu. It is an exploration into cryptocurrency and a work in progress. Any contributions are welcome. The software still needs comprehensive documentation and commenting.

## Built using packages
* GO-Ethereum https://geth.ethereum.org/
* Fyne GUI https://fyne.io/

## Contents

## Accounts 
### Opening a new account
The Fyne Wallet stores wallet information between program sessions by serializing the wallet struct `accountops.LocalWallet` and saving it to binary. Wallet information files are organised in directories according to the `accountops.USERNAME` of the account currently accessing the program.  In order to create a new account, the user must fill the login window form as seen in figure 1. Saving the wallet state after creating a new account will cause a file structure like the one in figure 2 to be created.  
### Accessing an account
The Fyne Wallet allows accessing wallets attached to already created accounts. The login window allows selecting a subdirectory of the `/accounts/` folder to access (figure 3). The correct `accountops.PASSWORD` must be input to successfully access the account wallets. 
### Encryption 
All Fyne Wallet accounts are encrypted using AES-256 bit encryption built into Golang.  Encryption keys fixed-length and generated from `accountops.PASSWORD` via the SHA256 Hash function. Without the correct password, account wallet files cannot be decrypted and de-serialized (figure 4). 

Figure 1: New user login window.
Figure 2: Account file structure.
Figure 3: Unlock user login window
Figure 4: AES Decryption of locked user. 

## Wallet 
### The Wallet information container
* View wallet address, key-sets, and ballance.
### The Wallet control buttons
* Assign a wallet nickname.
* Delete, Add, and Generate wallets. 
* Add wallets via CSV data (address and private key).
* Save account state (serialize wallets into account directory)
### The Wallet toolbar
* Export wallets via CSV data.
* Import wallets via CSV data (address and private key).
* Open a new transaction window
* ~~View historical transactions~~ (Not implemented)
* ~~Create a new smart contract~~ (Not implemented)

Figure 5: Import/Exporting a wallet from CSV
Figure 6: Renaming a wallet
Figure 7: The Wallet Interface

## Transaction Menu
The transaction menu allows easy processing of transactions between local wallets, and to wallet addresses (figure 8). Further, gas price and gas limit can be assigned manually, or assigned to the market determined gas price. The present market gas price can be viewed and refreshed before completing the transactoin (figure 9). 

| Transaction Variable          | Type                        |
|-------------------------------|-----------------------------|
| Sending and Recieving address | `string` (Hex)              |
| Gas price                     | `big.Int` (Measured in WEI) |
| Gas limit                     | `uint64` (Measured in WEI)  |
| Ethereum value                | `big.Float` (Ethereum)      |

Figure 8: Sending funds to a local account/ to an address.
Figure 9: Transaction window.

