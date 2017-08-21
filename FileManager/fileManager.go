package FileManager

import (
	"io/ioutil"
	"wuzzapcom/TelegramTorrentBot/Constants"
	"log"
	"encoding/json"
	"os"
	"strconv"
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
	Name string
	FileNames []string

}

func (data *Data) ToString() string{

	result := data.Name + "\n"

	result += "    Size : " + strconv.FormatInt(data.SizeOfSource, 10) + "\n"

	for _, name := range data.FileNames{

		result += "    " + name + "\n"

	}

	return result

}

func InitFileManager(pathToData string) FileManager {

	data, err := ioutil.ReadFile(pathToData + "/" + Constants.INFO_FILE_NAME)
	if err != nil{
		log.Println(err)
		return FileManager{pathToData:pathToData, loadedData: DataArray{Data:make([]Data, 0, 5)}}
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

	file, err := os.Create(fileManager.pathToData + "/" + Constants.INFO_FILE_NAME)
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

