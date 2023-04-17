package main

import "fmt"

func Questao1(compra float64, historico float64) float64  {
	var compra float64
	var historico float64
	var soma float64
	var desconto float64
	slice := []float64{}
	fmt.Print("Insira o valor da compra atual: ")
	fmt.Scan(&compra)
	if compra <= 0 {
		fmt.Println("Erro: Valor de compra inválido")
		return
	}
	for true {
		fmt.Print("Insira um valor do histórico do cliente(0 para parar): ")
		fmt.Scan(&historico)
		soma += historico
		if historico == 0 {
			break
		}
		slice = append(slice, historico)
		fmt.Print(slice)
	}
	media := soma / float64(len(slice))
	if media > 1000 {
		desconto = 0.2
	} else if len(slice) == 0 {
		desconto = 0.1
	} else if soma <= 500 {
		desconto = 0.02
	} else if soma <= 1000 {
		desconto = 0.05
	} else {
		desconto = 0.1
	}
	fmt.Println("Valor do desconto: ", compra*desconto)
	fmt.Println("Erro: nil")
}
