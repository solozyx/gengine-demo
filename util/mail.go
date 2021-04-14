package util

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"gengine/common"
)

func SendMail(to string, code string) error {
	sender := &MailSender{
		NoSSL:    true,
		Username: common.Cfg.Mail.Username,
		Password: common.Cfg.Mail.Password,
		Host:     "smtp.163.com",
		Port:     465,
	}
	content := fmt.Sprintf("验证码是 %s", code)
	err := sender.SendMail("验证码", content, []string{to}, map[string]string{
		"Content-Type": "text/plain; charset=\"utf-8\"",
	})
	if err != nil {
		logrus.Errorf("SendMail err=%v", err.Error())
	}
	return err
}

type fakeAuth struct {
	smtp.Auth
}

func (a fakeAuth) Start(server *smtp.ServerInfo) (string, []byte,
	error) {
	server.TLS = true
	return a.Auth.Start(server)
}

//MailSender 邮件发送人
type MailSender struct {
	NoSSL    bool
	Username string
	Password string
	Host     string
	Port     int
}

type MailContent struct {
	Subject     string
	Content     string
	Contact     []string
	ContentType string
	Other       map[string]string
}

func (sender *MailSender) Send(mailConent *MailContent) error {
	if mailConent.Other == nil {
		mailConent.Other = map[string]string{}
	}
	if mailConent.ContentType == "" {
		mailConent.ContentType = "text/plain; charset=\"utf-8\""
	}
	mailConent.Other["Content-Type"] = mailConent.ContentType
	return sender.SendMail(mailConent.Subject, mailConent.Content, mailConent.Contact, mailConent.Other)

}

//SendMail 发送邮件
//subject 邮件标题
//content 邮件内容
//contact 联系人
func (sender *MailSender) SendMail(subject string, content string, contact []string, args ...map[string]string) error {
	if len(contact) == 0 {
		return fmt.Errorf("联系人不能为空. Contact cannot be empty!\n")
	}
	from := mail.Address{"", sender.Username}
	//to := mail.Address{"", strings.Join(contact, ";")}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = strings.Join(contact, ";")
	headers["Subject"] = subject
	if args != nil {
		arg := args[0]
		for key, value := range arg {
			headers[key] = value
		}
	}
	// Build Email Content
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + content
	auth := smtp.PlainAuth("", sender.Username, sender.Password, sender.Host)
	if sender.NoSSL {
		auth = fakeAuth{auth}
	}

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         sender.Host}

	conn, err := tls.Dial("tcp", sender.Host+":"+strconv.Itoa(sender.Port), tlsconfig)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, sender.Host)
	if err != nil {
		return err
	}
	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(from.Address); err != nil {
		return err
	}
	for _, addr := range contact {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()
	return nil
}
