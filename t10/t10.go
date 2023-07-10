package t10

import "fmt"

func temp_koleb(arr []float32) map[float32][]float32 {
	// groups
	// -40:-30
	// -30:-20
	// -20:-10
	// -10:0
	// 0:10
	// 10:20
	// 20:30
	// 30:40

	grad := []float32{-40, -30, -20, -10, 0, 10, 20, 30, 40}
	result := make(map[float32][]float32, len(grad))

	for _, elem := range arr {
		for i := 0; i < len(grad)-1; i++ {
			// если попадает в группу - пихаем в нее
			if grad[i] <= elem && elem <= grad[i+1] {
				result[grad[i]] = append(result[grad[i]], elem)
			}
		}
	}

	return result
}

func DoTask() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(temp_koleb(arr))
}
