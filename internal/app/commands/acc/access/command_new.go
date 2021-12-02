package access

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/acc/access"
)

func (c *accAccessCommander) New(inputMessage *tgbotapi.Message) {
	var outMsgText string

	acc := access.Access{}

	err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &acc)
	if err != nil {
		outMsgText = "Wrong argument. Example: {\"role_id\":1,\"resource_id\":1}"
	} else {
		id, err := c.accessService.Create(acc)
		if err != nil {
			outMsgText = "Create access error"
		} else {
			acc.ID = id
			acc_out, _ := json.Marshal(acc)
			outMsgText = "New access: " + string(acc_out)
		}
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outMsgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("AccAccessCommander.Get: error sending reply message to chat - %v", err)
	}
}
