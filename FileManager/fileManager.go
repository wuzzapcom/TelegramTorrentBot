package FileManager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"wuzzapcom/TelegramTorrentBot/Constants"
)

type FileManager struct {
	pathToData string
	loadedData DataArray
}

type DataArray struct {
	Data []Data
}

type Data struct {
	PathToSource string //may be file or folder
	SizeOfSource int64
	Name         string
	FileNames    []Pair
}

type Pair struct {
	S string
	B bool
}

func (data *Data) ToString() string {

	result := data.Name + "\n"

	result += "    Size : " + strconv.FormatFloat((float64(data.SizeOfSource/1024)/1024)/1024, 'f', 2, 64) + "GB \n"

	for _, name := range data.FileNames {

		var ending string

		if name.B {
			ending = " OK\n"
		} else {
			ending = "\n"
		}

		result += "    " + name.S + ending

	}

	return result

}

func InitFileManager(pathToData string) FileManager {

	data, err := ioutil.ReadFile(pathToData + Constants.INFO_FILE_NAME)
	if err != nil {
		log.Println(err)
		return FileManager{pathToData: pathToData, loadedData: DataArray{Data: make([]Data, 0, 5)}}
	}

	var dataArray DataArray

	json.Unmarshal(data, &dataArray)

	return FileManager{pathToData: pathToData, loadedData: dataArray}

}

func (fileManager *FileManager) GetListOfFiles() []Data {

	return fileManager.loadedData.Data

}

func (fileManager *FileManager) Save() {

	data, err := json.Marshal(fileManager.loadedData)
	if err != nil {
		log.Panic(err)
	}

	err = os.Remove(fileManager.pathToData + Constants.INFO_FILE_NAME)
	if err != nil {
		log.Println(err)
	}

	file, err := os.Create(fileManager.pathToData + Constants.INFO_FILE_NAME)
	if err != nil {
		log.Panic(err)
	}

	file.Write(data)
	file.Close()

}

func (dataArray *DataArray) Add(data Data) {

	dataArray.Data = append(dataArray.Data, data)

}

func (fileManager *FileManager) Add(data Data) {

	fileManager.loadedData.Add(data)

}
