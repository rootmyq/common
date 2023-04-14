package com

import (
	"net/smtp"
	"strings"
)

// EmailReq 邮件信息
type EmailReq struct {
	//接收人邮箱
	To string
	//发送人邮箱
	From string

	//发送人密码
	FromPassword string

	//邮箱服务 ip+端口
	Host string

	//邮件表头
	Subject string

	//邮件内容
	Body string
}

// Mail 邮件服务
type Mail struct {
}

// SendMail 发送邮件
func (m *Mail) SendMail(emailReq EmailReq) error {
	hp := strings.Split(emailReq.Host, ":")
	auth := smtp.PlainAuth("", emailReq.From, emailReq.FromPassword, hp[0])
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + emailReq.To + "\r\nFrom: " + emailReq.From + ">\r\nSubject: " + emailReq.Subject + "\r\n" + contentType + "\r\n\r\n" + emailReq.Body)
	var sendTo []string
	sendTo = append(sendTo, emailReq.To)
	return smtp.SendMail(emailReq.Host, auth, emailReq.From, sendTo, msg)
}

//
//
//func SendToMail(user, password, host, to, subject, body, mailtype string) error {
//	hp := strings.Split(host, ":")
//	auth := smtp.PlainAuth("", user, password, hp[0])
//	var content_type string
//	if mailtype == "html" {
//		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
//	} else {
//		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
//	}
//
//	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
//	send_to := strings.Split(to, ";")
//	err := smtp.SendMail(host, auth, user, send_to, msg)
//	return err
//}
//
//func Main() {
//	user := "muyunqiang@aliyun.com"
//	password := "mu63780348mu"
//	host := "smtp.aliyun.com:25"
//	to := "340401140@qq.com"
//
//	subject := "使用Golang发送邮件"
//
//	body := `
//		<html>
//		<body>
//		<h3>
//		"Test golang send to email"
//		</h3>
//		</body>
//		</html>
//		`
//	fmt.Println("send email")
//	err := SendToMail(user, password, host, to, subject, body, "html")
//	if err != nil {
//		fmt.Println("Send mail error!")
//		fmt.Println(err)
//	} else {
//		fmt.Println("Send mail success!")
//	}

//}
