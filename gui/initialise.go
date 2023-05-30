package gui;

import (
	fyne 			"fyne.io/fyne/v2"
	app 			"fyne.io/fyne/v2/app"
	theme  			"fyne.io/fyne/v2/theme"
)

const url = "https://cloudflare-eth.com";

type UserInterface struct{
	MyApp_							fyne.App;
	// Login Window
	LoginInterface_  				LoginInterface;
	// Transaction window and canvas 
	TransactionInterface_   		TransactionInterface;
	// Wallet window and canvas
	WalletInterface_  				WalletInterface;
}

func InitialiseUI() (ui UserInterface){
	ui = UserInterface{}

	// Intialsing the app
	ui.MyApp_ = app.New();
	ui.MyApp_.Settings().SetTheme(theme.DarkTheme());

	// Initialising the wallet window
	ui.WalletInterface_ 					= WalletInterface{};
	ui.WalletInterface_.WalletWindow_ 		= ui.MyApp_.NewWindow("Fyne Wallet");
	ui.WalletInterface_.WalletWindow_.Resize(fyne.NewSize(750, 500));
	ui.WalletInterface_.WalletCanvas_ 		= ui.WalletInterface_.WalletWindow_.Canvas();
	
	// Initialising the transaction window
	ui.TransactionInterface_						= TransactionInterface{};
	ui.TransactionInterface_.TransactionWindow_ 	= ui.MyApp_.NewWindow("New Transaction");
	ui.TransactionInterface_.TransactionWindow_.Resize(fyne.NewSize(200,100));
	ui.TransactionInterface_.TransactionCanvas_ 	= ui.TransactionInterface_.TransactionWindow_.Canvas();
	
	// Initialising the login widow
	ui.LoginInterface_ 				= LoginInterface{};
	ui.LoginInterface_.LoginWindow_ = ui.MyApp_.NewWindow("Login");
	ui.LoginInterface_.LoginWindow_.Resize(fyne.NewSize(200,100));

	return;
}