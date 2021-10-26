package access

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *accAccessCommander) Edit(inputMessage *tgbotapi.Message) {
	var outMsgText string

	outMsgText = "Edit not implemented"
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outMsgText,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("AccAccessCommander.Get: error sending reply message to chat - %v", err)
	}
}
