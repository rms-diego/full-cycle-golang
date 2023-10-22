package main

const a = "hello, world!" // sintaxe para declara constantes

var (
	b bool    = true                        // Se não inicializado o tipo é inferido como false
	c int     = 10                          // Se não inicializado o tipo é inferido como 0
	d string  = "this variable is a string" // Se não inicializado o tipo é inferido como uma string vazia
	e float64 = 20.2                        // Se não inicializado o tipo é inferido como 0
)

func main() {
	b = true
	f := "short variable declaration" // Declaração curta, desse modo inferi o tipo na inicialização da variável

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
