package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить .env:", err)
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не найден")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(token)
	if err != nil {
		log.Fatal(err)
	}

	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/start",
		bot.MatchTypeExact,
		startHandler,
	)

	log.Println("Бот запущен")

	b.Start(ctx)
}

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Hello, IT!",
	})

	if err != nil {
		log.Println(err)
	}
}
