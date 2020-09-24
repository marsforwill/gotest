package main

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/smtp"
)

type LoginAuth struct {
	username, password string
}

func NewLoginAuth(username, password string) smtp.Auth {
	return &LoginAuth{username, password}
}

func (a *LoginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}

func SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	host, _, _ := net.SplitHostPort(addr)
	if err != nil {
		fmt.Println("call dial")
		return err
	}
	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: host, InsecureSkipVerify: true}
		if err = c.StartTLS(config); err != nil {
			fmt.Println("call start tls")
			return err
		}
	}

	if a != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(a); err != nil {
				fmt.Println("check auth with err:", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}

	header := make(map[string]string)
	header["Subject"] = "Golang发送邮件测试"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString(msg)
	_, err = w.Write([]byte(message))

	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func main() {
	// Set up authentication information.
	//auth := smtp.PlainAuth("", "liushimin@sensetime.com", "Mars940118", "smtp.partner.outlook.cn")
	//// Connect to the server, authenticate, set the sender and recipient,
	////and send the email all in one step.
	//to := []string{"liushimin@sensetime.com"}
	//msg := []byte("To: liushimin@sensetime.com\r\n" +     "Subject: discount Gophers!\r\n" +     "\r\n" +     "This is the email body.\r\n")
	//err := smtp.SendMail("smtp.partner.outlook.cn:587", auth, "liushimin@sensetime.com", to, msg)
	//if err != nil {
	//	log.Fatal(err)
	//}

	auth := NewLoginAuth("liushimin@sensetime.com", "Mars940118")

	to := []string{"zhangmingyang@sensetime.com"}
	msg := []byte("这是一封来自go的测试邮件")
	err := SendMail("smtp.partner.outlook.cn:587", auth, "liushimin@sensetime.com", to, msg)
	if err != nil {
		fmt.Println("with err:", err)
	}
	fmt.Println("please check mailbox")

}
