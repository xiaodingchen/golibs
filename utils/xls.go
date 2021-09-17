package utils

import (
	"bytes"
	"encoding/csv"
	"os"

	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
)

type XlsFile struct {
	Path     string
	FileType string
	Binary   []byte
	Charset  string `default:"utf-8"` // "utf-8"
}

func (x *XlsFile) Parse(index int) (rows [][]string, err error) {
	if x.FileType == "xls" {
		return x.ParseXls(index)
	}

	if x.FileType == "xlsx" {
		return x.ParseXlsx(index)
	}

	if x.FileType == "csv" {
		return x.ParseCsv(index)
	}

	return
}

// 解析xls文件
func (x *XlsFile) ParseXls(index int) ([][]string, error) {
	rows := make([][]string, 0)
	var f *xls.WorkBook
	var err error
	if len(x.Path) > 0 {
		f, err = xls.Open(x.Path, x.Charset)
	} else if len(x.Binary) > 0 {
		reader := bytes.NewReader(x.Binary)
		f, err = xls.OpenReader(reader, x.Charset)
	}

	if err != nil {
		return rows, err
	}

	if f == nil {
		return rows, err
	}
	sheet := f.GetSheet(0)
	rowCount := int(sheet.MaxRow + 1)
	sf := func(sheet *xls.WorkSheet, i int) *xls.Row {
		defer func() {
			if err := recover(); err != nil {
			}
		}()
		return sheet.Row(i)
	}

	for i := 0; i < rowCount; i++ {
		row := sf(sheet, i)
		if row == nil {
			continue
		}
		count := row.LastCol()
		trow := make([]string, 0)
		for c := 0; c < count; c++ {
			trow = append(trow, RemoveCsvBoom(row.Col(c)))
		}
		rows = append(rows, trow)
	}
	if index < len(rows) {
		return rows[index:], nil
	}
	return [][]string{}, nil
}

func (x *XlsFile) ParseXlsx(index int) ([][]string, error) {
	rows := make([][]string, 0)
	var err error
	var xlFile *xlsx.File
	if len(x.Path) > 0 {
		xlFile, err = xlsx.OpenFile(x.Path)
	} else if len(x.Binary) > 0 {
		xlFile, err = xlsx.OpenBinary(x.Binary)
	}

	if err != nil {
		return rows, err
	}
	if xlFile == nil {
		return rows, nil
	}
	sheet := xlFile.Sheets[0]
	rowCount := len(sheet.Rows)
	if rowCount == 0 {
		return rows, nil
	}

	for _, v := range sheet.Rows {
		if v == nil {
			continue
		}
		trow := make([]string, 0)
		for _, c := range v.Cells {
			trow = append(trow, RemoveCsvBoom(c.String()))
		}
		rows = append(rows, trow)
	}
	if index < len(rows) {
		return rows[index:], nil
	}
	return [][]string{}, nil
}

func (x *XlsFile) ParseCsv(index int) ([][]string, error) {
	rows := make([][]string, 0)
	var err error
	var reader *csv.Reader
	if len(x.Path) > 0 {
		handler, err := os.Open(x.Path)
		if err != nil {
			return rows, err
		}
		reader = csv.NewReader(handler)
	} else if len(x.Binary) > 0 {
		reader = csv.NewReader(bytes.NewReader(x.Binary))
	}

	if reader == nil {
		return rows, err
	}
	allRows, err := reader.ReadAll()
	if err != nil {
		return rows, err
	}

	for _, v := range allRows {
		row := make([]string, 0)
		for _, vv := range v {
			row = append(row, RemoveCsvBoom(vv))
		}
		rows = append(rows, row)
	}
	if index < len(rows) {
		return rows[index:], nil
	}
	return rows, err
}

func RemoveCsvBoom(str string) string {
	b := bytes.ReplaceAll([]byte(str), []byte{0xEF, 0xBB, 0xBF}, []byte(""))
	return bytes.NewBuffer(b).String()
}
