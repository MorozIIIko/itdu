package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/MorozIIIko/itdu/internal/database"
	"github.com/MorozIIIko/itdu/internal/repository"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"models.InlineKeyboardButton"
)

//err := godotenv.Load()

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(" не удалось загрузить .env:", err)
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL не найден")

	}
	db, err := database.Connect(databaseURL)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	defer db.Close()
	err = database.Migrate(db)
	if err != nil {
		log.Fatal("Не удалось выполнить миграцию базы данных:", err)
	}

	log.Println("Подключение к БД успешно")
	log.Println("Миграция базы данных выполнена успешно")
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
		startHandler(db),
	)

	log.Println("Бот запущен")

	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/quiz",
		bot.MatchTypeExact,
		quizHandler(db),
	)
	b.RegisterHandler(
		bot.HandlerTypeCallbackQueryData,
		"",
		bot.MatchTypePrefix,
		quizCallbackHandler,
	)
	b.Start(ctx)
}

func startHandler(db *sql.DB) func(
	context.Context,
	*bot.Bot,
	*models.Update,
) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}
		telegramID := update.Message.From.ID
		username := update.Message.From.Username
		//repository.CreateUser(db, telegramID, username)
		err := repository.CreateUser(db, telegramID, username)
		if err != nil {
			log.Println("ошибка", err)
			return

		}
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   " Hello , it !",
		})
		if err != nil {
			log.Println("ошибка", err)
			return
		}
	}
}
func quizHandler(db *sql.DB) func(
	context.Context,
	*bot.Bot,
	*models.Update,
) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}
		quizzes, err := repository.GetQuizzes(db)
		if err != nil {
			log.Println("ошибка на получении квиза", err)
			return
		}

		for _, quiz := range quizzes {
			log.Println("quiz:", quiz.ID, quiz.Name)
		}

		var keyboard [][]models.InlineKeyboardButton
		for _, quiz := range quizzes {
			button := models.InlineKeyboardButton{
				Text:         quiz.Name,
				CallbackData: strconv.FormatInt(quiz.ID, 10),
			}
			keyboard = append(keyboard, []models.InlineKeyboardButton{button})
		}
		//[]models.InlineKeyboardButton{button}
		//keyboard = append(keyboard,button)
		markup := models.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		}
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Выбери опросник",
			ReplyMarkup: &markup,
		})
		if err != nil {
			log.Println("ошибка отправки сообщения", err)
			return
		}

	}
}
func quizCallbackHandler(
	ctx context.Context,
	b *bot.Bot,
	update *models.Update,
) {
	if update.CallbackQuery == nil {
		return
	}
	data := update.CallbackQuery.Data
	log.Println(data)
}
