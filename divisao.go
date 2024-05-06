package main

import "fmt"
type Divisao struct{
     Operacao
}



func (d *Divisao) Calcular(x,y float64) float64 {
    if d.Num2 == 0 {
        fmt.Println("Erro: divisão por zero")
        return 0
    }
    return x / y
}