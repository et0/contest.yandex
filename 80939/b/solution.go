package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		a, b, c    float64 // длина дороги от доме до магазина, дома до пункта выдачи, магазина до пункта выдачи
		v0, v1, v2 float64 // скорость без/одним/двумя грузом
	)
	fmt.Fscan(in, &a, &b, &c, &v0, &v1, &v2)

	strategy := [12]float64{
		// сразу покупаем/забираем дойдя до магазина или пункта выдачи
		a/v0 + c/v1 + b/v2, // д-ш-п-д
		b/v0 + c/v1 + a/v2, // д-п-ш-д

		a/v0 + a/v1 + b/v0 + b/v1, // д-ш-д-п-д
		b/v0 + b/v1 + a/v0 + a/v1, // д-п-д-ш-д

		b/v0 + c/v1 + c/v2 + b/v2, // д-п(з)-ш(к)-п-д
		a/v0 + c/v1 + c/v2 + a/v2, // д-ш(к)-п(з)-ш-д

		a/v0 + c/v0 + c/v1 + a/v2, // д-ш-п(з)-ш(к)-д
		b/v0 + c/v0 + c/v1 + b/v2, // д-п-ш(к)-п(з)-д

		a/v0 + c/v0 + b/v1 + a/v0 + a/v1, // д-ш-п(з)-д-ш(к)-д
		b/v0 + c/v0 + a/v1 + b/v0 + b/v1, // д-п-ш(к)-д-п(з)-д

		a/v0 + c/v0 + c/v1 + a/v1 + a/v0 + a/v1, // д-ш-п(з)-ш-д-ш(к)-д
		b/v0 + c/v0 + c/v1 + b/v1 + b/v0 + b/v1, // д-п-ш(к)-п-д-п(з)-д
	}

	minTime := strategy[0]
	for i := 1; i < len(strategy); i++ {
		minTime = min(minTime, strategy[i])
	}

	fmt.Fprintf(out, "%.15f\n", minTime)
}
