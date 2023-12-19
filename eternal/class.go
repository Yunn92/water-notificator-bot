package eternal

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	BotAPI    *tgbotapi.BotAPI
	UserList  map[int64]*User
	Keyboards *KeyboardList
	Timer     *Timer
}

func BotInitilisation() Bot {
	return Bot{
		UserList:  make(map[int64]*User),
		Keyboards: NewKeyboardList(),
		Timer:     MakerNewTimer(),
	}
}

// Создаёт нового пользователя со стандартными настройками
func (b *Bot) SayHello(id int64) {
	msg := tgbotapi.NewMessage(id, "Салам, это бот для того чтобы клуши не забывали пить воду")
	msg.ReplyMarkup = b.Keyboards.TrunOnButton

	b.UserList[id] = MakerUser()

	if _, err := b.BotAPI.Send(msg); err != nil {
		log.Println(err)
	}
}

// Функция отображения настроек
func (b *Bot) ShowSettings(id int64) {
	var isEnabled string
	var turnOffButton, turnOnButton = tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Отключить напоминание")), tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Включить напоминание"))

	switch b.UserList[id].Settings.IsEnabled {
	case true:
		isEnabled = "Да"
		b.Keyboards.Settings.Keyboard = append(b.Keyboards.Settings.Keyboard, turnOffButton)
	case false:
		isEnabled = "Нет"
		b.Keyboards.Settings.Keyboard = append(b.Keyboards.Settings.Keyboard, turnOnButton)
	}

	msgText := fmt.Sprintf("Текущие настройки:\nБот включен: %s\nТекщий интервал напоминаний: %s час", isEnabled, strconv.Itoa(b.UserList[id].Settings.Interval))

	msg := tgbotapi.NewMessage(id, msgText)
	msg.ReplyMarkup = b.Keyboards.Settings

	b.BotAPI.Send(msg)
}

func (b *Bot) TellAJoke(id int64) {
	msg := tgbotapi.NewMessage(id, "Главное правило хорошего раввина: \n\"Не покупай кольца кальмара около синагоги\"")

	b.BotAPI.Send(msg)
}

type User struct {
	Id       int64
	Step     string
	Settings *UserSettings
}

func MakerUser() *User {
	return &User{
		Id:       0,
		Step:     "initiation",
		Settings: NewDefaultSettings(),
	}
}

func (u *User) ChangeMode() {
	u.Settings.IsEnabled = !u.Settings.IsEnabled
}

type UserSettings struct {
	Interval  int
	IsEnabled bool
}

func NewDefaultSettings() *UserSettings {
	return &UserSettings{
		Interval:  1,
		IsEnabled: false,
	}
}

type Timer struct {
	NextNotification time.Time
}

func MakerNewTimer() *Timer {
	return &Timer{
		NextNotification: time.Time{},
	}
}

type KeyboardList struct {
	Main     *tgbotapi.ReplyKeyboardMarkup
	Settings *tgbotapi.ReplyKeyboardMarkup

	TrunOnButton *tgbotapi.ReplyKeyboardMarkup
	ChooseMode   *tgbotapi.InlineKeyboardMarkup
}

func NewKeyboardList() *KeyboardList {
	return &KeyboardList{
		Main:         MainKeyboard(),
		Settings:     Settings(),
		TrunOnButton: TurnOnButton(),
		ChooseMode:   TypesOfEnable(),
	}
}
