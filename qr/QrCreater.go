package qr

import (
	qrcode "github.com/skip2/go-qrcode"
	"fmt"
)

func CreateQr(publicDir string, ipAddr string) error {
	fileName := fmt.Sprintf("%s/connect.png", publicDir)
	return qrcode.WriteFile(ipAddr, qrcode.Medium, 256, fileName)
}
