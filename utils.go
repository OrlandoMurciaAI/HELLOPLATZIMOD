package utils

import "fmt"

// recordar que la letra mayuscula es que la funcion
// esta disponible para otros modulos , letra miniscula
// solo disponible adentro del mismo modulo

func HelloWorld(name string) {
	fmt.Printf("Hello %s These are your first steps from Platzi in Go's World", name)
}
