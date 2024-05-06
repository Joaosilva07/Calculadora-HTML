package main

import ("math";"fmt")

type Radiciacao struct{
    Operacao
}

func (r *Radiciacao) Calcular(x,y float64) float64 {
   if r.Num2 == 0 {
       fmt.Println("Erro: radiciação por zero")
       return 0
   }
   return math.Pow(x , 1/y)
}