package Constants

const HELP_MESSAGE = "Вы сейчас общаетесь со своим BitTorrent сервером.\n Чтобы получить информацию о скачиваемых торрентах, введите /checkTorrents. Чтобы скачать новый торрент, отправьте в чат .torrent файл. Чтобы получить список скачанных торрентов, введите /getFiles."
const UNKNOWN_COMMAND = "Не понимаю, чего вы хотите, повторите запрос."
const UNKNOWN_USER = "Ваши права доступа не подтверждены, вы не являетесь хозяином. До свидания."
const NO_TORRENTS_DOWNLOADING = "У вас нет ни одного торрента."
const FILE_DOWNLOADED_1 = "Торрент "
const FILE_DOWNLOADED_2 = " скачался!"
const TORRENT_STARTED = "Торрент начал скачиваться!"
const BOT_OWNER_ID =  36187514

const INFO_FILE_NAME = ".TelegramTorrentBotInfo"

/*

	1) Finish FileManager
		* FileManager should show all subfiles, such as tv series episodes
	2) Implement FileManager in main project
	3) Don`t call DownloadAll function in TorrentDownloader, but download all files consistently(for convenient downloading tv series)
	4) Send message to user about every downloaded file/episode (optional, user should set this as flag manually)

 */
