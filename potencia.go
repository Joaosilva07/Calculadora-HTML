package main

import "math"

type Potencia struct{
     Operacao
}


func (p *Potencia) Calcular(x,y float64) float64 {
    return math.Pow(x, y)
}