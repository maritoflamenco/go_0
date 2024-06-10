package main

import (
	"fmt"
	"github/fl4m3sec/go_0/variables"
)

func main() {
	estado, texto := variables.ConvierteTexto(207)
	fmt.Println(estado)
	fmt.Println(texto)
}
