package third

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

// @author: [maplepie](https://github.com/maplepie)
// @function: Email
// @description: Email发送方法
// @param: subject string, body string
// @return: error

type Email struct {
	To       string `json:"to"`
	Port     int    `json:"port"`
	From     string `json:"from"`
	Host     string `json:"host"`
	IsSSL    bool   `json:"is_ssl"`
	Secret   string `json:"secret"`
	Nickname string `json:"nickname"`
}

func (e *Email) Send(subject string, body string) error {
	to := strings.Split(e.To, ",")
	return e.send(to, subject, body)
}

// @author: [SliverHorn](https://github.com/SliverHorn)
// @function: ErrorToEmail
// @description: 给email中间件错误发送邮件到指定邮箱
// @param: subject string, body string
// @return: error

func (e *Email) ErrorToEmail(subject string, body string) error {
	to := strings.Split(e.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return e.send(to, subject, body)
}

// @author: [maplepie](https://github.com/maplepie)
// @function: EmailTest
// @description: Email测试方法
// @param: subject string, body string
// @return: error

func (e *Email) EmailTest(subject string, body string) error {
	to := []string{e.From}
	return e.send(to, subject, body)
}

// @author: [maplepie](https://github.com/maplepie)
// @function: send
// @description: Email发送方法
// @param: subject string, body string
// @return: error

func (e *Email) send(to []string, subject string, body string) error {
	from := e.From
	nickname := e.Nickname
	secret := e.Secret
	host := e.Host
	port := e.Port
	isSSL := e.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	email := email.NewEmail()
	if nickname != "" {
		email.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		email.From = from
	}
	email.To = to
	email.Subject = subject
	email.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = email.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = email.Send(hostAddr, auth)
	}
	return err
}
