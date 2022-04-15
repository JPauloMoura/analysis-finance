package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type File struct {
	Name string
	Size int
}

// HandleUpload é responsalvel por realizar o upload de arquivos
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		// upload maximo de 10mb
		r.ParseMultipartForm(10 << 20)

		// retorna o primeiro arquivo com o nome especificado no formulario
		file, headers, err := r.FormFile("uploadFile")
		if headers == nil {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		if err != nil {
			log.Printf("%v: %v",
				_errorRetrievingFile,
				err,
			)
			return
		}
		defer file.Close()

		// cria um arquivo temporário no diretorio uploads
		tempFile, err := ioutil.TempFile("/tmp", "upload-*.csv")
		if err != nil {
			log.Printf("%v: %v, name: %s",
				_errorCreateingTempFile,
				err,
				headers.Filename,
			)
		}
		defer tempFile.Close()

		// faz a leitura dos bytes do arquivo de upload
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("%v: %v, name:%s",
				_errorReadingFile,
				err,
				headers.Filename,
			)
		}

		// escreve no arquivo temporário
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!
		log.Printf("successfully: upload, name:%s",
			headers.Filename,
		)

		infoFile := &File{
			Name: headers.Filename,
			Size: toMegabytes(int(headers.Size)),
		}

		log.Printf("fileName: %+v\n", infoFile.Name)
		log.Printf("fileSize: %+v\n", infoFile.Size)

		home.ExecuteTemplate(w, "Home", infoFile)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func toMegabytes(size int) int {
	return size / 1024
}
