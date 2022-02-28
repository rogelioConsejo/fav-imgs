package persistence

import (
	"encoding/json"
	"fav-imgs/gallery/image"
	. "fav-imgs/gallery/interfaces"
	"fmt"
	"io/ioutil"
	"os"
)

type reader struct {
	images map[string]Image
}

// GetImages TODO: Handle errors better + refactor
func (r reader) GetImages() map[string]Image {
	jsonData, err := os.OpenFile(jsonFile, os.O_RDONLY, 0660)
	defer jsonData.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	stored, err := ioutil.ReadAll(jsonData)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	imagesJson := map[string]image.Json{}
	err = json.Unmarshal(stored, &imagesJson)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	images := make(map[string]Image)
	for key, imageJson := range imagesJson {
		images[key] = image.NewImage(imageJson.Title, imageJson.Url)
	}

	return images
}

func GetPersistenceReader() Read {
	return reader{}
}

type adder struct {
}

// AddImage TODO: Handle errors better + refactor + check if ID exists
func (a adder) AddImage(newImage Image) (id string) {

	id = makeId()
	fileReader, err := os.OpenFile(jsonFile, os.O_RDONLY|os.O_CREATE, 0660)
	defer func(jsonData *os.File) {
		err := jsonData.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(fileReader)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	stored, err := ioutil.ReadAll(fileReader)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	var images map[string]image.Json
	if string(stored) != "" {
		images = map[string]image.Json{}
		err = json.Unmarshal(stored, &images)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
	} else {
		images = make(map[string]image.Json)
	}

	images[id] = image.Json{
		Title: newImage.GetTitle(),
		Url:   newImage.GetUrl(),
	}
	newData, err := json.Marshal(images)
	fmt.Printf("%s\n", string(newData))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	fileWriter, err := os.OpenFile(jsonFile, os.O_TRUNC|os.O_WRONLY, 0660)
	_, err = fileWriter.Write(newData)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return id
}

func makeId() string {
	return RandStringRunes(15)
}

func GetPersistenceAdder() Add {
	return adder{}
}

const jsonFile = "./data.json"
