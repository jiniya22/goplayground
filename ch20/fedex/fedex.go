package fedex

import "fmt"

type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %s parcel\n", parcel)
}

// github.com/jiniya22/goplayground/ch20/fedex
