package jsonreader

import (
	"encoding/json"
	"io/ioutil"
	"os"
)



type Translations struct {
	Translations []Translation `json:"translations"`
}

type Translation struct {
	Linux string `json:"linux"`
	Windows string `json:"windows"`
}



// convert json to struct
func  ToStruct() (*Translations,error) {
	// read json file
	jsonFile, err := os.Open("./lib.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var translations Translations
	json.Unmarshal(byteValue, &translations)
	return &translations, nil


}







	


