package main

type Soma struct{
    Operacao
}


func (s *Soma) Calcular(x, y float64) float64 {
   return x + y
}