package gui;

import (
	fyne            "fyne.io/fyne/v2"
	widget 			"fyne.io/fyne/v2/widget"
	dialog    		"fyne.io/fyne/v2/dialog"
	container		"fyne.io/fyne/v2/container"
	accountops	    "wallet/accountops"
)

func (wi WalletInterface) InitialiseTransactionWindowFunction(){
	tw := NewTransactionWindow(wi);
	tw = tw.GenerateTransactionInterface();
	tw = tw.ShowTransactionWindow();
}

func (wi WalletInterface) LoadWalletFromCSVFunction(){
	dialog.ShowFileOpen(
		func(reader fyne.URIReadCloser, err error){
		 	if (reader != nil){
		 		path := reader.URI().Path();
				name := reader.URI().Name();
				localaccount, err := accountops.LoadFromCSV(path);			
				if (err == nil){
					wi.CreateTabExistingWallet("Wallet", localaccount);
					wi.DrawTabs();
					wi.UpdateWalletInterface();
					dialog.ShowInformation("Success.", "Wallet successfully loaded from " + name, wi.WalletWindow_);	
				} else {
					dialog.ShowError(err, wi.WalletWindow_);
				}
			 	
		 	}
		}, wi.WalletWindow_,
	);
}

func (wi WalletInterface) SaveWalletToCSVFunction(){
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error){
			if (writer != nil && err == nil){
				path := writer.URI().Path();
				name := writer.URI().Name();
				err = accountops.ExportToCSV(wi.GetCurrentAccount(), path);
				if (err != nil){
					dialog.ShowError(err, wi.WalletWindow_);
				} else {
					dialog.ShowInformation("Success", "Wallet informatoin has been saved to CSV file: " + name, wi.WalletWindow_);
				}
			}
		}, wi.WalletWindow_,
	);
}

func (wi *WalletInterface) MainWindowControlBar(){
	// Button to launch a new transaction 
	newTransactionButton := widget.NewButton("New transaction.", wi.InitialiseTransactionWindowFunction);
	
	// Button to start a new smart contract
	newContractButton := widget.NewButton("New smart contract.", func(){ /* DUMMY */ });
	// Button to view past transactions
	viewTransactionsButton := widget.NewButton("View transactions.", func(){ /* DUMMY */ });
	// Button to load wallets from CSV files
	loadWalletInfoButton := widget.NewButton("Load wallet from csv.", wi.LoadWalletFromCSVFunction);
	// Button to export wallet information to CSV file
	exportWalletInfoButton := widget.NewButton("Export wallet information.", wi.SaveWalletToCSVFunction);
	// Accordian items for the transaction options
	transactionsAccorianItem := widget.NewAccordionItem(
		"Transactions", 
		container.NewVBox(
			newTransactionButton,
			newContractButton,
		    viewTransactionsButton,
		),
	);
	// Transaction accordian
	transactionOptions_ := widget.NewAccordion(
		transactionsAccorianItem,
	);
	// Accordian items for the wallet saving options
	walletOpsAccordianItem := widget.NewAccordionItem(
		"Wallet Operations", 
		container.NewVBox(
			loadWalletInfoButton,
			exportWalletInfoButton,
		),
	);
	// Wallet accordian
	walletOptions_ := widget.NewAccordion(
		walletOpsAccordianItem,
	);
	// Master container for wallet toolbar
	wi.WalletWindowToolBarContainer_ = container.NewVBox(
		transactionOptions_,
		widget.NewSeparator(),
		walletOptions_,
		widget.NewSeparator(),
	);
}