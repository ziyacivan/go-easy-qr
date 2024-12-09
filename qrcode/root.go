package qrcode

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(url string) ([]byte, error) {
	qrCode, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}

	return qrCode, err
}
