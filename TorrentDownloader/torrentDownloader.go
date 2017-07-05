package TorrentDownloader

import (
	"github.com/anacrolix/torrent"
	"fmt"
	"strconv"
	"wuzzapcom/TelegramTorrentBot/Constants"
)

type TorrentDownloader struct{

	client *torrent.Client

}

func NewTorrentDownloader(torrentDataPath string) (torrentDownloader *TorrentDownloader, err error) {

	torrentDownloader = &TorrentDownloader{}

	config := &torrent.Config{}
	config.DataDir = torrentDataPath

	torrentDownloader.client, err = torrent.NewClient(config)
	if err != nil{
		fmt.Println(err)
	}

	return

}

func (torrentDownloader *TorrentDownloader) GetTorrents() (torrents []*Torrent) {

	torrents = NewTorrents(torrentDownloader.client.Torrents())

	return

}

func (torrentDownloader *TorrentDownloader) DownloadTorrent(path string){

	t := torrentDownloader.addTorrent(path)

	t.DownloadAll()

}

func (torrentDownloader *TorrentDownloader) addTorrent(path string) (t *torrent.Torrent) {

	t, err := torrentDownloader.client.AddTorrentFromFile(path)
	if err != nil{
		fmt.Println(err)
		return nil
	}

	return

}

func (torrentDownloader *TorrentDownloader) GetListOfTorrents() (result string) {

	torrents := torrentDownloader.GetTorrents()

	if len(torrents) == 0{
		return Constants.NO_TORRENTS_DOWNLOADING
	}

	for i, torr := range torrents{

		if torr.IsDownloaded(){
			continue
		}

		result += strconv.Itoa(i+1) + ") " + torr.GetName() + "   " + strconv.FormatFloat(torr.GetDownloadSpeed(), 'f', 2, 64) + "mb/s   " + strconv.FormatFloat(torr.GetProgress(), 'f', 2, 64) + "%\n"

	}

	return

}

func (torrentDownloader *TorrentDownloader) Close(){
	torrentDownloader.client.Close()
}

