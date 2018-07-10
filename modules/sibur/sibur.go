package sibur

import "github.com/tealeg/xlsx"

type API interface {
	GetTKP(string) (interface{}, error)
}

type sibur struct{}

func New() API {
	return &sibur{}
}

func (s *sibur) GetTKP(inPath string) (interface{}, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "I am a cell!"
	err = file.Save("./output.xlsx")
	if err != nil {
		return nil, err
	}
	return "MyXLSXFile.xlsx", nil
}
