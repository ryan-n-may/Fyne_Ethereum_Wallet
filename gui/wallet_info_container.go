package gui;

import (
	fmt 			"fmt"
	fyne 			"fyne.io/fyne/v2"
	widget 			"fyne.io/fyne/v2/widget"
	container		"fyne.io/fyne/v2/container"
	canvas    		"fyne.io/fyne/v2/canvas"
	theme 			"fyne.io/fyne/v2/theme";
	dialog 			"fyne.io/fyne/v2/dialog";
	color			"image/color"
	clipboard  		"golang.design/x/clipboard"
	accountops 		"wallet/accountops"
)


func (wi *WalletInterface) GenerateWalletPage(account accountops.LocalAccount) (*fyne.Container){
	err := clipboard.Init()
	if err != nil {
	    panic(err)
	}
	// Progress bar as we generate a new wallet
	progress := dialog.NewProgressInfinite("Generating New Wallet", "Calculating private and public keys.", wi.WalletWindow_);
	progress.Show();
	// Creating new wallet
	accountAddress := account.AddressHex_;
	privatekeyHex := account.PrivateKeyHex_;
	publickeyHex := account.PublicKeyHex_;
	fmt.Println("Getting ballance from generate wallet page");
	fmt.Println(url);
	ballance, pending := account.GetBallance(url, nil);
	// Hiding progress dialog once completed
	progress.Hide();
	// Setting wallet canvas text to monostyle
	Monostyle := fyne.TextStyle{};
	Monostyle.Monospace = true;
	//Address information and copy button
	accountAddressLabel := canvas.NewText("Address: " + accountAddress, color.White);
	accountAddressLabel.TextSize = 25;
	accountAddressLabel.TextStyle = Monostyle;
	accountAddressCopyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func(){
		clipboard.Write(clipboard.FmtText, []byte(accountAddress));
	})
	accountAddressBox := container.NewHBox(accountAddressCopyButton, accountAddressLabel);
	// Private key information, copy button, and show/hide toggle. 
	privateKeyHexLabel := canvas.NewText("Private key: ***", color.White);
	privateKeyHexLabel.TextSize = 15;
	privateKeyHexLabel.TextStyle = Monostyle;
	privateKeyHexShowButton := widget.NewCheck("show", func(value bool){
		if value{
			privateKeyHexLabel.Text = "Private key: " + privatekeyHex;
			dialog.ShowInformation("Your private key should not be shared publically.",
								 "Your private key is used to sign transactions and validate your wallet's identity.",
								 wi.WalletWindow_);
		} else {
			privateKeyHexLabel.Text = "Private key: **** ";
		}
		
	})
	privateKeyCopyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func(){
		clipboard.Write(clipboard.FmtText, []byte(privatekeyHex));
		dialog.ShowInformation("Your private key should not be shared publically.",
								 "Your private key is used to sign transactions and validate your wallet's identity.",
								 wi.WalletWindow_);
	})
	privateKeyHexBox := container.NewHBox(privateKeyCopyButton, privateKeyHexShowButton, privateKeyHexLabel);
	// Public key information, copy button, and show/hide toggle.
	publicKeyHexLabel := canvas.NewText("Public key: ***", color.White);
	publicKeyHexLabel.TextSize = 15;
	publicKeyHexLabel.TextStyle = Monostyle;
	publicKeyHexShowButton := widget.NewCheck("show", func(value bool){
		//clipboard.Write(clipboard.FmtText, []byte(accountAddress));
		if value{
			publicKeyHexLabel.Text = "Public key: " + publickeyHex[0:len(privatekeyHex)] + "...";
		} else {
			publicKeyHexLabel.Text = "Public key: **** ";
		}
		
	})
	publicKeyCopyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func(){
		clipboard.Write(clipboard.FmtText, []byte(publickeyHex));
	})
	publicKeyHexBox := container.NewHBox(publicKeyCopyButton, publicKeyHexShowButton, publicKeyHexLabel);
	// Etherium ballance text
	ballanceLabel := canvas.NewText(ballance.String() + " Ethereum",  color.White);
	ballanceLabel.TextSize = 30;
	ballanceLabel.TextStyle = Monostyle;
	// Pending Etherium ballance text
	pendingLabel := canvas.NewText(pending.String() + " Pending", color.White);
	pendingLabel.TextSize = 30;
	pendingLabel.TextStyle = Monostyle;
	// Container to stack etherium ballance and pending ballance
	ballanceStack := container.NewGridWithColumns(1, ballanceLabel, pendingLabel);
	ballanceBox := container.NewCenter(ballanceStack);
	// Main container for all wallet infromation output
	output := container.NewVBox(
					accountAddressBox,
					privateKeyHexBox,
					publicKeyHexBox,
					ballanceBox,
	);
	return output;
}