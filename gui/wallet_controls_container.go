package gui;

import (
	fmt 			"fmt"
	fyne 			"fyne.io/fyne/v2"
	widget 			"fyne.io/fyne/v2/widget"
	container		"fyne.io/fyne/v2/container"
	accountops 		"wallet/accountops"
	dialog 			"fyne.io/fyne/v2/dialog"
	layout    		"fyne.io/fyne/v2/layout"
	saveops   		"wallet/saveops";
)

func (wi *WalletInterface) AddWalletFunction(){
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
					dialog.ShowInformation(
						"Success.", 
						"Wallet successfully loaded from " + name, 
						wi.WalletWindow_,
					);	
				} else {
					dialog.ShowError(err, wi.WalletWindow_);
				}
			 	
		 	}
		 }, wi.WalletWindow_);
}

func (wi *WalletInterface) NewWalletFunction(){
	fmt.Println("Creating tab");
	wi.CreateTab("New Wallet");
	fmt.Println("Drawing tabs");
	wi.DrawTabs();
	fmt.Println("Updating wallet interface");
	wi.UpdateWalletInterface();
	fmt.Print("Tab container tabs: ");
	fmt.Println(len(wi.TabContainer_.Items));
}

func (wi *WalletInterface) RenameWalletFunction(){
	dialog.ShowEntryDialog("Changing Wallet Nickname", "Nickname:", 
		func(name string){
			wi.RenameTab(name);
			wi.DrawTabs();
			wi.UpdateWalletInterface();
		}, wi.WalletWindow_);
}

func (wi *WalletInterface) DeleteWalletFunction(){
	dialog.ShowConfirm("Deleting Wallet", "Are you sure you want to remove this wallet?", func(yes bool){
			if yes{
				wi.DeleteTab();
				wi.DrawTabs();
				wi.UpdateWalletInterface();
			}
		}, wi.WalletWindow_);
}

func (wi *WalletInterface) SaveStateFunction(){
	for i := 0; i < len(*wi.Tabs_); i++{
		fmt.Println("saving wallet");
		err := saveops.SaveEncryptedStruct(accountops.USERNAME, accountops.PASSWORD, (*wi.Tabs_)[i].Account_);
		if err != nil{
			dialog.NewError(err, wi.WalletWindow_).Show();
		}
	}	
}

func (wi *WalletInterface) TabControlButtons() {
	// Button to add a wallet
	addWalletButton_ := widget.NewButton("Add Wallet", wi.AddWalletFunction);
	// Button to add a new Wallet (generate wallet address)
	newWalletButton_ := widget.NewButton("New Wallet", wi.NewWalletFunction);
	// Button to rename a wallet
	renameWalletButton_ := widget.NewButton("Rename Wallet", wi.RenameWalletFunction);
	// Button to delete a wallet
	deleteWalletButton_ := widget.NewButton("Delete Wallet", wi.DeleteWalletFunction);
	// Button to save Wallet window state (to preserve info between runtimes)
	saveButton_ := widget.NewButton("Save state", wi.SaveStateFunction);

	wi.TabControlContainer_ = container.NewHBox(
		addWalletButton_,
		newWalletButton_, 
		widget.NewSeparator(),
		renameWalletButton_,
		widget.NewSeparator(),
		deleteWalletButton_,
		layout.NewSpacer(),
		saveButton_,
	);
}

func (wi *WalletInterface) RenameTab(name string){
	if (name == ""){
		name = "Empty name";
	}
	tabItem := wi.TabContainer_.CurrentTab()
	for tab := 0; tab < len(*wi.Tabs_); tab++ {
		if tabItem == (*wi.Tabs_)[tab].Tab_ {
			(*wi.Tabs_)[tab].Name_ = name;
			(*wi.Tabs_)[tab].Account_.Nickname_ = name;
		}
	}
}

func (wi *WalletInterface) DeleteTab(){
	tabItem := wi.TabContainer_.CurrentTab()
	for tab := 0; tab < len(*wi.Tabs_); tab++ {
		if tabItem == (*wi.Tabs_)[tab].Tab_ {
			(*wi.Tabs_)[tab] = (*wi.Tabs_)[len(*wi.Tabs_)-1]
			(*wi.Tabs_) = (*wi.Tabs_)[:len(*wi.Tabs_)-1]
		}
	}
}