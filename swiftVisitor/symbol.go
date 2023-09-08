package swiftVisitor

import (
    "fmt"
    "os"
    "html/template"
)

type Symbol struct {
    Id      	string
	Tipo      	string
    TipoDato 	string
    Scope       string
    Line        int
    Column      int
}

type SymbolList struct {
    Symbols []*Symbol
}

func NewSymbolList() *SymbolList {
    return &SymbolList{}
}

func (e *SymbolList) AddSymbol(id , Tipo, TipoDato, scope string, line, column int) {
    sym := &Symbol{
        Id:      id,
        Tipo: Tipo,
		TipoDato: TipoDato,
        Scope:       scope,
        Line:        line,
        Column:      column,
    }
    e.Symbols = append(e.Symbols, sym)
}

func (e *SymbolList) PrintSymbols() {
    for _, err := range e.Symbols {
        fmt.Printf("Symbol %s: %s\nScope: %s\nLine: %d\nColumn: %d\n", err.Id, err.Tipo, err.Scope, err.Line, err.Column)
    }
}

func (s *SymbolList) GenerateStyledHTMLSymbolTable(filename string) error {
    // Abrir un archivo de salida para guardar el informe HTML
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    // Definir la plantilla HTML
    const htmlTemplate = `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Symbol Table</title>
        <!-- Enlace a Bootstrap de forma online -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
    </head>
    <body>
        <div class="container">
            <h1 class="mt-4">Symbol Table</h1>
            <table class="table table-bordered mt-4">
                <thead>
                    <tr>
                        <th>Identifier</th>
                        <th>Type</th>
                        <th>Data Type</th>
                        <th>Scope</th>
                        <th>Line</th>
                        <th>Column</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .Symbols}}
                    <tr>
                        <td>{{.Id}}</td>
                        <td>{{.Tipo}}</td>
                        <td>{{.TipoDato}}</td>
                        <td>{{.Scope}}</td>
                        <td>{{.Line}}</td>
                        <td>{{.Column}}</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
        <!-- Enlace a los scripts de Bootstrap y jQuery de forma online -->
        <script src="https://code.jquery.com/jquery-3.5.1.slim.min.js"></script>
        <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@1.16.0/dist/umd/popper.min.js"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js"></script>
    </body>
    </html>
    `

    // Crear una nueva plantilla HTML
    tmpl, err := template.New("symbolTable").Parse(htmlTemplate)
    if err != nil {
        return err
    }

    // Escribir los datos de SymbolList en la plantilla y guardarla en el archivo
    err = tmpl.Execute(file, s)
    if err != nil {
        return err
    }

    return nil
}

func (s *SymbolList) ResetSymbols() {
    // Reiniciar la lista de símbolos asignándole una nueva lista vacía
    s.Symbols = []*Symbol{}
}