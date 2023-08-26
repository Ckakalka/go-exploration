package main

import "fmt"

func main() {
	priceList := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200,
		"колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700,
		"буженина": 700, "помидоры": 250, "рыба": 300, "хамон": 1500}
	for k, v := range priceList {
		if v > 500 {
			fmt.Println(k)
		}
	}
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	orderCost := 0
	for _, k := range order {
		orderCost += priceList[k]
	}
	fmt.Printf("Стоимость заказа = %d\n", orderCost)
}
