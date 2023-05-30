package main;

import (
	gui				"wallet/gui"
);

func main(){
	ui := gui.InitialiseUI();
	ui, li := ui.InitialiseLoginInterface();
	li.StartLoginWindow();
}	