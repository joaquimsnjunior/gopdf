# 📄 GoPDF - Gerador de Relatórios PDF XP Investimentos

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Active-brightgreen.svg?style=for-the-badge)](README.md)

## 📋 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Funcionalidades](#funcionalidades)
- [Arquitetura](#arquitetura)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Como Usar](#como-usar)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Configuração](#configuração)
- [Exemplos](#exemplos)
- [API Reference](#api-reference)
- [Roadmap](#roadmap)
- [Contribuição](#contribuição)
- [Licença](#licença)

## 🎯 Sobre o Projeto

O **GoPDF** é uma aplicação Go desenvolvida para gerar relatórios PDF profissionais que simulam documentos oficiais. O sistema utiliza templates HTML customizáveis com dados dinâmicos para criar relatórios de posição consolidada de investimentos.

### 🌟 Principais Características

- **Design Profissional**: Interface idêntica aos relatórios oficiais da XP Investimentos
- **Templates Responsivos**: Layout adaptativo para diferentes dispositivos e impressão
- **Dados Dinâmicos**: Sistema de templates Go para inserção automática de dados
- **Performance Otimizada**: Geração rápida de PDFs com baixo uso de recursos
- **Fácil Customização**: Templates HTML/CSS facilmente modificáveis

## ⚡ Funcionalidades

### 📊 Relatórios Suportados

- [x] **Posição Consolidada de Investimentos**
  - Carteira de ações (PETR4, ITUB4, etc.)
  - Produtos de renda fixa (CDB, LCI, etc.)
  - Cálculo automático de rentabilidade
  - Resumo financeiro completo

- [x] **Informações do Cliente**
  - Dados pessoais (Nome, CPF)
  - Informações da conta
  - Assessor responsável
  - Data de posição

- [x] **Análise de Performance**
  - Valor investido vs. valor atual
  - Resultado absoluto (R$)
  - Rentabilidade percentual
  - Indicadores visuais (positivo/negativo)

### 🎨 Design Features

- **Header Profissional**: Logo XP dourado em fundo preto
- **Layout Responsivo**: Adaptável para desktop, tablet e mobile
- **Tabelas Interativas**: Design moderno com hover effects
- **Cores Corporativas**: Paleta oficial da XP Investimentos
- **Tipografia Otimizada**: Fonte Segoe UI para máxima legibilidade

## 🏗️ Arquitetura

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   main.go       │───▶│   service/       │───▶│   template.html │
│                 │    │                  │    │                 │
│ ┌─────────────┐ │    │ ┌──────────────┐ │    │ ┌─────────────┐ │
│ │ DataPosicao │ │    │ │ LoadHtmlFile │ │    │ │ HTML + CSS  │ │
│ │   Invoice   │ │    │ │ CreatePDF    │ │    │ │ Go Templates│ │
│ └─────────────┘ │    │ └──────────────┘ │    │ └─────────────┘ │
└─────────────────┘    └──────────────────┘    └─────────────────┘
        │                        │                        │
        ▼                        ▼                        ▼
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│  Dados de       │    │  Processamento   │    │   PDF Final     │
│  Investimentos  │    │  Template        │    │   Gerado        │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

## 🛠️ Tecnologias Utilizadas

| Tecnologia | Versão | Descrição |
|------------|--------|-----------|
| **Go** | 1.21+ | Linguagem principal do backend |
| **HTML5** | - | Estrutura dos templates |
| **CSS3** | - | Estilização e layout responsivo |
| **Go Templates** | - | Sistema de templates nativo |
| **wkhtmltopdf** | - | Engine de conversão HTML para PDF |

### 📦 Dependências

```go
// Principais dependências do projeto
github.com/joaquimsnjunior/gopdf/service  // Serviço de PDF customizado
```

## 📋 Pré-requisitos

Antes de começar, certifique-se de ter instalado:

- **Go 1.21 ou superior**
  ```bash
  go version
  ```

- **wkhtmltopdf** (para conversão PDF)
  ```bash
  # Ubuntu/Debian
  sudo apt-get install wkhtmltopdf
  
  # macOS
  brew install wkhtmltopdf
  
  # Windows
  # Baixar de: https://wkhtmltopdf.org/downloads.html
  ```

- **Git** (para clonagem do repositório)
  ```bash
  git --version
  ```

## 🚀 Instalação

### 1. Clone o repositório
```bash
git clone https://github.com/joaquimsnjunior/gopdf.git
cd gopdf
```

### 2. Inicialize o módulo Go
```bash
go mod init github.com/joaquimsnjunior/gopdf
go mod tidy
```

### 3. Instale as dependências
```bash
go get ./...
```

### 4. Crie a estrutura de diretórios
```bash
mkdir -p tmp
```

## 💻 Como Usar

### Execução Básica

```bash
# Executar o projeto
go run main.go
```

### Exemplo de Saída
```
PDF gerado com sucesso: file/ebe89729-8993-488e-86e1-3db6755beb62.pdf
```

### Personalização de Dados

Edite o arquivo `main.go` para customizar os dados:

```go
invoice := Invoice{
    NomeCliente: "Seu Nome",
    CPF:         "000.000.000-00",
    Conta:       "98765-4",
    Assessor:    "Seu Assessor",
    // ... mais configurações
}
```

## 📁 Estrutura do Projeto

```
gopdf/
├── 📄 README.md              # Documentação principal
├── 📄 main.go                # Arquivo principal da aplicação
├── 📄 index.html             # Template HTML principal
├── 📄 template.html          # Template alternativo
├── 📁 service/               # Pacote de serviços
│   ├── 📄 pdf.go            # Lógica de geração PDF
│   └── 📄 template.go       # Processamento de templates
├── 📁 tmp/                  # Diretório de arquivos temporários
│   └── 📄 *.pdf            # PDFs gerados
├── 📁 assets/              # Recursos estáticos (futuro)
│   ├── 📁 css/
│   ├── 📁 images/
│   └── 📁 fonts/
├── 📄 README.md

```

## ⚙️ Configuração

### Variáveis de Ambiente

Crie um arquivo `.env` (opcional):

```env
# Configurações do PDF
PDF_OUTPUT_DIR=tmp
PDF_QUALITY=high
PDF_FORMAT=A4

# Configurações de Debug
DEBUG_MODE=false
LOG_LEVEL=info
```

### Configuração do Service

```go
// Personalizar o serviço PDF
wk := service.NewHtmlToPdf("custom_output_dir")
```

## 📚 Exemplos

### Exemplo 1: Relatório Simples

```go
invoice := Invoice{
    NomeCliente: "João Silva",
    CPF:         "123.456.789-00",
    Items: []DataPosicao{
        {
            Nome:               "PETR4",
            Quantidade:         100,
            PrecoMedio:        "25,50",
            Valor:             "2.550,00",
            ResultadoPositivo: true,
        },
    },
}
```

### Exemplo 2: Múltiplos Ativos

```go
Items: []DataPosicao{
    {Nome: "PETR4", Quantidade: 100, /* ... */},
    {Nome: "ITUB4", Quantidade: 200, /* ... */},
    {Nome: "VALE3", Quantidade: 150, /* ... */},
}
```

## 📖 API Reference

### Estruturas Principais

#### `DataPosicao`
```go
type DataPosicao struct {
    Nome               string  // Código do ativo (ex: "PETR4")
    Quantidade         int     // Quantidade de cotas/ações
    PrecoMedio        string  // Preço médio de compra
    CotacaoAtual      string  // Cotação atual do ativo
    ValorInvestido    string  // Valor total investido
    Valor             string  // Valor atual da posição
    Resultado         string  // Resultado em reais
    Rentabilidade     string  // Rentabilidade percentual
    ResultadoPositivo bool    // Indica se o resultado é positivo
}
```

#### `Invoice`
```go
type Invoice struct {
    // Dados do cliente
    NomeCliente string
    CPF         string
    Conta       string
    Assessor    string
    
    // Controle temporal
    DataPosicao string
    DataGeracao string
    
    // Portfolio
    Items []DataPosicao
    
    // Resumo financeiro
    TotalInvestido         string
    ValorAtualCarteira     string
    ResultadoTotal         string
    RentabilidadeTotal     string
    ResultadoTotalPositivo bool
}
```

### Métodos do Service

```go
// Criar nova instância do serviço
wk := service.NewHtmlToPdf(outputDir string)

// Carregar template HTML com dados
err := service.LoadHtmlFile(templateFile string, data interface{})

// Gerar PDF
filename, err := wk.CreatePDF(htmlFile string)
```

## 🤝 Contribuição

Contribuições são sempre bem-vindas! Veja como você pode ajudar:

### Como Contribuir

1. **Fork** o projeto
2. **Clone** seu fork
   ```bash
   git clone https://github.com/seu-usuario/gopdf.git
   ```
3. **Crie** uma branch para sua feature
   ```bash
   git checkout -b feature/nova-funcionalidade
   ```
4. **Commit** suas mudanças
   ```bash
   git commit -m 'feat: adiciona nova funcionalidade'
   ```
5. **Push** para a branch
   ```bash
   git push origin feature/nova-funcionalidade
   ```
6. **Abra** um Pull Request

### Padrões de Código

- **Convenção de Commits**: Seguimos [Conventional Commits](https://www.conventionalcommits.org/)
- **Formatação**: Use `go fmt` antes de commitar
- **Testes**: Adicione testes para novas funcionalidades
- **Documentação**: Mantenha a documentação atualizada

### Tipos de Contribuição

| Tipo | Descrição | Label |
|------|-----------|-------|
| 🐛 **Bug Fix** | Correção de bugs | `bug` |
| ⚡ **Feature** | Nova funcionalidade | `enhancement` |
| 📚 **Docs** | Melhoria na documentação | `documentation` |
| 🎨 **Style** | Melhorias de design/CSS | `design` |
| 🔧 **Config** | Configurações e setup | `config` |

---

### 📊 Status do Projeto

![GitHub issues](https://img.shields.io/github/issues/joaquimsnjunior/gopdf)
![GitHub pull requests](https://img.shields.io/github/issues-pr/joaquimsnjunior/gopdf)
![GitHub last commit](https://img.shields.io/github/last-commit/joaquimsnjunior/gopdf)
![GitHub repo size](https://img.shields.io/github/repo-size/joaquimsnjunior/gopdf)

---

<div align="center">
  <p>Feito com ❤️ por <strong>Joaquim Silva Junior</strong></p>
  <p>⭐ Se este projeto te ajudou, considere dar uma estrela!</p>
</div>
