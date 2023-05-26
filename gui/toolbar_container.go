package gui;

import (
	fyne            "fyne.io/fyne/v2"
	widget 			"fyne.io/fyne/v2/widget"
	dialog    		"fyne.io/fyne/v2/dialog"
	container		"fyne.io/fyne/v2/container"
	accountops	    "wallet/accountops"
)

func (ui UserInterface) MainWindowControlBar() UserInterface{
	// Button to launch a new transaction 
	newTransactionButton := widget.NewButton("New transaction.", func(){
		ui = ui.InitialiseTransactionWindow();
		ui = ui.GenerateTransactionContainer();
		ui = ui.FillTransactionWindow();
		ui = ui.ShowTransactionWindow();
	});
	// Button to start a new smart contract
	newContractButton := widget.NewButton("New smart contract.", func(){
		
	});
	// Button to view past transactions
	viewTransactionsButton := widget.NewButton("View transactions.", func(){
		
	});
	// Button to load wallets from CSV files
	loadWalletInfoButton := widget.NewButton("Load wallet from csv.", func() { 
		dialog.ShowFileOpen(
		 func(reader fyne.URIReadCloser, err error){
		 	if (reader != nil){
		 		path := reader.URI().Path();
				name := reader.URI().Name();
				localaccount, err := accountops.LoadFromCSV(path);			
				if (err == nil){
					ui = ui.CreateTabExistingWallet("Wallet", localaccount);
					ui = ui.DrawTabs();
					ui = ui.UpdateMainWindow();
					dialog.ShowInformation("Success.", "Wallet successfully loaded from " + name, ui.MainWindow_);	
				} else {
					dialog.ShowError(err, ui.MainWindow_);
				}
			 	
		 	}
		 }, ui.MainWindow_);
	});
	// Button to export wallet information to CSV file
	exportWalletInfoButton := widget.NewButton("Export wallet information.", func(){
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error){
			if (writer != nil && err == nil){
				path := writer.URI().Path();
				name := writer.URI().Name();
				err = accountops.ExportToCSV(ui.GetCurrentAccount(), path);
				if (err != nil){
					dialog.ShowError(err, ui.MainWindow_);
				} else {
					dialog.ShowInformation("Success", "Wallet informatoin has been saved to CSV file: " + name, ui.MainWindow_);
				}
			}
		}, ui.MainWindow_);
		
	});
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
	ui.WalletWindowToolBarContainer_ = container.NewVBox(
		transactionOptions_,
		widget.NewSeparator(),
		walletOptions_,
		widget.NewSeparator(),
	);
	return ui;
}