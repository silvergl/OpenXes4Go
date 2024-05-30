package OpenXes4Go

import (
	"encoding/xml"
	"fmt"
	"github.com/openxes4go/model"
	"os"
)

type FromXMLToModel struct {
	file os.File
}

func (f *FromXMLToModel) parseFile() model.XLog {
	x := model.XLog{}
	xlog := model.XLog{}
	return x
}

func writeFIle(log model.XLog) {
	xmlData, err := xml.MarshalIndent(log, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	xmlData = []byte(xml.Header + string(xmlData))
}
