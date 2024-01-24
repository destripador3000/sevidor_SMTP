package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	host := "smtp.gmail.com"
	port := "587"
	from := "juanesvalon02@gmail.com"
	password := "ualwjimsbanzmcfk"

	toList := []string{"destripador.30001@gmail.com", "juanestebanvalencia722@gmail.com"}
	msg := "Subject:Correo de prueba\r\n" + "\r\n" + "¡¡¡Esto es una prueba!!!"
	body := []byte(msg)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Correo Enviado exitosamente")
}
