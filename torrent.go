package main

import (
	//"wuzzapcom/TelegramTorrentBot/TorrentDownloader"
	//"fmt"
	"wuzzapcom/TelegramTorrentBot/Telegram"
	"log"
	"io/ioutil"
)

func main(){

	telegram, err  := Telegram.NewTelegram(openAuthFile("auth.token"), "/Users/wuzzapcom/", "/Users/wuzzapcom")
	if err != nil{
		log.Println(err)
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

func testTorrent(){

	//torrentDownloader, err := TorrentDownloader.NewTorrentDownloader()
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//torrents := torrentDownloader.GetTorrents()
	//
	//fmt.Println(len(torrents))
	//
	////for _, torr := range torrents{
	////	fmt.Println(torr.GetProgress())
	////}
	//
	////go torrentDownloader.DownloadTorrent("/Users/wuzzapcom/LinkinPark.torrent")
	////
	////for len(torrentDownloader.GetTorrents()) == 0 {}
	////
	////for /*len(torrentDownloader.GetTorrents()) == 0 || */!torrentDownloader.GetTorrents()[0].IsDownloaded(){
	////
	////	fmt.Println(torrentDownloader.GetTorrents()[0].GetProgress())
	////
	////}
	//
	//torrentDownloader.Close()

}



