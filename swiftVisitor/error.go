package swiftVisitor

import (
    "fmt"
    "os"
    "html/template"
)

type Error struct {
    Number      int
    Description string
    Scope       string
    Line        int
    Column      int
}

type ErrorList struct {
    Errors []*Error
}

func NewErrorList() *ErrorList {
    return &ErrorList{}
}

func (e *ErrorList) AddError(number int, description, scope string, line, column int) {
    err := &Error{
        Number:      number,
        Description: description,
        Scope:       scope,
        Line:        line,
        Column:      column,
    }
    e.Errors = append(e.Errors, err)
}

func (e *ErrorList) ResetErrors() {
    // Reiniciar la lista de errores asignándole una nueva lista vacía
    e.Errors = []*Error{}
}

func (e *ErrorList) PrintErrors() {
    for _, err := range e.Errors {
        fmt.Printf("Error %d: %s\nScope: %s\nLine: %d\nColumn: %d\n", err.Number, err.Description, err.Scope, err.Line, err.Column)
    }
}
func (e *ErrorList) GenerateStyledHTMLReport(filename string) error {
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
        <title>Error Report</title>
        <!-- Enlace a Bootstrap de forma online -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
    </head>
    <body>
        <div class="container">
            <h1 class="mt-4">Error Report</h1>
            <table class="table table-bordered mt-4">
                <thead>
                    <tr>
                        <th>No Error</th>
                        <th>Descripción</th>
                        <th>Ambito</th>
                        <th>Linea</th>
                        <th>Columna</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .Errors}}
                    <tr>
                        <td>{{.Number}}</td>
                        <td>{{.Description}}</td>
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
    tmpl, err := template.New("report").Parse(htmlTemplate)
    if err != nil {
        return err
    }

    // Escribir los datos de ErrorList en la plantilla y guardarla en el archivo
    err = tmpl.Execute(file, e)
    if err != nil {
        return err
    }

    return nil
}