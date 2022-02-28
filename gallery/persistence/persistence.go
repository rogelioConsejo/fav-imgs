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

func (r reader) GetImages() map[string]Image {
	images := makeImages(readFromFile())
	return images
}

func makeImages(imagesJson map[string]image.Json) map[string]Image {
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

func (a adder) AddImage(newImage Image) (id string) {
	images := readFromFile()

	id = makeId()
	_, exists := images[id]
	for exists {
		id = makeId()
		_, exists = images[id]
	}

	images[id] = makeJsonTranslatableStruct(newImage)
	writeToFile(images)

	return id
}

func makeJsonTranslatableStruct(newImage Image) image.Json {
	return image.Json{
		Title: newImage.GetTitle(),
		Url:   newImage.GetUrl(),
	}
}

//TODO: Handle errors
func readFromFile() map[string]image.Json {
	fileReader, _ := os.OpenFile(jsonFile, os.O_RDONLY|os.O_CREATE, 0660)
	defer fileReader.Close()

	stored, _ := ioutil.ReadAll(fileReader)

	var images map[string]image.Json
	if string(stored) != "" {
		images = map[string]image.Json{}
		_ = json.Unmarshal(stored, &images)
	} else {
		images = make(map[string]image.Json)
	}
	return images
}

//TODO: Handle errors
func writeToFile(images map[string]image.Json) {
	newData, _ := json.Marshal(images)
	fmt.Printf("%s\n", string(newData))
	fileWriter, _ := os.OpenFile(jsonFile, os.O_TRUNC|os.O_WRONLY, 0660)
	defer fileWriter.Close()
	_, _ = fileWriter.Write(newData)
}

func makeId() string {
	return RandStringRunes(15)
}

func GetPersistenceAdder() Add {
	return adder{}
}

const jsonFile = "./data.json"
