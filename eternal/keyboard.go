package eternal

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MainKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	Keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Включить напоминание"),
		),
	)

	return &Keyboard
}

func Settings() *tgbotapi.ReplyKeyboardMarkup {
	Keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Настроить интервал"),
		),
	)

	return &Keyboard
}

func TypesOfEnable() *tgbotapi.InlineKeyboardMarkup {
	Keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Включить на день", "enable_day"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Включить на 3 дня", "enable_day"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Просто включить", "enable_day"),
		),
	)

	return &Keyboard
}

func TurnOnButton() *tgbotapi.ReplyKeyboardMarkup {
	Keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Включить напоминание"),
		),
	)

	Keyboard.OneTimeKeyboard = true

	return &Keyboard
}
