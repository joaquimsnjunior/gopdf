package service

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)


func LoadHtmlFile(htmlFileDir string, dataToLoad any) error {
	// Carrega o arquivo HTML e injeta os dados na estrutura fornecida
	content, err := ioutil.ReadFile(htmlFileDir)
	var htmlLoader = template.Must(template.New("htmlLoader").Parse(string(content)))
	if err != nil {
		log.Fatal("Erro ao ler o arquivo HTML: ", err)
	}
	// Cria um novo arquivo HTML temporário para armazenar o conteúdo processado
	fout, err := os.Create("template.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	// Executa o template com os dados fornecidos e escreve no arquivo temporário
	if err := htmlLoader.Execute(fout, dataToLoad); err != nil {
		log.Fatal("Erro ao executar o template: ", err)
		return err
	}
	return nil
}