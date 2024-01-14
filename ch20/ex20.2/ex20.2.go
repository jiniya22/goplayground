package main

import (
	"github.com/jiniya22/goplayground/ch20/fedex"
	"github.com/jiniya22/goplayground/ch20/koreapost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}
func main() {
	fedexSender := &fedex.FedexSender{}
	postSender := &koreapost.PostSender{}
	SendBook("ABC 신나는 영어", fedexSender)
	SendBook("너에게 보내는 편지", postSender)
}
