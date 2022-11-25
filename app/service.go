package app

import "fmt"

type empData struct {
	A0  string
	A1  string
	A2  string
	A3  string
	A4  string
	A5  string
	A6  string
	A7  string
	A8  string
	A9  string
	A10 string
	A11 string
}

func Run(csvPath string) {
	fmt.Println(csvPath)
	linhas := LeitorCsv(csvPath)
	for _, line := range linhas {

		emp := empData{
			A0:  line[0],
			A1:  line[1],
			A2:  line[2],
			A3:  line[3],
			A4:  line[4],
			A5:  line[5],
			A6:  line[6],
			A7:  line[7],
			A8:  line[8],
			A9:  line[9],
			A10: line[10],
		}
		fmt.Printf("\n%s, %s, %s, %s", emp.A3, emp.A7, emp.A8, emp.A9)
	}
}
