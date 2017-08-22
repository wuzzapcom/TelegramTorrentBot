package Constants

const HELP_MESSAGE = "Вы сейчас общаетесь со своим BitTorrent сервером.\n Чтобы получить информацию о скачиваемых торрентах, введите /checkTorrents. Чтобы скачать новый торрент, отправьте в чат .torrent файл. Чтобы получить список скачанных торрентов, введите /getFiles."
const UNKNOWN_COMMAND = "Не понимаю, чего вы хотите, повторите запрос."
const UNKNOWN_USER = "Ваши права доступа не подтверждены, вы не являетесь хозяином. До свидания."
const NO_TORRENTS_DOWNLOADING = "У вас нет ни одного торрента."
const FILE_DOWNLOADED_1 = "Торрент "
const FILE_DOWNLOADED_2 = " скачался!"
const TORRENT_STARTED = "Торрент начал скачиваться!"
const NO_FILES_DOWNLOADED = "Вы пока не скачали ни один торрент."
const PICK_FILES_TO_DOWNLOAD = "Выберите файлы, которые стоит скачать. Вводите номера файлов цифрами, каждая цифра - новая строка. Если нужно скачать все файлы, то введите одно число -1. Нумерация с нуля."
const BOT_OWNER_ID = 36187514

const INFO_FILE_NAME = ".TelegramTorrentBotInfo"

const (
	WAIT_FOR_FILES_TO_DOWNLOAD_STATE = iota
	NO_STATE
)

/*
	TODO

	0) FIX BUG WITH FINISH NOTIFICATION : return func in separate goroutine, keep information about number of bytes torrent should download and compare with it, not with all bytes of torrent
	1) Add way to finish downloading files
	2) Add emoji for visibility
	3) TODO from torrentDownloader : add array of torrents for simultaneous work with few torrents

*/
