package gui;

import (
	"fmt"
	big		 		"math/big"
	//errors 			"errors"
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

type TransactionInterface struct{
	AppPointer_   					*fyne.App;

	TransactionWindow_   			fyne.Window;
	TransactionCanvas_   			fyne.Canvas;

	ParentInterface_ 				WalletInterface;

	MasterContainer_ 				*fyne.Container;

	TransactionForm_     			*widget.Form;

	SendingAccountSelector_ 		*widget.Select;
	SendingAccountBallance_  		*widget.Label;

	DeliveryTypeRadio_  			*widget.RadioGroup;
	DeliveryAccountSelector_		*widget.Select;
	DeliveryAccountEntry_   		*widget.Entry;
	ValidateDeliveryAddressButton_	*widget.Button;

	EthereumToTransfer_   			*widget.Entry;
	GasPriceCheck_    				*widget.Check;
	GasPriceEntry_   				*widget.Entry;
	GasLimitCheck_    				*widget.Check;
	GasLimitEntry_   				*widget.Entry;

	MarketGasPriceText_    			*widget.Label;
	MarketGasPriceButton_   		*widget.Button;

	SendButton_ 					*widget.Button;
}

func (tw TransactionInterface) DisableGasPriceEntry(checked bool){
	if (checked){
		tw.GasPriceEntry_.Disable();
	} else {
		tw.GasPriceEntry_.Enable();
	}
}

func (tw TransactionInterface) DisableGasLimitEntry(checked bool){
	if (checked){
		tw.GasLimitEntry_.Disable();
	} else {
		tw.GasLimitEntry_.Enable();
	}
}

func (tw TransactionInterface) SendTransaction(
	ethereum_s,
	gasPrice_s,
	gasLimit_s,
	sendingAccountHex_s,
	deliveryAccountHex_s string,
){
	var a accountops.LocalAccount
	a = tw.ParentInterface_.TabMap_[sendingAccountHex_s];

	ethereum_bf, _, _	:= new(big.Float).Parse(ethereum_s, 10);
	gasLimit_int, _	 	:= strconv.Atoi(gasLimit_s);
	gasPrice_bf, _ 		:= new(big.Int).SetString(gasPrice_s, 10);

	err := transactions.Transfer(
		a, 
		url, 
		*ethereum_bf, 
		uint64(gasLimit_int), 
		*gasPrice_bf, 
		deliveryAccountHex_s,
	);
	if (err != nil){
		dialog.NewError(err, tw.TransactionWindow_).Show();
	} else {
		dialog.NewInformation("Successful transfer.", "", tw.TransactionWindow_);
	}
}

func (tw TransactionInterface) ValidateTransaction(){
	// Variable collections
	ethereumToTransfer_s := tw.EthereumToTransfer_.Text;
	sendingAccountHex_s := tw.SendingAccountSelector_.Selected;
	
	deliveryAccountHex_s := tw.DeliveryAccountEntry_.Text;
	if tw.DeliveryTypeRadio_.Selected != "Address"{
		deliveryAccountHex_s = tw.DeliveryAccountSelector_.Selected;
	}

	gasPrice_s := tw.GasPriceEntry_.Text;
	if tw.GasPriceCheck_.Checked == true{
		gasPrice_s = "0";
	}

	gasLimit_s := tw.GasLimitEntry_.Text;
	if tw.GasLimitCheck_.Checked == true{
		gasLimit_s = "0";
	} 

	fmt.Println(gasPrice_s);
	fmt.Println(gasLimit_s);

	selectedSending_localaccount := tw.ParentInterface_.TabMap_[sendingAccountHex_s];
	ballance_bf, _ := selectedSending_localaccount.GetBallance(url, nil); 
	ballance_s := ballance_bf.String();
	ballance_f, _ := strconv.ParseFloat(ballance_s, 64);
	if !accountops.IsAddressValid(sendingAccountHex_s) || !accountops.IsAddressValid(deliveryAccountHex_s){
		dialog.NewInformation("Invalid Address", 
			"Delivery or sending address is not a valid hexadecimal address.",
			tw.TransactionWindow_,
		).Show();
	} else if _, err := strconv.ParseFloat(gasPrice_s, 64); err != nil{
		dialog.NewInformation("Gas Price is invalid.",
			"The Gas price entered is not a valid numeric value",
			tw.TransactionWindow_,
		).Show();
	} else if _, err := strconv.ParseFloat(gasLimit_s, 64); err != nil{
		dialog.NewInformation("Gas limit is invalid.",
			"The Gas limit entered is not a valid numeric value",
			tw.TransactionWindow_,
		).Show();
	} else if _, err := strconv.ParseFloat(ethereumToTransfer_s, 64); err != nil{
		dialog.NewInformation("Transfer value is invalid.",
			"The transfer value entered is not a valid numeric value",
			tw.TransactionWindow_,
		).Show();
	} else if e, _ := strconv.ParseFloat(ethereumToTransfer_s, 64); ballance_f < e {
		dialog.NewInformation("Insufficient funds.",
			"Insufficient funds in transfer for selected account.",
			tw.TransactionWindow_,
		).Show();
	} else {
		tw.SendTransaction(
			ethereumToTransfer_s,
			gasPrice_s,
			gasLimit_s,
			sendingAccountHex_s,
			deliveryAccountHex_s,
		);
	}	
}

func (tw TransactionInterface) UpdateMarketGasPrice(){
	gasPrice := transactions.SuggestGasPrice(url);
	gasPriceEth := accountops.WeiToValue(*gasPrice);
	tw.MarketGasPriceText_.SetText(gasPrice.String() + " wei, \n" + gasPriceEth.String() + " ethereum .");
}

func (tw TransactionInterface) ValidateDeliveryAddress(){
	address := tw.DeliveryAccountEntry_.Text;
	fmt.Println(address);
	if accountops.IsAddressValid(address) == true{
		dialog.NewInformation(
			"Invalid Address", 
			"The address provided is not a hexadecimal address beginning in 0x.", 
			tw.TransactionWindow_,
		).Show();
	} else {
		dialog.NewInformation(
			"Valid Address", 
			"The address provided is a hexadecimal address.", 
			tw.TransactionWindow_,
		).Show();
	}
}

func (tw TransactionInterface) EnableDisableDeliveryTypeWidgets(selection string){
	fmt.Println(selection);
	if (selection == "Local wallet"){
		tw.DeliveryAccountEntry_.Disable();
		tw.DeliveryAccountSelector_.Enable();
		tw.ValidateDeliveryAddressButton_.Disable();
	} else {
		tw.DeliveryAccountEntry_.Enable();
		tw.DeliveryAccountSelector_.Disable();
		tw.ValidateDeliveryAddressButton_.Enable();
	}
}

func (tw TransactionInterface) GetBallanceOfWallet(address string){
	fmt.Println(address);
	var selectedAccount accountops.LocalAccount;
	selectedAccount = tw.ParentInterface_.TabMap_[address];
	ballance, _ := selectedAccount.GetBallance(url, nil); 
	fmt.Println(ballance.String());
	tw.SendingAccountBallance_.SetText(ballance.String() + " Ethereum");
}

func (tw TransactionInterface) GenerateTransactionInterface() (TransactionInterface){
	// Title style
	Monostyle := fyne.TextStyle{};
	Monostyle.Monospace = true;
	// Array of local wallets 
	localWallets := []string{};
	for num := 0; num < len(*tw.ParentInterface_.Tabs_); num ++{
		localWallets = append(localWallets, (*tw.ParentInterface_.Tabs_)[num].Account_.AddressHex_);
	}
	// Title
	title := canvas.NewText("Transaction: ", color.White);
	title.TextSize = 25;
	title.TextStyle = Monostyle;
	// Sending account ballance
	tw.SendingAccountBallance_ = widget.NewLabel("0 Ethereum");
	// Sending account selector
	tw.SendingAccountSelector_ = widget.NewSelect(localWallets, tw.GetBallanceOfWallet);
	// Delivery address selector (from local wallet) and delivery address entry
	tw.DeliveryAccountSelector_ = widget.NewSelect(localWallets, func(string){})
	tw.DeliveryAccountEntry_ = widget.NewEntry();
	tw.ValidateDeliveryAddressButton_ = widget.NewButton("Validate Address", tw.ValidateDeliveryAddress);
	tw.DeliveryTypeRadio_ = widget.NewRadioGroup([]string{"Local wallet", "Address"}, tw.EnableDisableDeliveryTypeWidgets);
	tw.DeliveryTypeRadio_.SetSelected("Address");
	tw.DeliveryTypeRadio_.Horizontal = true;
	// Ethereum to transfer
	tw.EthereumToTransfer_ = widget.NewEntry();
	// Gas price and gas limit
	tw.GasPriceEntry_ = widget.NewEntry();
	tw.GasPriceCheck_ = widget.NewCheck("Market determined", tw.DisableGasPriceEntry);
	tw.GasPriceCheck_.SetChecked(true);
	tw.GasLimitEntry_ = widget.NewEntry();
	tw.GasLimitCheck_ = widget.NewCheck("Default Gas limit", tw.DisableGasLimitEntry);
	tw.GasLimitCheck_.SetChecked(true);
	// Market gas price 
	tw.MarketGasPriceText_ = widget.NewLabel("");
	tw.MarketGasPriceButton_ = widget.NewButton("Refresh", tw.UpdateMarketGasPrice);

	// Send button
	tw.SendButton_ = widget.NewButton("Send", tw.ValidateTransaction);

	// Main form
	tw.TransactionForm_ = widget.NewForm(
		widget.NewFormItem("Sending Wallet:", 				tw.SendingAccountSelector_),
		widget.NewFormItem("Sending Wallet Ballance:", 		tw.SendingAccountBallance_),
		widget.NewFormItem("Delivery Wallet Type:", 		tw.DeliveryTypeRadio_),
		widget.NewFormItem("Delviery Account Selector:",	tw.DeliveryAccountSelector_),
		widget.NewFormItem("Delivery Account:", 			tw.DeliveryAccountEntry_),
		widget.NewFormItem("", 								tw.ValidateDeliveryAddressButton_),
		widget.NewFormItem("Ethereum to Transfer:", 		tw.EthereumToTransfer_),
		widget.NewFormItem("Gas price:", 					tw.GasPriceCheck_),
		widget.NewFormItem("Gas Price:", 					tw.GasPriceEntry_),
		widget.NewFormItem("Gas Limit:", 					tw.GasLimitCheck_),
		widget.NewFormItem("Gas Limit:", 					tw.GasLimitEntry_),
		widget.NewFormItem("Market Gas price: ", 			tw.MarketGasPriceText_),
		widget.NewFormItem("", 								tw.MarketGasPriceButton_),
	);

	tw.MasterContainer_ = container.NewCenter(
		container.NewVBox(
			title,
			tw.TransactionForm_,
			tw.SendButton_,
		),
	);
	return tw;
}

func NewTransactionWindow(ParentInterface WalletInterface) (TransactionInterface){
	tw := TransactionInterface{};
	tw.ParentInterface_ = ParentInterface;
	tw.AppPointer_ = ParentInterface.AppPointer_;
	tw.TransactionWindow_ = (*tw.AppPointer_).NewWindow("New Transaction");
	tw.TransactionCanvas_ = tw.TransactionWindow_.Canvas();
	tw.TransactionWindow_.Resize(fyne.NewSize(200,100));
	return tw;
}

func (tw TransactionInterface) ShowTransactionWindow() (TransactionInterface){
	tw.TransactionWindow_.SetContent(tw.MasterContainer_);
	tw.TransactionWindow_.Show();
	return tw;
}
