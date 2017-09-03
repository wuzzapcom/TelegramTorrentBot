package FileManager

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestFileManager(t *testing.T) {

	dataString := "{\"Data\":[{\"PathToSource\":\"/home\",\"SizeOfSource\":999,\"Name\":\"Test1\",\"FileNames\":[{\"S\":\"File1\",\"B\":false},{\"S\":\"File2\",\"B\":true},{\"S\":\"File3\",\"B\":false}]}]}"

	gopath := os.Getenv("GOPATH")

	fileManager := InitFileManager(gopath + "/src/github.com/wuzzapcom/TelegramTorrentBot/")
	data := Data{
		PathToSource: "/home",
		SizeOfSource: 999,
		Name:         "Test1",
		FileNames:    []Pair{{"File1", false}, {"File2", true}, {"File3", false}},
	}

	fileManager.Add(data)

	res := fileManager.GetListOfFiles()
	if len(res) != 1 && !reflect.DeepEqual(res[0], data) {
		t.Error("Failed : readed data is wrong. Expected : true got : false")
	}

	fileManager.Save()

	b, err := ioutil.ReadFile(gopath + "/src/github.com/wuzzapcom/TelegramTorrentBot/.TelegramTorrentBotInfo")
	if err != nil {
		t.Error("Failed to read .TelegramTorrentBotInfo\n", err)
	}

	if string(b) != dataString {
		t.Error("Readed .TelegramTorrentBotInfo contains wrong data . Expected : " + dataString + "got : " + string(b))
	}

	os.Remove(gopath + "/src/github.com/wuzzapcom/TelegramTorrentBot/.TelegramTorrentBotInfo")

}
