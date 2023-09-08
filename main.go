package main

import (
	"encoding/json"
	"fmt"
	parser "modulo/parser"
	"modulo/swiftVisitor"
	"net/http"
	"github.com/antlr4-go/antlr/v4"
)

func Compilar(salida string) string {
	inputStream := antlr.NewInputStream(salida)
	lexer := parser.NewSwiftLanLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSwiftLanParser(tokens)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	tree := p.Inicio()
	var visitor parser.SwiftLanVisitor = swiftVisitor.NewVisitorEvalue()
	visitor.Visit(tree)
	salidaF := swiftVisitor.OutData()
	errores:=swiftVisitor.OutErrors()
	simbolos:=swiftVisitor.OutSimbolos()
	simbolos.GenerateStyledHTMLSymbolTable("ReporteSimbolos.html")
	errores.GenerateStyledHTMLReport("ReporteErrores.html")
	return salidaF
}

// Area del Servidor
type Mensaje struct {
	Contenido string `json:"contenido"`
}

func compilarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var mensaje Mensaje
		err := json.NewDecoder(r.Body).Decode(&mensaje)
		if err != nil {
			http.Error(w, "Error al leer el cuerpo de la solicitud JSON", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		contenido := mensaje.Contenido
		sal := Compilar(contenido)
		w.Write([]byte(sal))

	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/compilar", compilarHandler)    // Nuevo manejador para recibir el contenido POST
	http.Handle("/", http.FileServer(http.Dir("."))) // Servir archivos estáticos desde la carpeta actual
	fmt.Println("Servidor Activo - http://localhost:3030/")
	fmt.Println("-|-|-|-|P1 OLC2|-|-|-|-")
	http.ListenAndServe(":3030", nil)
}