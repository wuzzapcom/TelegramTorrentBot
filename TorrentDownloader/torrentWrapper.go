package TorrentDownloader

import (
	"github.com/anacrolix/torrent"
	"time"
)

type Torrent struct{

	t *torrent.Torrent

}

func NewTorrent(torr *torrent.Torrent) (t *Torrent){

	t = &Torrent{}

	t.t = torr

	return

}

func NewTorrents(torrents []*torrent.Torrent) (t []*Torrent){

	t = make([](*Torrent), len(torrents))

	for i, torr := range torrents {

		t[i] = NewTorrent(torr)

	}

	return

}

func (t *Torrent) GetProgress() float64{

	return float64(t.t.BytesCompleted()) / float64(t.t.Length()) * 100

}

func (t *Torrent) GetSize() int64{

	return t.t.Length()

}

func (t *Torrent) GetFilenames() []string{

	files := t.t.Files()
	res := make([]string, 0, len(files))

	for _, file := range files {

		res = append(res, file.Torrent().Name())

	}

	return res

}

func (t *Torrent) GetDownloadSpeed() float64{

	startValue := t.t.BytesCompleted()

	time.Sleep(time.Second / 10)

	endValue := t.t.BytesCompleted()

	delta := 10 * (endValue - startValue)

	delta = delta / 1024

	return  float64(delta) / 1024

}


func (t *Torrent) IsDownloaded() bool {

	return  t.t.BytesCompleted() == t.t.Length()

}

func (t *Torrent) GetName() string {

	return t.t.Name()

}
