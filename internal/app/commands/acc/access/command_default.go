package access

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *accAccessCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("AccAccessCommander.Help: error sending reply message to chat - %v", err)
	}
}
