package main

/*
21.
Реализовать паттерн «адаптер» на любом примере.
*/

func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)

	windowsMachine := &windows{}

	windowsMachineAdapter := windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertSquareUsbInComputer(&windowsMachineAdapter)
}
