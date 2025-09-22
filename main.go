package main

import (
    "fmt"
    "time"

    "github.com/joaquimsnjunior/gopdf/service"
)

// Estruturas que vamos utilizar para inserir dados no PDF
type DataPosicao struct {
    Nome               string
    Quantidade         int
    PrecoMedio        string
    CotacaoAtual      string
    ValorInvestido    string
    Valor             string
    Resultado         string
    Rentabilidade     string
    ResultadoPositivo bool
}

type Invoice struct {
    // Dados do cliente
    NomeCliente string
    CPF         string
    Conta       string
    Assessor    string
    
    // Data da posição e geração
    DataPosicao string
    DataGeracao string
    
    // Items da carteira
    Items []DataPosicao
    
    // Resumo financeiro
    TotalInvestido        string
    ValorAtualCarteira    string
    ResultadoTotal        string
    RentabilidadeTotal    string
    ResultadoTotalPositivo bool
}

func main() {
    invoice := Invoice{
        // Dados do cliente
        NomeCliente: "Joaquim Silva",
        CPF:         "123.456.789-00",
        Conta:       "12345-6",
        Assessor:    "Maria Assessora",
        
        // Datas
        DataPosicao: "31/12/2024",
        DataGeracao: time.Now().Format("02/01/2006 15:04:05"),
        
        // Posições
        Items: []DataPosicao{
            {
                Nome:               "PETR4",
                Quantidade:         100,
                PrecoMedio:        "22,45",
                CotacaoAtual:      "27,30",
                ValorInvestido:    "2.245,00",
                Valor:             "2.730,00",
                Resultado:         "485,00",
                Rentabilidade:     "21,56",
                ResultadoPositivo: true,
            },
            {
                Nome:               "ITUB4",
                Quantidade:         130,
                PrecoMedio:        "18,90",
                CotacaoAtual:      "22,10",
                ValorInvestido:    "2.457,00",
                Valor:             "2.873,00",
                Resultado:         "416,00",
                Rentabilidade:     "16,92",
                ResultadoPositivo: true,
            },
            {
                Nome:               "CDB PREFIXADO",
                Quantidade:         1,
                PrecoMedio:        "1.275,60",
                CotacaoAtual:      "1.318,20",
                ValorInvestido:    "1.275,60",
                Valor:             "1.318,20",
                Resultado:         "42,60",
                Rentabilidade:     "3,34",
                ResultadoPositivo: true,
            },
        },
        
        // Resumo
        TotalInvestido:        "5.977,60",
        ValorAtualCarteira:    "6.921,20",
        ResultadoTotal:        "943,60",
        RentabilidadeTotal:    "15,79",
        ResultadoTotalPositivo: true,
    }

    // Instanciando PDF service
    wk := service.NewHtmlToPdf("tmp")

    // Carregando o HTML com nossa estrutura invoice
    err := service.LoadHtmlFile("index.html", invoice)
    if err != nil {
        fmt.Printf("Erro ao carregar HTML: %v\n", err)
        return
    }

    // Criando o PDF a partir do HTML que foi gerado na função acima
    filePDFName, err := wk.CreatePDF("index.html") // Mudei para index.html
    if err != nil {
        fmt.Printf("Erro ao criar PDF: %v\n", err)
        return
    }
    
    fmt.Println("PDF gerado com sucesso: ", filePDFName)
}