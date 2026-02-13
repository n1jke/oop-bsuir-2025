package infrastructure

import "fmt"

// SmtpMailer - имитация почтового сервиса
type SmtpMailer struct {
	server string
}

func NewSmtpMailer(svr string) *SmtpMailer {
	return &SmtpMailer{server: svr}
}

func (s *SmtpMailer) Notify(to string, subject string, body string) {
	fmt.Printf(">> Connecting to SMTP server %s...\n", s.server)
	fmt.Printf(">> Sending EMAIL to %s\n   Subject: %s\n   Body: %s\n", to, subject, body)
}

// TelegramMailer - имитация бота в телеграмм дл отправки сообщений менеджеру
type TelegramMailer struct {
	connString string
}

func NewTelegramMailer(conn string) *TelegramMailer {
	return &TelegramMailer{connString: conn}
}

func (t *TelegramMailer) Notify(to string, subject string, body string) {
	fmt.Printf(">> Connecting to Telegram bot %s...\n", t.connString)
	fmt.Printf(">> Sending MESSAGE to %s\n   Subject: %s\n   Body: %s\n", to, subject, body)
}
