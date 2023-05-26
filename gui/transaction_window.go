package gui;

import (
	"fmt"
	big		 		"math/big"
	errors 			"errors"
	strconv    		"strconv"
	widget 			"fyne.io/fyne/v2/widget"
	container		"fyne.io/fyne/v2/container"
	canvas 			"fyne.io/fyne/v2/canvas"
	fyne   			"fyne.io/fyne/v2"
	color   		"image/color"
	accountops 		"wallet/accountops"
	transactions    "wallet/transactions"
	dialog 			"fyne.io/fyne/v2/dialog"
)

func (ui UserInterface) GenerateTransactionContainer() (UserInterface){
	// Title 
	Monostyle := fyne.TextStyle{};
	Monostyle.Monospace = true;
	accountAddressLabel := canvas.NewText("Transaction: ", color.White);
	accountAddressLabel.TextSize = 25;
	accountAddressLabel.TextStyle = Monostyle;

	// Select account to send from
	fromSelectorText := widget.NewLabel("From Address:");
	options := []string{};
	for num := 0; num < len(ui.Tabs_); num ++{
		options = append(options, ui.Tabs_[num].Account_.AddressHex_);
	}
	fromBallance := widget.NewLabel("0 Ethereum");
	fromSelector := widget.NewSelect(options, func(address string){
		var selectedAccount accountops.LocalAccount;
		fmt.Println(address);
		selectedAccount = ui.TabMap_[address];
		ballance, _ := selectedAccount.GetBallance(url, nil); 
		fromBallance.SetText(ballance.String() + " Ethereum");	
	});
	fromSelectorContainer := container.NewGridWithRows(1,
		fromSelectorText, 
		fromSelector, 
		fromBallance,
	);
	
	// Type delivery address (hexadecimal)
	sendText := widget.NewLabel("Delivery Address:");
	sendField := widget.NewEntry();
	sendSelector := widget.NewSelect(options, func(address string){
		
	});
	sendContainer := container.NewGridWithRows(1,
		sendText, 
		sendField, 
		sendSelector,
	);

	// Transaction type selector
	transactionTypeRadio := widget.NewRadioGroup([]string{"Local wallet", "Address"},
		func(value string){
			if value == "Local wallet"{
				sendField.Disable();
				sendSelector.Enable();
			} else {
				sendSelector.Disable();
				sendField.Enable();
			}
		},
	);
	transactionTypeRadio.SetSelected("Address");
	transactionTypeLabel := widget.NewLabel("Transaction destination:");
	transactionTypeContainer := container.NewVBox(transactionTypeLabel, transactionTypeRadio);

	// Etherium to transfer
	ethValue := widget.NewLabel("Transfer value (Ethereum): ");
	ethField := widget.NewEntry();
	ethContainer := container.NewGridWithRows(1, 
		ethValue, 
		ethField,
	);

	// Gas price 
	gasPriceLabel := widget.NewLabel("Gas price:");
	gasPriceEntry := widget.NewEntry();
	gasPriceMarketToggle := widget.NewCheck("Market determined", func(value bool){
		if (value == true){
			gasPriceEntry.Disable();
		} else {
			gasPriceEntry.Enable();
		}
	});
	gasPriceMarketToggle.SetChecked(true);
	gasPriceContainer := container.NewGridWithRows(1,
		gasPriceLabel, 
		gasPriceMarketToggle, 
		gasPriceEntry,
	);

	// Gas limit 
	gasPriceLimitLabel := widget.NewLabel("Gas limit:");
	gasPriceLimitEntry := widget.NewEntry();
	gasPriceLimitMarketToggle := widget.NewCheck("Default limit", func(value bool){
		if (value == true){
			gasPriceLimitEntry.Disable();
		} else {
			gasPriceLimitEntry.Enable();
		}
	});
	gasPriceLimitMarketToggle.SetChecked(true);
	gasPriceLimitContainer := container.NewGridWithRows(1,
		gasPriceLimitLabel, 
		gasPriceLimitMarketToggle, 
		gasPriceLimitEntry,
	);


	// Send Button
	sendButton := widget.NewButton("Send", func(){
		fmt.Println("Inside send");

		if (fromSelector.SelectedIndex() == -1 || 
			(sendSelector.SelectedIndex() == -1 && transactionTypeRadio.Selected == "Local wallet")){
			fmt.Println("Inside error dialog");
			err := errors.New("Send or delivery account not selected.");
			dialog.NewError(err, ui.TransactionWindow_).Show();
			return;
		}
		// Collecting information
		fromAddress := options[fromSelector.SelectedIndex()];
		deliveryAddress := sendField.Text;
		if transactionTypeRadio.Selected == "Local wallet" {
			deliveryAddress = options[sendSelector.SelectedIndex()];
		}
		eth := ethField.Text;
		gasPrice := gasPriceEntry.Text;
		if gasPriceMarketToggle.Checked == true {
			gasPrice = "0";	
		}
		gasLimit := gasPriceLimitEntry.Text;
		if gasPriceLimitMarketToggle.Checked == true {
			gasLimit = "0";	
		}

		var selectedAccount accountops.LocalAccount;
		selectedAccount = ui.TabMap_[fromAddress];
		eth_f, ok := new(big.Float).SetString(eth);
		gasPrice_f, ok := new(big.Int).SetString(gasPrice, 10);
		gasLimit_int, err := strconv.Atoi(gasLimit);
		gasLimit_uint64 := uint64(gasLimit_int);
		// Checking for eth errors, gas price errors, or gas limit errors
		if (ok == false || err != nil){
			err := errors.New("Feilds are not filled correctly.");
			dialog.NewError(err, ui.TransactionWindow_).Show();
			return;
		}

		if (accountops.IsAddressValid(fromAddress) == false){
			err := errors.New("Sending Address is not valid.");
			dialog.NewError(err, ui.TransactionWindow_).Show();
			return;
		}

		if (accountops.IsAddressValid(deliveryAddress) == false){
			err := errors.New("Delivery Address is not valid.");
			dialog.NewError(err, ui.TransactionWindow_).Show();
			return;
		}


		dialog.NewConfirm("Confirmation", 
			"Transaction is to be set from \n" + fromAddress +
				" to " + deliveryAddress +
				"\n of value " + eth + " (Ethereum). " +
				" Gas price = " + gasPrice +
				" Gas limit = " + gasLimit +
				" ( 0 = Market determined)", func(yes bool){
					if(yes){
						fmt.Println("Sending transaction");
						err := transactions.Transfer(selectedAccount, url, *eth_f, gasLimit_uint64, *gasPrice_f, deliveryAddress);
						if (err != nil){
							dialog.NewError(err, ui.TransactionWindow_).Show();
						}
					}
				},
			ui.TransactionWindow_,
		).Show();
	});

	ui.TransactionWindowContainer_ = container.NewCenter(
		container.NewVBox(
			accountAddressLabel,
			fromSelectorContainer, 
			transactionTypeContainer,
			sendContainer,
			ethContainer,
			gasPriceContainer,
			gasPriceLimitContainer,
			sendButton,
		),
	);
	return ui;
}

func (ui UserInterface) InitialiseTransactionWindow() (UserInterface){
	ui.TransactionWindow_ = ui.MyApp_.NewWindow("New Transaction");
	ui.TransactionCanvas_ = ui.TransactionWindow_.Canvas();
	ui.TransactionWindow_.Resize(fyne.NewSize(200,100));
	return ui;
}

func (ui UserInterface) FillTransactionWindow() (UserInterface){
	tabLayout := container.NewVBox(
		ui.TransactionWindowContainer_,
	);
	ui.TransactionWindow_.SetContent(tabLayout);
	return ui;
}

func (ui UserInterface) ShowTransactionWindow() (UserInterface){
	ui.TransactionWindow_.Show();
	return ui;
}
