package access

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	rows_per_page = uint64(3)
)

func (c *accAccessCommander) List(inputMessage *tgbotapi.Message) {

	pageData := CallbackListData{}
	err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &pageData)

	startIndex := uint64(0)

	if err != nil {
		startIndex = uint64(pageData.Offset)
	}

	outputMsgText := "Here all the accesses: \n\n"

	accesses, isLast := c.accessService.List(startIndex, rows_per_page)
	for _, p := range accesses {
		outputMsgText += c.accessService.String(p)
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if !isLast {
		serializedData, _ := json.Marshal(CallbackListData{
			Offset: int(rows_per_page),
		})

		callbackPath := path.CallbackPath{
			Domain:       "acc",
			Subdomain:    "access",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("AccAccessCommander.List: error sending reply message to chat - %v", err)
	}
}
