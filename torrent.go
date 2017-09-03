package main

import (
	"github.com/wuzzapcom/TelegramTorrentBot/Telegram"
	"log"
	"io/ioutil"
	"os"
	"fmt"
)

func main(){

	if len(os.Args) != 4 {
		fmt.Println("You should provide path to auth file, path to folder with torrent files and path to folder with torrent data.")
	}

	telegram, err  := Telegram.NewTelegram(openAuthFile(os.Args[1]), os.Args[2], os.Args[3])
	if err != nil{
		log.Fatal(err)
	}
	telegram.Start()
}

func openAuthFile(pathToAuthFile string) string{

	data, err := ioutil.ReadFile(pathToAuthFile)

	if err != nil{
		log.Println("Failed reading from file")
		log.Println(err)
		panic(err)
	}

	log.Println(string(data))

	return string(data)

}
