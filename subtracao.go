package main

type Subtracao struct{
    Operacao
}



func (s *Subtracao) Calcular(x, y float64) float64 {
   return x - y
}
