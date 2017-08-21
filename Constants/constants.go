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

	1) Update broke notifications when not all torrent downloaded(why am i use notification func?)
	2) /getFiles shows all files from torrent, but not just downloaded files
	3) TEST THIS
	4) TODO from torrentDownloader

*/
