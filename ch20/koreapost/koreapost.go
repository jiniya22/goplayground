package koreapost

import "fmt"

type PostSender struct {
}

func (f *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %s 를 보냅니다\n", parcel)
}
