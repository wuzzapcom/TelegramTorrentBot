package TorrentDownloader

import "github.com/anacrolix/torrent"

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



func (t *Torrent) IsDownloaded() bool {

	return  t.t.BytesCompleted() == t.t.Length()

}

func (t *Torrent) GetName() string {

	return t.t.Name()

}
