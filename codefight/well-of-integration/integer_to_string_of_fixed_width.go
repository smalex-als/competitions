package switchlights

import "fmt"

func integerToStringOfFixedWidth(number int, width int) string {
	res := "00000" + fmt.Sprintf("%d", number)
	return res[len(res)-width:]
}
