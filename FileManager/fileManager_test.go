package FileManager

import (
	"testing"
	"reflect"
	"io/ioutil"
	"os"
)

func TestFileManager(t *testing.T){

	dataString := "{\"Data\":[{\"PathToSource\":\"/home\",\"SizeOfSource\":999,\"DownloadDate\":532145,\"Name\":\"Test1\",\"FileNames\":[\"File1\",\"File2\",\"File3\"]}]}"

	fileManager := InitFileManager("/home/wuzzapcom/Coding/Golang/src/wuzzapcom/TelegramTorrentBot")
	data := Data{
		PathToSource:"/home",
		SizeOfSource:999,
		DownloadDate:532145,
		Name: "Test1",
		FileNames:[]string{"File1", "File2", "File3"}}

	fileManager.Add(data)

	res := fileManager.GetListOfFiles()
	if len(res) != 1 && !reflect.DeepEqual(res[0], data) {
		t.Error("Failed : readed data is wrong. Expected : true got : false")
	}

	fileManager.Save()

	b, err := ioutil.ReadFile("/home/wuzzapcom/Coding/Golang/src/wuzzapcom/TelegramTorrentBot/.TelegramTorrentBotInfo")
	if err != nil{
		t.Error("Failed to read .TelegramTorrentBotInfo\n", err)
	}

	if string(b) != dataString{
		t.Error("Readed .TelegramTorrentBotInfo contains wrong data . Expected : " + dataString + "got : " + string(b))
	}

	os.Remove("/home/wuzzapcom/Coding/Golang/src/wuzzapcom/TelegramTorrentBot/.TelegramTorrentBotInfo")

}
