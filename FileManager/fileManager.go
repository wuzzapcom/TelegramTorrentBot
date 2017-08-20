package FileManager

import (
	"io/ioutil"
	"wuzzapcom/TelegramTorrentBot/Constants"
	"log"
	"encoding/json"
	"os"
)

type FileManager struct{

	pathToData string
	loadedData DataArray

}

type DataArray struct{

	Data []Data

}

type Data struct{

	PathToSource string //may be file or folder
	SizeOfSource int64
	DownloadDate int64
	Name string

}

func InitFileManager(pathToData string) FileManager {

	data, err := ioutil.ReadFile(pathToData + "/" + Constants.INFO_FILE_NAME)
	if err != nil{
		log.Panic(err)
	}

	var dataArray DataArray

	json.Unmarshal(data, &dataArray)

	return FileManager{pathToData: pathToData, loadedData: dataArray}

}

func (fileManager *FileManager) GetListOfFiles() []Data{

	return fileManager.loadedData.Data

}

func (fileManager *FileManager) Save() {

	data, err := json.Marshal(fileManager.loadedData)
	if err != nil{
		log.Panic(err)
	}

	err = os.Remove(fileManager.pathToData + "/" + Constants.INFO_FILE_NAME)
	if err != nil {
		log.Println(err)
	}

	file, err := os.Open()
	if err != nil{
		log.Panic(err)
	}

	file.Write(data)
	file.Close()

}

func (dataArray *DataArray) Add(data Data){

	dataArray.Data = append(dataArray.Data, data)

}

func (fileManager *FileManager) Add(data Data){

	fileManager.loadedData.Add(data)

}

