package main

import (
	"log"

	p1XX "github.com/Tech-Arch1tect/p1XX-go"
)

func main() {
	d := p1XX.New("xxx.xxx.xxx.xxx", "email/username", "password")
	if err := d.Handshake(); err != nil {
		log.Panic(err)
	}
	if err := d.Login(); err != nil {
		log.Panic(err)
	}
	deviceInfo, err := d.GetDeviceInfo()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Device info: %+v", deviceInfo)
}
