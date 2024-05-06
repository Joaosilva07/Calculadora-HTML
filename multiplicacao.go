package main

type Multiplicacao struct{
    Operacao
}

func (m *Multiplicacao) Calcular(x, y float64) float64 {
   return x * y
}