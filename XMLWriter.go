package OpenXes4Go

import (
	"context"
	"github.com/openxes4go/model"
	"os"
)

type FromXMLToModel struct {
	file os.File
}

type FromModelToXML struct {
	xlog model.XLog
}

func (f *FromXMLToModel) parseFile() model.XLog {
	x := model.XLog{}
	xlog:= model.XLog{
		
	}
	return x
}

func (f *FromModelToXML) writeFIle{

}
