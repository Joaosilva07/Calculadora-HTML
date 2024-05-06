package main


type Operacao struct {
    Operador string
    Num1, Num2 float64
}

type Calcular interface {
    Calcular(x,y float64) float64
}

