package main

import "fmt"

// В плохом варианте для каждого возможного сочетания каналов оповещения нам
// пришлось бы создавать отдельный подкласс. Например, если у нас есть три типа
// оповещений — Email, SMS, и Slack — и пользователи могут использовать их в
// любом сочетании, то для каждого сочетания придется создавать отдельный класс.
// Число классов быстро увеличивается по мере добавления новых типов оповещений.

// Недостатки этого подхода Комбинаторный взрыв подклассов: В этом примере у нас
// всего три типа оповещений, но уже шесть классов для покрытия различных
// комбинаций. Если добавить ещё один тип (например, Telegram), то количество
// необходимых комбинаций возрастёт. При четырёх типах оповещений получится 15
// различных классов!

// Проблемы с поддержкой и маcштабируемостью: Если нам нужно изменить способ
// отправки для одного из каналов, придется вносить изменения в нескольких
// классах. Это делает код сложным для поддержки и увеличивает вероятность
// ошибок.

// Отсутствие гибкости: Мы ограничены только предопределёнными комбинациями.
// Если нужно добавить ещё один канал или убрать один из существующих, придется
// создавать новые классы или модифицировать старые, что опять же усложняет код.

// Заключение Использование паттерна Декоратор, как в предыдущем примере,
// позволяет динамически добавлять новые функции без создания комбинаторного
// количества классов.

// Интерфейс Notifier, который реализуют все классы
type Notifier interface {
	Send(message string)
}

// Класс для отправки уведомлений только по Email
type EmailNotifier struct {
	email string
}

func NewEmailNotifier(email string) *EmailNotifier {
	return &EmailNotifier{email: email}
}

func (e *EmailNotifier) Send(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.email, message)
}

// Класс для отправки уведомлений только по SMS
type SMSNotifier struct {
	phoneNumber string
}

func NewSMSNotifier(phoneNumber string) *SMSNotifier {
	return &SMSNotifier{phoneNumber: phoneNumber}
}

func (s *SMSNotifier) Send(message string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.phoneNumber, message)
}

// Класс для отправки уведомлений только в Slack
type SlackNotifier struct {
	channel string
}

func NewSlackNotifier(channel string) *SlackNotifier {
	return &SlackNotifier{channel: channel}
}

func (s *SlackNotifier) Send(message string) {
	fmt.Printf("Sending Slack message to channel %s: %s\n", s.channel, message)
}

// Класс для отправки уведомлений по Email и SMS
type EmailAndSMSNotifier struct {
	email       string
	phoneNumber string
}

func NewEmailAndSMSNotifier(email, phoneNumber string) *EmailAndSMSNotifier {
	return &EmailAndSMSNotifier{email: email, phoneNumber: phoneNumber}
}

func (e *EmailAndSMSNotifier) Send(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.email, message)
	fmt.Printf("Sending SMS to %s: %s\n", e.phoneNumber, message)
}

// Класс для отправки уведомлений по Email и в Slack
type EmailAndSlackNotifier struct {
	email   string
	channel string
}

func NewEmailAndSlackNotifier(email, channel string) *EmailAndSlackNotifier {
	return &EmailAndSlackNotifier{email: email, channel: channel}
}

func (e *EmailAndSlackNotifier) Send(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.email, message)
	fmt.Printf("Sending Slack message to channel %s: %s\n", e.channel, message)
}

// Класс для отправки уведомлений по SMS и в Slack
type SMSAndSlackNotifier struct {
	phoneNumber string
	channel     string
}

func NewSMSAndSlackNotifier(phoneNumber, channel string) *SMSAndSlackNotifier {
	return &SMSAndSlackNotifier{phoneNumber: phoneNumber, channel: channel}
}

func (s *SMSAndSlackNotifier) Send(message string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.phoneNumber, message)
	fmt.Printf("Sending Slack message to channel %s: %s\n", s.channel, message)
}

// Класс для отправки уведомлений по Email, SMS и в Slack
type EmailSMSAndSlackNotifier struct {
	email       string
	phoneNumber string
	channel     string
}

func NewEmailSMSAndSlackNotifier(email, phoneNumber, channel string) *EmailSMSAndSlackNotifier {
	return &EmailSMSAndSlackNotifier{email: email, phoneNumber: phoneNumber, channel: channel}
}

func (e *EmailSMSAndSlackNotifier) Send(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.email, message)
	fmt.Printf("Sending SMS to %s: %s\n", e.phoneNumber, message)
	fmt.Printf("Sending Slack message to channel %s: %s\n", e.channel, message)
}

func main() {
	// Если нужно отправить уведомление по Email, SMS и в Slack, выбираем соответствующий класс
	notifier := NewEmailSMSAndSlackNotifier("admin@example.com", "+123456789", "#alerts")
	notifier.Send("Critical system error!")
}
