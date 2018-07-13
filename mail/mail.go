package mail

import (
	"gopkg.in/gomail.v2"
	"crypto/tls"
)

func SendMail(host string, port int, username string, password string, from string, to string, subject string, body string, emailType string) (bool, error) {
	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify:true}
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	//m.SetAddressHeader("Cc", "postmaster@mail.feehi.com", "Dan")
	m.SetBody(emailType, body)
	err := d.DialAndSend(m)
	if(err != nil){
		return false,err
	}
	return true,nil
}