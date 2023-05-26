package gui;

import (
	container		"fyne.io/fyne/v2/container"
	fyne			"fyne.io/fyne/v2"
	accountops 		"wallet/accountops"
	localaccount    "wallet/accountops"
)

func (ui UserInterface) CreateTab(name string) (UserInterface){
	account := accountops.Create();
	output := ui.GenerateWalletPage(account);
	tabContainer := fyne.NewContainer(output);
	tab := container.NewTabItem(name, tabContainer);
	newTab := TabContents{	tab,
							tabContainer,
							output,
							name,
							account };
	ui.TabMap_[account.AddressHex_] = account;
	ui.Tabs_ = append(ui.Tabs_, newTab);
	return ui;
}

func (ui UserInterface) CreateTabExistingWallet(name string, account accountops.LocalAccount) (UserInterface){
	output := ui.GenerateWalletPage(account);
	tabContainer := fyne.NewContainer(output);
	tab := container.NewTabItem(name, tabContainer);
	newTab := TabContents{	tab,
							tabContainer,
							output,
							name,
							account };
	ui.TabMap_[account.AddressHex_] = account;
	ui.Tabs_ = append(ui.Tabs_, newTab);
	return ui;
}

func (ui UserInterface) DrawTabs() (UserInterface){
	ui.TabContainer_ = container.NewAppTabs();
	for tab := 0; tab < len(ui.Tabs_); tab++ {
		ui.Tabs_[tab].Tab_ = container.NewTabItem(ui.Tabs_[tab].Name_, ui.Tabs_[tab].TabContainer_);
		ui.TabContainer_.Append(ui.Tabs_[tab].Tab_);
	}
	ui.TabContainer_.SetTabLocation(container.TabLocationTop);
	return ui;
}

func (ui UserInterface) GetCurrentAccount() (localaccount.LocalAccount){
	tabItem := ui.TabContainer_.CurrentTab();
	for tab := 0; tab < len(ui.Tabs_); tab++ {
		if tabItem == ui.Tabs_[tab].Tab_ {
			return ui.Tabs_[tab].Account_;
		}
	}
	return localaccount.LocalAccount{};
}

