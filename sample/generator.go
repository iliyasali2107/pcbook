package sample

import "pcbook/pb"

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}
