package Telegram

import (
	"github.com/Syfaro/telegram-bot-api"
	"wuzzapcom/TelegramTorrentBot/TorrentDownloader"
	"log"
	"net/http"
	"os"
	"io"
	"wuzzapcom/TelegramTorrentBot/Constants"
	"time"
	"wuzzapcom/TelegramTorrentBot/FileManager"
)

type Telegram struct{
	bot     *tgbotapi.BotAPI
	updates <-chan tgbotapi.Update
	torrentDownloader *TorrentDownloader.TorrentDownloader
	torrentFilesPath string
}

func NewTelegram(authToken string, torrentFilesPath string, torrentDataPath string) (telegram *Telegram, err error){

	telegram = &Telegram{}

	telegram.torrentDownloader, err = TorrentDownloader.NewTorrentDownloader(torrentDataPath)
	if err != nil {
		return  nil, err
	}

	telegram.bot, err = tgbotapi.NewBotAPI(authToken)

	if err != nil {
		return
	}

	config := telegram.configureAPI()

	telegram.updates, err = telegram.bot.GetUpdatesChan(config)
	if err != nil{
		return nil, err
	}

	telegram.torrentFilesPath = torrentFilesPath

	return

}

func (telegram *Telegram) Start(){

	go telegram.sendNotification()

	for update := range telegram.updates{ //update.Message.Chat.ID

		if update.Message == nil{
			continue
		}

		telegram.handleUpdate(update)
	}

}

func (telegram *Telegram) handleUpdate(update tgbotapi.Update){

	if update.Message.Chat.ID != Constants.BOT_OWNER_ID {

		telegram.sendMessage(Constants.UNKNOWN_USER, update.Message.Chat.ID)

	}

	if update.Message.Command() == "help" {

		telegram.sendMessage(Constants.HELP_MESSAGE, update.Message.Chat.ID)

	}else if update.Message.Command() == "checkTorrents" {

		telegram.sendMessage(telegram.checkTorrents(), update.Message.Chat.ID)

	}else if update.Message.Command() == "getFiles" {

		telegram.sendMessage(telegram.getFiles(), update.Message.Chat.ID)

	}else {

		if update.Message.Document != nil {

			filename, err := telegram.downloadFile(update.Message.Document.FileID)
			if err != nil {
				log.Println(err)
			}else {
				telegram.downloadTorrent(filename)
				telegram.sendMessage(Constants.TORRENT_STARTED, update.Message.Chat.ID)
			}

		}else {

			telegram.sendMessage(Constants.UNKNOWN_COMMAND, update.Message.Chat.ID)

		}

	}


}

func (telegram *Telegram) getFiles() string{

	fileManager := FileManager.InitFileManager(telegram.torrentFilesPath)
	dataArray := fileManager.GetListOfFiles()

	result := ""

	for _, data := range dataArray{

		result += data.ToString()

	}

	if result == "" {
		return Constants.NO_FILES_DOWNLOADED
	}

	return result

}

func (telegram *Telegram) sendNotification(){

	finishedTorrents := make([]int, 1)
	finishedTorrents[0] = -1
	breakCycle := false

	for {

		torrents := telegram.torrentDownloader.GetTorrents()

		breakCycle = false

		for i, torr := range torrents {

			for _, value := range finishedTorrents{
				if i == value{
					breakCycle = true
				}
			}

			if torr.IsDownloaded() && !breakCycle {

				telegram.addInfo(torr)

				telegram.sendMessage(Constants.FILE_DOWNLOADED_1+torr.GetName()+Constants.FILE_DOWNLOADED_2, Constants.BOT_OWNER_ID)
				finishedTorrents = append(finishedTorrents, i)

			}

		}

		time.Sleep(time.Second)

	}

}

func (telegram *Telegram) addInfo(torrent *TorrentDownloader.Torrent){

	fileManager := FileManager.InitFileManager(telegram.torrentFilesPath)

	data := FileManager.Data{
		PathToSource : telegram.torrentFilesPath,
		SizeOfSource : torrent.GetSize(),
		Name : torrent.GetName(),
		FileNames: torrent.GetFileNames(),
	}

	fileManager.Add(data)
	fileManager.Save()

}


func (telegram *Telegram) checkTorrents() string {

	return telegram.torrentDownloader.GetListOfTorrents()

}

func (telegram *Telegram) downloadTorrent(filename string){

	telegram.torrentDownloader.DownloadTorrent(filename)

}

func (telegram *Telegram) downloadFile(documentID string) (string,error) {

	fileConfig := tgbotapi.FileConfig{FileID: documentID}

	file, err := telegram.bot.GetFile(fileConfig)
	if err != nil {
		return "", err
	}

	link := file.Link(telegram.bot.Token)

	log.Println("Link : " + link)

	filename, err := telegram.downloadFileFromLink(file.FilePath, link)
	if err != nil{
		return "", err
	}

	return filename, nil

}

func (telegram *Telegram) downloadFileFromLink(filename, link string) (string, error){

	response, err := http.Get(link)
	if err != nil{
		return "", err
	}

	log.Println("filename is " + telegram.torrentFilesPath + filename)
	file, err := os.Create(telegram.torrentFilesPath + filename)
	if err != nil{
		return "", err
	}

	log.Println(response.Status)

	_, err = io.Copy(file, response.Body)
	if err != nil{
		return "", err
	}

	file.Close()
	response.Body.Close()

	return telegram.torrentFilesPath + filename, nil

}

func (telegram *Telegram) sendMessage(message string, id int64){

	log.Println("Send message to user with message : " + message + "id : " + string(id))

	telegram.bot.Send(tgbotapi.NewMessage(id, message))

}

func (telegram *Telegram) configureAPI() tgbotapi.UpdateConfig{

	telegram.bot.Debug = false
	config := tgbotapi.NewUpdate(0)//todo check in documentation for value
	config.Timeout = 60 //todo check in documentation for value

	return config //TODO configure this

}
