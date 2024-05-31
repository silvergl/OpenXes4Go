package OpenXes4Go

import (
	"encoding/xml"
	"fmt"
	"github.com/openxes4go/model"
	"os"
)

func WriteFile(log model.XLog) error {
	bytes, err := xml.Marshal(log)
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return err
	}
	file, err2 := os.Create("output.xml")
	if err2 != nil {
		return err2
	}
	defer file.Close()

	_, err3 := file.Write(bytes)
	if err3 != nil {
		return err3
	}
	return nil

}
