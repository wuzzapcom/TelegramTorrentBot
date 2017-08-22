package TorrentDownloader

import (
	"fmt"
	"strconv"
	"wuzzapcom/TelegramTorrentBot/Constants"

	"github.com/anacrolix/torrent"
)

type TorrentDownloader struct {
	client         *torrent.Client
	currentTorrent *torrent.Torrent //TODO in next version add array of torrents for simultaneous work with few torrents

}

func NewTorrentDownloader(torrentDataPath string) (torrentDownloader *TorrentDownloader, err error) {

	torrentDownloader = &TorrentDownloader{}

	config := &torrent.Config{}
	config.DataDir = torrentDataPath

	torrentDownloader.client, err = torrent.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}

	return

}

func (torrentDownloader *TorrentDownloader) GetCurrentTorrent() *Torrent {

	return NewTorrent(torrentDownloader.currentTorrent)

}

func (torrentDownloader *TorrentDownloader) GetTorrents() (torrents []*Torrent) {

	torrents = NewTorrents(torrentDownloader.client.Torrents())

	return

}

func (torrentDownloader *TorrentDownloader) GetFilenamesFromTorrent() string {

	torrent := NewTorrent(torrentDownloader.currentTorrent)

	torrentNames := torrent.GetFileNames()

	result := ""

	for _, name := range torrentNames {

		result += name + "\n"

	}

	return result

}

func (torrentDownloader *TorrentDownloader) DownloadTorrent(downloadFilesIndexes []int) {

	if len(downloadFilesIndexes) == -1 && downloadFilesIndexes[0] == -1 {
		torrentDownloader.currentTorrent.DownloadAll()
	}

	for i, file := range torrentDownloader.currentTorrent.Files() {

		for _, j := range downloadFilesIndexes {

			if i == j {

				file.Download()

			}

		}

	}

}

func (torrentDownloader *TorrentDownloader) AddTorrent(path string) {

	t, err := torrentDownloader.client.AddTorrentFromFile(path)
	if err != nil {
		fmt.Println(err)
		torrentDownloader.currentTorrent = nil
	}

	torrentDownloader.currentTorrent = t

}

func (torrentDownloader *TorrentDownloader) GetListOfTorrents() (result string) {

	torrents := torrentDownloader.GetTorrents()

	if len(torrents) == 0 {
		return Constants.NO_TORRENTS_DOWNLOADING
	}

	for i, torr := range torrents {

		if torr.IsDownloaded() {
			continue
		}

		result += strconv.Itoa(i+1) + ") " + torr.GetName() + "   " + strconv.FormatFloat(torr.GetDownloadSpeed(), 'f', 2, 64) + "mb/s   " + strconv.FormatFloat(torr.GetProgress(), 'f', 2, 64) + "%\n"

	}

	return

}

func (torrentDownloader *TorrentDownloader) Close() {
	torrentDownloader.client.Close()
}
