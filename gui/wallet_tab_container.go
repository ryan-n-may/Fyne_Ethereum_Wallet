package gui;

import (
	fmt 			"fmt"
	container		"fyne.io/fyne/v2/container"
	fyne			"fyne.io/fyne/v2"
	accountops 		"wallet/accountops"
	localaccount    "wallet/accountops"
)

func (wi *WalletInterface) CreateTab(name string){
	account := accountops.Create();
	output := wi.GenerateWalletPage(account);
	tabContainer := fyne.NewContainer(output);
	tab := container.NewTabItem(name, tabContainer);
	newTab := TabContents{	tab,
							tabContainer,
							output,
							name,
							account };
	fmt.Println(account.AddressHex_);
	wi.TabMap_[account.AddressHex_] = account;
	*wi.Tabs_ = append(*wi.Tabs_, newTab);
}

func (wi *WalletInterface) CreateTabExistingWallet(name string, account accountops.LocalAccount){
	output := wi.GenerateWalletPage(account);
	tabContainer := fyne.NewContainer(output);
	tab := container.NewTabItem(name, tabContainer);
	newTab := TabContents{	tab,
							tabContainer,
							output,
							name,
							account };
	wi.TabMap_[account.AddressHex_] = account;
	*wi.Tabs_ = append(*wi.Tabs_, newTab);
}

func (wi *WalletInterface) DrawTabs(){
	wi.TabContainer_ = container.NewAppTabs();
	for tab := 0; tab < len(*wi.Tabs_); tab++ {
		fmt.Println((*wi.Tabs_)[tab].Name_);
		(*wi.Tabs_)[tab].Tab_ = container.NewTabItem((*wi.Tabs_)[tab].Name_, (*wi.Tabs_)[tab].TabContainer_);
		wi.TabContainer_.Append((*wi.Tabs_)[tab].Tab_);
	}
	fmt.Print("After appening tabs: ");
	fmt.Println(len((*wi.TabContainer_).Items));
	wi.TabContainer_.SetTabLocation(container.TabLocationTop);
}

func (wi *WalletInterface) GetCurrentAccount() (localaccount.LocalAccount){
	tabItem := wi.TabContainer_.CurrentTab();
	for tab := 0; tab < len(*wi.Tabs_); tab++ {
		if tabItem == (*wi.Tabs_)[tab].Tab_ {
			return (*wi.Tabs_)[tab].Account_;
		}
	}
	return localaccount.LocalAccount{};
}

