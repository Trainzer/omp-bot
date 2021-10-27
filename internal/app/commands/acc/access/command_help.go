package access

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *accAccessCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__acc__access - help\n"+
			"/get__acc__access - list access by id\n"+
			"/list__acc__access - list accesses\n"+
			"/delete__acc__access - delete access\n"+
			"/new__acc__access - new access\n"+
			"/edit__acc__access - edit access\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("AccAccessCommander.Help: error sending reply message to chat - %v", err)
	}
}
