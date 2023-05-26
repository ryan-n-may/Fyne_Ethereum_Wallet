package gui;

import (
	fyne 			"fyne.io/fyne/v2"
	app 			"fyne.io/fyne/v2/app"
	theme  			"fyne.io/fyne/v2/theme"
	container		"fyne.io/fyne/v2/container"
	accountops 		"wallet/accountops"
)

const url = "https://cloudflare-eth.com";

type UserInterface struct{
	MyApp_							fyne.App;
	// Transaction window and canvas 
	TransactionWindow_   			fyne.Window;
	TransactionCanvas_   			fyne.Canvas;
	// Wallet window and canvas
	MainWindow_ 					fyne.Window;
	MainCanvas_						fyne.Canvas;
	// Wallet window containers
	TabMap_   						map[string]accountops.LocalAccount;
	Tabs_							[]TabContents;
	MasterContainer_   				*fyne.Container;
	TabContainer_					*container.AppTabs;
	TabControlContainer_			*fyne.Container;
	WalletWindowToolBarContainer_   *fyne.Container;
	// Transaction window containers
	TransactionWindowContainer_     *fyne.Container;
}

type TabContents struct{
	Tab_ 				*container.TabItem;
	TabContainer_		*fyne.Container;
	WalletContents_		*fyne.Container;
	Name_  				string;
	Account_ 			accountops.LocalAccount;
}

func InitialiseUI() (ui UserInterface){
	MyApp_ := app.New();
	MyApp_.Settings().SetTheme(theme.DarkTheme());
	MainWindow_ := MyApp_.NewWindow("Fyne Wallet");
	MainWindow_.Resize(fyne.NewSize(750, 500));
	MainCanvas_ := MainWindow_.Canvas();
	TransactionWindow_ := MyApp_.NewWindow("New Transaction");
	TransactionCanvas_ := TransactionWindow_.Canvas();
	TransactionWindow_.Resize(fyne.NewSize(200,100));
	TabContainer_ := container.NewAppTabs();
	TabMap_ := make(map[string]accountops.LocalAccount);
	ui = UserInterface{}
	// App
	ui.MyApp_ =	MyApp_;
	// Main window and canvas
	ui.MainWindow_ = MainWindow_;
	ui.MainCanvas_ = MainCanvas_;
	// Transaction window and canvas
	ui.TransactionWindow_ = TransactionWindow_;
	ui.TransactionCanvas_ = TransactionCanvas_;
	// Intialising the tab container
	ui.TabContainer_ = TabContainer_;
	ui.TabMap_ = TabMap_;
	return;
}

func (ui UserInterface) UpdateMainWindow() (UserInterface){
	ui = ui.DrawTabs();
	ui = ui.TabControlButtons();
	ui = ui.MainWindowControlBar();
	tabLayout := container.NewBorder(
		nil,
		nil,
		ui.WalletWindowToolBarContainer_,
		container.NewVBox(ui.TabControlContainer_, container.NewMax(ui.TabContainer_)),
	);
	ui.MainWindow_.SetContent(tabLayout);
	return ui;
}

func (ui UserInterface) Show() (UserInterface){
	ui.MainWindow_.ShowAndRun();
	return ui;
}