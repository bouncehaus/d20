package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// func roller() {
// 	rand.Seed(time.Now().UnixNano())
// 	min := 1
// 	max := 20
// 	response := (rand.Intn(max-min+1) + min)
// 	fix := strconv.Itoa(response)
// }

func main() {
	sess, err := discordgo.New("$token")
	if err != nil {
		log.Fatal(err)
	}

	// rand.Seed(time.Now().UnixNano())
	// min := 1
	// max := 20
	// //fmt.Println("You rolled:", rand.Intn(max-min+1)+min)
	// response := (rand.Intn(max-min+1) + min)
	// fix := strconv.Itoa(response)

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "!roll" {
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 20
			//fmt.Println("You rolled:", rand.Intn(max-min+1)+min)
			response := (rand.Intn(max-min+1) + min)
			fix := strconv.Itoa(response)
			s.ChannelMessageSend(m.ChannelID, fix)
		}

	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("the bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
