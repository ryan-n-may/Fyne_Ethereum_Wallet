package gui;

import (
	big  			"math/big"
	fyne 			"fyne.io/fyne/v2"
	container		"fyne.io/fyne/v2/container"
	dialog 			"fyne.io/fyne/v2/dialog"
	accountops 		"wallet/accountops"
	saveops 		"wallet/saveops"
)


type WalletInterface struct {
	AppPointer_   					*fyne.App;

	WalletWindow_ 					fyne.Window;
	WalletCanvas_					fyne.Canvas;
	// Wallet window containers
	TabMap_   						map[string]accountops.LocalAccount;
	Tabs_							*[]TabContents;
	MasterContainer_   				*fyne.Container;
	TabContainer_					*container.AppTabs;
	TabControlContainer_			*fyne.Container;
	WalletWindowToolBarContainer_   *fyne.Container;
	MasterContianer_ 				*fyne.Container;
}

type TabContents struct{
	Tab_ 				*container.TabItem;
	TabContainer_		*fyne.Container;
	WalletContents_		*fyne.Container;
	Name_  				string;
	Account_ 			accountops.LocalAccount;
}

func (ui* UserInterface) InitialiseWalletInterface() (WalletInterface){
	ui.WalletInterface_ 					= WalletInterface{};
	ui.WalletInterface_.WalletWindow_ 		= ui.MyApp_.NewWindow("Fyne Wallet: " + accountops.USERNAME);
	ui.WalletInterface_.WalletWindow_.Resize(fyne.NewSize(750, 500));
	ui.WalletInterface_.WalletCanvas_ 		= ui.WalletInterface_.WalletWindow_.Canvas();
	ui.WalletInterface_.AppPointer_ 		= &ui.MyApp_;
	ui.WalletInterface_.TabContainer_ 		= container.NewAppTabs();
	ui.WalletInterface_.TabMap_ 			= make(map[string]accountops.LocalAccount);
	ui.WalletInterface_.Tabs_ 				= new([]TabContents);
	return ui.WalletInterface_;
}

func (wi *WalletInterface) PreLoadAccount(){	
	account_array, err := saveops.ReadEncryptedAccountDirectory(accountops.USERNAME, accountops.PASSWORD);
	if err != nil{
		dialog.NewError(err, wi.WalletWindow_);
	}
	for _, account := range account_array{
		account.PrintWalletInfo(url, *big.NewInt(0));
		wi.CreateTabExistingWallet(account.Nickname_, account);
	}
}

func (wi *WalletInterface) UpdateWalletInterface(){
	wi.DrawTabs();
	wi.TabControlButtons();
	wi.MainWindowControlBar();
	wi.MasterContainer_ = container.NewBorder(
		nil,
		nil,
		wi.WalletWindowToolBarContainer_,
		container.NewVBox(wi.TabControlContainer_, container.NewMax(wi.TabContainer_)),
	);
	wi.WalletWindow_.SetContent(wi.MasterContainer_);
}

func (wi* WalletInterface) ShowWalletInterface(){
	wi.WalletWindow_.Show();
}