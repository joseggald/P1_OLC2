var codeMirrorInstance

function compilarData() {
    const contenido = codeMirrorInstance.getValue();
    console.log(contenido)
    fetch('/compilar', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ contenido })
    })
    .then(response => response.text())
    .then(data => {
        console.log('Respuesta del servidor:', data);
    })
    .catch(error => console.error('Error:', error));
}


// Agregar el evento al botón cuando se cargue la página
window.onload = function () {

};


document.addEventListener("DOMContentLoaded", function () {
    // Obtener el contenedor y el textarea del editor de código
    var codeContainer = document.querySelector(".code-editor-container");
    var codeTextArea = document.getElementById("code-editor");
    
    // Configurar CodeMirror con el tema Dracula y la fuente monospace para el área de código
    codeMirrorInstance = CodeMirror.fromTextArea(codeTextArea, {
        lineNumbers: true,
        theme: "dracula",
        mode: "javascript",
        fontFamily: "Roboto Mono, monospace",
    });

    // Función para ajustar el tamaño del editor de código cuando se redimensiona
    function resizeCodeEditor() {
        codeMirrorInstance.setSize(
            codeContainer.clientWidth,
            codeContainer.clientHeight
        );
    }

    // Llamar a la función para ajustar el tamaño inicialmente y cuando se redimensione
    resizeCodeEditor();
    window.addEventListener("resize", resizeCodeEditor);

    // Función para ejecutar el código y mostrar el resultado en la consola
    function runCode() {
        try {
            var code = codeMirrorInstance.getValue();
            var result = eval(code);
            consoleTextArea.value = result;
        } catch (error) {
            consoleTextArea.value = "Error: " + error.message;
        }
    }

    // Ejecutar el código cuando se presione Enter en el área de código
    codeMirrorInstance.on("keypress", function (cm, event) {
        if (event.keyCode === 13 && !event.shiftKey) {
            runCode();
        }
    });

    // Ejecutar el código al hacer clic en un botón de ejecución (opcional)
    var runButton = document.getElementById("run-button");
    if (runButton) {
        runButton.addEventListener("click", runCode);
    }
});