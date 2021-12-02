package access

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c accAccessCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("AccAccessCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMsgText := ""

	accesses, isLast := c.accessService.List(uint64(parsedData.Offset), rows_per_page)
	for _, p := range accesses {
		outputMsgText += c.accessService.String(p)
		outputMsgText += "\n"
	}
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	if !isLast {
		serializedData, _ := json.Marshal(CallbackListData{
			Offset: int(rows_per_page) + parsedData.Offset,
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
