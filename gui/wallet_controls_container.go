package gui;

import (
	fyne 			"fyne.io/fyne/v2"
	widget 			"fyne.io/fyne/v2/widget"
	container		"fyne.io/fyne/v2/container"
	accountops 		"wallet/accountops"
	dialog 			"fyne.io/fyne/v2/dialog"
	layout    		"fyne.io/fyne/v2/layout"
)

func (ui UserInterface) TabControlButtons() UserInterface{
	// Button to add a wallet
	addWalletButton_ := widget.NewButton("Add Wallet", func() { 
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
	// Button to add a new Wallet (generate wallet address)
	newWalletButton_ := widget.NewButton("New Wallet", func() { 
		ui = ui.CreateTab("New Wallet");
		ui = ui.DrawTabs();
		ui = ui.UpdateMainWindow();
	});
	// Button to rename a wallet
	renameWalletButton_ := widget.NewButton("Rename Wallet", func() { 
		dialog.ShowEntryDialog("Wallet Nickname", "", func(name string){
			ui = ui.RenameTab(name);
			ui = ui.DrawTabs();
			ui = ui.UpdateMainWindow();
		}, ui.MainWindow_);
	});
	// Button to delete a wallet
	deleteWalletButton_ := widget.NewButton("Delete Wallet", func() { 
		dialog.ShowConfirm("Deleting Wallet", "Are you sure you want to remove this wallet?", func(yes bool){
			if yes{
				ui = ui.DeleteTab();
				ui = ui.DrawTabs();
				ui = ui.UpdateMainWindow();
			}
		}, ui.MainWindow_);
	});
	// Button to save Wallet window state (to preserve info between runtimes)
	saveButton_ := widget.NewButton("Save state", func(){
		//DUMMY FUNCTION ATM
	});

	ui.TabControlContainer_ = container.NewHBox(
		addWalletButton_,
		newWalletButton_, 
		widget.NewSeparator(),
		renameWalletButton_,
		widget.NewSeparator(),
		deleteWalletButton_,
		layout.NewSpacer(),
		saveButton_,
	);
	return ui;
}

func (ui UserInterface) RenameTab(name string) (UserInterface){
	if (name == ""){
		name = "Empty name";
	}
	tabItem := ui.TabContainer_.CurrentTab()
	for tab := 0; tab < len(ui.Tabs_); tab++ {
		if tabItem == ui.Tabs_[tab].Tab_ {
			ui.Tabs_[tab].Name_ = name;
		}
	}
	return ui;
}

func (ui UserInterface) DeleteTab() (UserInterface){
	tabItem := ui.TabContainer_.CurrentTab()
	for tab := 0; tab < len(ui.Tabs_); tab++ {
		if tabItem == ui.Tabs_[tab].Tab_ {
			ui.Tabs_[tab] = ui.Tabs_[len(ui.Tabs_)-1]
			ui.Tabs_ = ui.Tabs_[:len(ui.Tabs_)-1]
		}
	}
	return ui;
}