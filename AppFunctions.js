// Definir la función para obtener "hola" y actualizar el input
function obtenerHola() {
    fetch('/saludar')
        .then(response => response.text())
        .then(data => {
            document.getElementById('resultado').value = data;
        })
        .catch(error => console.error('Error:', error));
}

function enviarContenido() {
    const contenido = document.getElementById('inputTexto').value;
    fetch('/recibir', {
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
window.onload = function() {
    document.getElementById('saludarBtn').addEventListener('click', obtenerHola);
};
