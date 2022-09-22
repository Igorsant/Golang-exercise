package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	
	inputsNK, _ := reader.ReadString('\n')
	inputsNK = strings.Trim(inputsNK, " \n")

	inputList := strings.Split(inputsNK, " ")
	
	var studentsList []string
	N, _ := strconv.Atoi(inputList[0])
	K, _ := strconv.Atoi(inputList[1])

	for i := 0; i < N; i++ {
		var student string
		fmt.Scanln(&student)
		studentsList = append(studentsList, student)
	}
	
	

	for i := 0; i < len(studentsList); i++ {
		for j := 0; j < len(studentsList); j++ {
			if(int(studentsList[i][0]) < int(studentsList[j][0])){
				aux := studentsList[i]
				studentsList[i] = studentsList[j]
				studentsList[j] = aux
			}
		}
	}

	fmt.Println("Choosen student:", studentsList[K])
}