package main

import (
	"github.com/KrisInferno/PocketBot/pkg/repository"
	"github.com/KrisInferno/PocketBot/pkg/repository/boltdb"
	"github.com/KrisInferno/PocketBot/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1635881340:AAFQHRfSYiw1b59pFYN1AQusKg9PJLd5Lf8")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("98014-5129f80dad64c668d4079cf4")
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost/")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
