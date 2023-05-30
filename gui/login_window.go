package gui;

import (
	widget   		"fyne.io/fyne/v2/widget"
	dialog   		"fyne.io/fyne/v2/dialog"
	fyne 			"fyne.io/fyne/v2"
	fmt 			"fmt"
	saveops  		"wallet/saveops"

	accountops 		"wallet/accountops"
)

type LoginInterface struct{
	AppPointer_   			*fyne.App;
	UserInterfacePointer	*UserInterface;

	LoginWindow_ 			fyne.Window;

	LoginForm_   			*widget.Form;

	AccountSelector_  		*widget.Select;
	PasswordEntry_  		*widget.Entry;
	ShowPasswordCheck_  	*widget.Check;
	UnlockButton_   		*widget.Button;

	UsernameEntry_     		*widget.Entry;
	NewPasswordEntry_   	*widget.Entry;
	NewAccountButton_   	*widget.Button;
}

func (lw LoginInterface) UnlockAccount(){
	accountops.USERNAME = lw.AccountSelector_.Selected;
	accountops.PASSWORD = lw.PasswordEntry_.Text;
	wi := (*lw.UserInterfacePointer).InitialiseWalletInterface();
	progress := dialog.NewProgressInfinite("Opening account", "Decrypting via AES-256", lw.LoginWindow_);
	progress.Show();
	wi.PreLoadAccount();
	progress.Hide();
	wi.UpdateWalletInterface();
	wi.ShowWalletInterface();
}

func (lw LoginInterface) OpenNewAccount(){
	accountops.USERNAME = lw.UsernameEntry_.Text;
	fmt.Println(lw.UsernameEntry_.Text);
	accountops.PASSWORD = lw.NewPasswordEntry_.Text;
	fmt.Println("initialising wallet interface");
	wi := (*lw.UserInterfacePointer).InitialiseWalletInterface();
	fmt.Println("Updating wallet interface");
	wi.UpdateWalletInterface();
	fmt.Println("Showing wallet interface");
	wi.ShowWalletInterface();
}

func (ui UserInterface) InitialiseLoginInterface() (UserInterface, LoginInterface){
	ui.LoginInterface_ 					= LoginInterface{};
	ui.LoginInterface_.UserInterfacePointer = &ui;
	ui.LoginInterface_.LoginWindow_		= ui.MyApp_.NewWindow("Fyne Wallet");
	ui.LoginInterface_.LoginWindow_.Resize(fyne.NewSize(100, 200));
	ui.LoginInterface_.AppPointer_ 		= &ui.MyApp_;
	return ui, ui.LoginInterface_;
}

func (lw* LoginInterface) StartLoginWindow(){
	options, err := saveops.GatherUsernames("/accounts/");
	if err != nil{
		dialog.NewError(err, lw.LoginWindow_);
	}

	lw.AccountSelector_ = widget.NewSelect(options, func(string){});
	lw.PasswordEntry_ = widget.NewEntry();
	lw.PasswordEntry_.Password = true;
	lw.UnlockButton_ = widget.NewButton("Unlock", lw.UnlockAccount);

	lw.UsernameEntry_ = widget.NewEntry();
	lw.NewPasswordEntry_ = widget.NewEntry();
	lw.NewAccountButton_ = widget.NewButton("New Account", lw.OpenNewAccount);

	lw.LoginForm_ = widget.NewForm(
		widget.NewFormItem("Account:", lw.AccountSelector_),
		widget.NewFormItem("Password:", lw.PasswordEntry_),
		widget.NewFormItem("", lw.UnlockButton_),

		widget.NewFormItem("New Username:", lw.UsernameEntry_),
		widget.NewFormItem("New Password:", lw.NewPasswordEntry_),
		widget.NewFormItem("", lw.NewAccountButton_),
	);

	lw.LoginWindow_.SetContent(lw.LoginForm_);
	lw.LoginWindow_.ShowAndRun();
}