package persistence

import (
	"encoding/json"
	"fav-imgs/gallery/image"
	. "fav-imgs/gallery/interfaces"
	"fmt"
	"io/ioutil"
	"os"
)

type persistence struct {
	reader  reader
	adder   adder
	deleter deleter
	updater updater
}

func (p persistence) GetImages() map[string]Image {
	return p.reader.GetImages()
}

func (p persistence) AddImage(image Image) (id string) {
	return p.adder.AddImage(image)
}

func (p persistence) DeleteImage(id string) {
	p.deleter.DeleteImage(id)
}

func (p persistence) Update(id string, image Image) {
	p.updater.Update(id, image)
}

func GetPersistence() Persistence {
	return persistence{
		reader:  GetPersistenceReader(),
		adder:   GetPersistenceAdder(),
		deleter: GetPersistenceDeleter(),
		updater: GetPersistenceUpdater(),
	}
}

type reader struct {
	images map[string]Image
}

func (r reader) GetImages() map[string]Image {
	images := makeImages(readFromFile())
	return images
}

func GetPersistenceReader() reader {
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

func GetPersistenceAdder() adder {
	return adder{}
}

type deleter struct {
}

func (d deleter) DeleteImage(id string) {
	images := readFromFile()
	delete(images, id)
	writeToFile(images)
}

func GetPersistenceDeleter() deleter {
	return deleter{}
}

type updater struct {
}

func (u updater) Update(id string, image Image) {
	images := readFromFile()
	_, exists := images[id]
	if exists {
		images[id] = makeJsonTranslatableStruct(image)
	}
	writeToFile(images)
}

func GetPersistenceUpdater() updater {
	return updater{}
}

func makeImages(imagesJson map[string]image.Json) map[string]Image {
	images := make(map[string]Image)
	for key, imageJson := range imagesJson {
		images[key] = image.NewImage(imageJson.Title, imageJson.Url)
	}
	return images
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

const jsonFile = "./data.json"
