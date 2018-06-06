package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

const programmApiUrl = "http://pr0gramm.com/api/items/get?"
const programmDownloadUrl = "http://img.pr0gramm.com/"

func (c Command) Pr0() {
	downloadableImage, err := CallPr0API(buildAPIURL(c.params), ".jpg", ".png")
	if err != nil {
		SendError(c)
	}
	sendImage(downloadableImage, c)
}

func (c Command) Pr0gif() {
	downloadableImage, err := CallPr0API(buildAPIURL(c.params), ".gif", "")
	if err != nil {
		SendError(c)
	}
	sendDocument(downloadableImage, c)
}

func (c Command) Pr0vid() {
	downloadableImage, err := CallPr0API(buildAPIURL(c.params), ".mp4", ".webm")
	if err != nil {
		SendError(c)
	}
	sendDocument(downloadableImage, c)
}

func sendDocument(url string, c Command) {
	m := tgbotapi.NewDocumentUpload(c.m.Chat.ID, downloadAndcreateFileByte(url))
	c.bot.Send(m)
}

func SendError(c Command) {
	sticker := tgbotapi.NewStickerShare(c.m.Chat.ID, "CAADAgAD-SYAAktqAwABJKWBkTWlpzoC")
	msg := tgbotapi.NewMessage(c.m.Chat.ID, "`nichts gefunden`")
	msg.ParseMode = "Markdown"
	c.bot.Send(sticker)
	c.bot.Send(msg)
}

func sendImage(url string, c Command) {
	m := tgbotapi.NewPhotoUpload(c.m.Chat.ID, downloadAndcreateFileByte(url))
	m.Caption = "Stibot v2 hat immernoch nichts mit dem Content zu tun."
	c.bot.Send(m)
}

func downloadAndcreateFileByte(url string) tgbotapi.FileBytes {
	b := download(url)
	fb := tgbotapi.FileBytes{
		Name:  url,
		Bytes: b,
	}
	return fb
}

func CallPr0API(url string, filter string, filter2 string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	d := programmResponse{}
	json.NewDecoder(response.Body).Decode(&d)

	items := filterItems(d.Items, filter)
	if filter2 != "" {
		print("filter 2")
		items = append(items, filterItems(d.Items, filter2)...)
	}
	if len(items) > 0 {
		return items[rand.Intn(len(items))].Image, nil
	}
	return "", errors.New("Aint no Shit")

}

func filterItems(list []item, filter string) (ret []item) {
	for _, i := range list {
		if strings.HasSuffix(i.Image, filter) {
			ret = append(ret, i)
		}
	}
	return
}

func buildAPIURL(params []string) string {
	urlparams := "flags=1"
	if len(params) > 0 {
		switch params[0] {
		case "nsfw":
			urlparams = "flags=2"
			if len(params) >= 2 {
				urlparams += "&tags=" + strings.Join(params[1:], "+")
			}
		default:
			urlparams += "&tags=" + strings.Join(params, "+")
		}
	}
	print("debug output: " + programmApiUrl + urlparams)
	return programmApiUrl + urlparams
}

type programmResponse struct {
	Cache string `json:"cache"`
	Items []item `json:"items"`
}

type item struct {
	Image string `json:"image"`
}
