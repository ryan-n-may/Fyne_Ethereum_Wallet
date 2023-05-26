# Fyne Etherium Wallet
The Fyne Ethereum Wallet is a simple Ethereum wallet for storing and transacting Etherium. The wallet is build in Go using the Fyne GUI. 
## To Do List
* Simplify GUI structs and remove unnecessary UserInterface returns
![completion](https://progress-bar.dev/0)
* Implement smart contract capability
![completion](https://progress-bar.dev/0)
* Test Wallet GUI for error conditions
![completion](https://progress-bar.dev/50)
* Write test modules for Go-Etherium methods & Implement error handling
![completion](https://progress-bar.dev/15)
* Create a page to view historical transactions
![completion](https://progress-bar.dev/0)
* Write comprehensive in-code documentation
![completion](https://progress-bar.dev/10)
## Built using
* GO-Ethereum https://geth.ethereum.org/
* Fyne GUI https://fyne.io/
## The Wallet

![wallet GUI](https://github.com/ryan-n-may/Fyne_Etherium_Wallet/blob/main/screenshots/Screenshot%20from%202023-05-26%2022-14-26.png)
### Wallet operations
* View wallet key-sets and ballance.
* Assign a wallet nickname.
* Delete / Generate wallets. 
* Add wallets via CSV data (address and private key).
* Export wallets via CSV data.
## Transaction Menu
The transaction menu allows easy processing of transactions between local wallets (right), and to addresses (left). The transaction menu also has the option to assign gas price and gas limit manually. 

![transaction GUI](https://github.com/ryan-n-may/Fyne_Etherium_Wallet/blob/main/screenshots/Screenshot%20from%202023-05-26%2022-19-25.png)

