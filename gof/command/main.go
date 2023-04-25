package main

func main() {
	tv := &tv{}

	onComm := &onCommand{device: tv}
	offComm := &offCommand{device: tv}

	onButton := button{command: onComm}
	offButton := button{command: offComm}

	onButton.press()

	offButton.press()
}
