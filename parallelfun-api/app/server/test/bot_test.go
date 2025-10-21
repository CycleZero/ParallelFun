package test

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/msg"
	"github.com/Tnze/go-mc/bot/playerlist"
	"log"
	"sync"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	client := bot.NewClient()
	// ...
	player := basic.NewPlayer(client, basic.DefaultSettings, basic.EventsListener{})
	err := client.JoinServer("server.poyuan233.cn:25565")
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	log.Println("Login success")
	go func() {
		for {
			err = client.HandleGame()
			if err != nil {
				log.Fatal(err)
			}
		}
		wg.Done()
	}()
	wg.Add(1)
	playerList := playerlist.New(client)
	chatHandler := msg.New(client, player, playerList, msg.EventsHandler{})
	go func() {
		for {
			err = chatHandler.SendMessage("hello")

			if err != nil {
				log.Println(err)
			} else {
				log.Println("send message success", "hello")
			}
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Add(1)
	wg.Wait()

}
