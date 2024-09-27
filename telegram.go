package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func telegramBot() {
	// 使用你的 Telegram Bot Token
	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	// 設置 Bot 的 debug 模式
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	url := config.LaunchURL
	for update := range updates {
		if update.Message != nil && update.Message.Text == "/start"{
            // 創建一個 WebApp 按鈕
            webAppInfo := tgbotapi.WebAppInfo{
                URL: url, // 你的 URL
            }
            webAppButton := tgbotapi.NewInlineKeyboardButtonWebApp("Open WebApp", webAppInfo)

            // 創建鍵盤
            keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(webAppButton))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welecom to cactus-club")
			msg.ReplyMarkup = keyboard

			// 發送消息
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}

