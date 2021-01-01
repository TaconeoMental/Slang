fn iter(string, func) {
    i = 0
    for i < len(string) {
        func(string[i])
    }
}

// Crear operador para pasar argumentos por referencia a las funciones
fn print_item(item) {
    print item
}

fn main {
    string = "Hola qué tal"
    iter(string, print_item)
}

main()
