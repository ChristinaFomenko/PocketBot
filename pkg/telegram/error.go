package telegram

import "errors"

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

//msg.Text = "Ты не авторизован! Используй команду /start."

func (b *Bot) handleError(chatID int64, err error) {
	switch err {

	}
}
