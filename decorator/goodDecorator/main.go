package main

import "fmt"

// Объяснение кода
//
// Базовый интерфейс Notifier — имеет метод Send(message string), который все
// виды оповещений должны реализовать. Класс EmailNotifier — базовая реализация,
// которая отправляет сообщение по электронной почте. Декораторы SMSNotifier и
// SlackNotifier — принимают объект Notifier и добавляют дополнительные способы
// отправки уведомлений (по SMS и через Slack соответственно). Метод Send в
// каждом декораторе вызывает метод Send у декорированного объекта, добавляя при
// этом своё поведение. Результат выполнения кода При запуске программы, если
// вызовется notifierWithSMSAndSlack.Send("Critical system error!"), то будет
// выведено:

// Notifier - интерфейс для всех типов оповещений
type Notifier interface {
	Send(message string)
}

// EmailNotifier - базовая реализация Notifier, отправляет уведомление по email
type EmailNotifier struct {
	email string
}

// NewEmailNotifier - конструктор для EmailNotifier
func NewEmailNotifier(email string) *EmailNotifier {
	return &EmailNotifier{email: email}
}

// Send - отправка сообщения по email
func (e *EmailNotifier) Send(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.email, message)
}

// SMSNotifier - декоратор для добавления SMS-оповещения
type SMSNotifier struct {
	Notifier
	phoneNumber string
}

// NewSMSNotifier - конструктор для SMSNotifier
func NewSMSNotifier(notifier Notifier, phoneNumber string) *SMSNotifier {
	return &SMSNotifier{
		Notifier:    notifier,
		phoneNumber: phoneNumber,
	}
}

// Send - отправка сообщения по SMS и вызов метода Send у базового Notifier
func (s *SMSNotifier) Send(message string) {
	s.Notifier.Send(message)
	fmt.Printf("Sending SMS to %s: %s\n", s.phoneNumber, message)
}

// SlackNotifier - декоратор для добавления Slack-оповещения
type SlackNotifier struct {
	Notifier
	channel string
}

// NewSlackNotifier - конструктор для SlackNotifier
func NewSlackNotifier(notifier Notifier, channel string) *SlackNotifier {
	return &SlackNotifier{
		Notifier: notifier,
		channel:  channel,
	}
}

// Send - отправка сообщения в Slack и вызов метода Send у базового Notifier
func (s *SlackNotifier) Send(message string) {
	s.Notifier.Send(message)
	fmt.Printf("Sending Slack message to channel %s: %s\n", s.channel, message)
}

// Основная функция
func main() {
	// Создаем базовый EmailNotifier
	notifier := NewEmailNotifier("admin@example.com")

	// Оборачиваем его в SMSNotifier
	notifierWithSMS := NewSMSNotifier(notifier, "+123456789")

	// Оборачиваем его в SlackNotifier
	notifierWithSMSAndSlack := NewSlackNotifier(notifierWithSMS, "#alerts")

	// Отправляем уведомление, используя все декораторы
	notifierWithSMSAndSlack.Send("Critical system error!")
}
