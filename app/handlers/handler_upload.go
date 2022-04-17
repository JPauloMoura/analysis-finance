package handlers

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/analysis-finance/internal/entitys"
)

type templateUpload struct {
	Csv      *entitys.Csv
	ErrorMsg string
}

// HandleUpload é responsalvel por realizar o upload de arquivos
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		// upload maximo de 10mb
		r.ParseMultipartForm(10 << 20)

		// retorna o primeiro arquivo com o nome especificado no formulario
		file, headers, err := r.FormFile("uploadFile")
		if err != nil {
			log.Printf("%v: %v", _errRetrievingFile, err)
			return
		}
		defer file.Close()

		tmp := &templateUpload{}

		if headers == nil {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		} else if headers.Header["Content-Type"][0] != "text/csv" {
			log.Printf("error:%v", ErrorMsgInvalidType)

			tmp.ErrorMsg = ErrorMsgs[ErrorMsgInvalidType]
			home.ExecuteTemplate(w, "Home", tmp)
			return
		}

		if err := readCsv(file); err != nil {
			log.Printf("%v:%v", _errReadingFile, err)
		}

		tmp.Csv = entitys.NewCsv(headers.Filename, headers.Size)
		tmp.Csv.ConvertSizeToMB()

		log.Printf("successfully upload: 'name':%v,'size':%v", tmp.Csv.Name, tmp.Csv.Size)
		home.ExecuteTemplate(w, "Home", tmp)
	}

}

func readCsv(file io.Reader) error {
	reader := bufio.NewReader(file)
	var lineCount int64

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if lineCount == 0 {
			lineCount++
			continue
		}
		lineCount++

		r := csv.NewReader(strings.NewReader(string(line)))
		r.Comma = ','
		records, err := r.ReadAll()
		if err != nil {
			return err
		}

		for _, row := range records {
			log.Println(row)
		}
	}
	return nil
}
