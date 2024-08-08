package health

import (
	"context"
	"fmt"
)

func Healthmeasure(ctx context.Context) error{
	var a int
	var b int
	var c int
	fmt.Print("Rate in the scale of 2 , how much u rate ur sleep cycle")
	fmt.Scan(&a)
	fmt.Println("Rate you physical activity in scale of 10")
	fmt.Scan(&b)
	fmt.Println("Rate ur workload life in scale of 10")
	fmt.Scan(&c)

	if a+b+c > 20 {
		fmt.Println("Good")
	} else if a+b+c < 20 && a+b+c > 10 {
		fmt.Println("ok")
	} else {
		fmt.Println("Bad")
	}
	return nil

}
