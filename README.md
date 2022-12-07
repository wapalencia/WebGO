# WebGO
Nombre : Wilfredo Palencia
Secion 7122

exercise 1
This function is called inside the main, which receives a "string" parameter, that is, some text which will be validated if it is empty or if it is a palindrome, the function returns a boolean if it validates through a for loop and two conditional yes.

func IsPalindrome(w string) bool {
	if w == "" {
		return false

	}
	i := 0
	for _, item := range w {
		i++
		rev := item
		ini := w[len(w)-i]
		if string(rev) != string(ini) {
			println("el for", string(rev), "vs cadena", string(ini))
			return false
		}
	}
	return true
}

exercise 3
This function receives an array structure and returns an array of arrays. The condition is that it saves each position as long as it is equal to the previous one and when it finds one that is different it saves it in an array of arrResultContiguous as long as it is greater than 1 positions.

func ReadArry(Str []string) [][]string {

	palabra := Str
	var arrItemContiguo = make([]string, 0)
	var arrResultContiguo = make([][]string, 0)

	var len = cap(Str)
	var lenArrItemContiguo = cap(arrItemContiguo)
	//var lenArrResultContiguo = cap(arrResultContiguo)
	//fmt.Println("cap de vector ... ", lenArrItemContiguo)

	//console.log(len)
	//fmt.Println("len de vector ... ", len)

	for i := 0; i < len; i++ {

		//fmt.Println("Inicio for:  iteracion", i+1, "letra: ", palabra[i], "array contiguo:", arrItemContiguo)
		if lenArrItemContiguo == 0 {
			//fmt.Println("arrItemContiguo.length == 0 ", arrItemContiguo)
			arrItemContiguo = append(arrItemContiguo, palabra[i])
			//fmt.Println("primera asignacion ", arrItemContiguo, "lenght: ", cap(arrItemContiguo))
			lenArrItemContiguo++
		} else {
			//fmt.Println("mas de uno valor item ", arrItemContiguo[0], "letra: ", palabra[i])
			if arrItemContiguo[0] == palabra[i] {
				//fmt.Println(" segunda condicion : arrItemContiguo[0] ==", arrItemContiguo[0], " palabra[i] ", palabra[i])
				arrItemContiguo = append(arrItemContiguo, palabra[i])
				//fmt.Println("segunda asignacion", arrItemContiguo)
				lenArrItemContiguo++
			} else {
				if lenArrItemContiguo > 1 {
					//fmt.Println("tercer if: arrItemContiguo.length > 1", arrItemContiguo, "result ", arrResultContiguo)

					arrResultContiguo = append(arrResultContiguo, arrItemContiguo)
					arrItemContiguo = nil
					arrItemContiguo = []string{palabra[i]}
					lenArrItemContiguo = 1

					//fmt.Println("asignacion tercer caso: ", palabra[i], "array contiguo: ", arrItemContiguo, "array result :", arrResultContiguo, cap(arrResultContiguo))
					//fmt.Println("resultado hasta ahora, item", arrItemContiguo, "result ", arrResultContiguo, cap(arrResultContiguo))
				} else {
					//fmt.Println("cuarto caso")
					arrItemContiguo = nil
					arrItemContiguo = []string{palabra[i]}
					lenArrItemContiguo = 1
				}
			}
		}
	}

	//fmt.Println("asignacion final: ", "array contiguo: ", arrItemContiguo, "array result :", arrResultContiguo, cap(arrResultContiguo))
	if lenArrItemContiguo > 1 {
		//fmt.Println("ultima asignacion: arrItemContiguo.length > 1", arrItemContiguo, "result ", arrResultContiguo)
		//arrResultContiguo[arrResultContiguo.length] = arrItemContiguo
		arrResultContiguo = append(arrResultContiguo, arrItemContiguo)

	}

	//console.log(arrResultContiguo)

	//fmt.Println("VALOR DE C2 SI EXISTE c2 ", c2)
	return arrResultContiguo
}

exercise 2
This function receives two string-type parameters, considering the phrase to be greater than the word and with two to register one for the phrase and another for the word, searching for the match between phrase and word. "returns true" if the word exists.

func ExistWord(str, sub string) bool {

	if str == "" || sub == "" {
		return false
	}

	if len(sub) > len(str) {
		return false
	}
	for i := 0; i < len(str)-len(sub)+1; i++ {

		//fmt.Println("primer for", string(str[i]))
		if str[i] != sub[0] {
			continue
		}
		exists := true
		for j := 1; j < len(sub) && exists; j++ {
			//fmt.Println("primer segundo for", string(sub[j]))
			if str[i+j] == sub[j] {
				continue
			}
			exists = false
		}
		if exists {
			return true
		}
	}
	return false
}
exercise 4
it's Receive two array and return an integer by rearising the following operation:

["a","b","c","a"] if that is the array and the weights is ["y","x","a","b"] the values resulting 2= the index of "a", 3=index of b, c, It has no weight since it is not in the list of weights, therefore the formula is: (2+3+2)+(2*3*2) = 19

func CountSeed(arrySeed, arrWeight []string) int {

	var arrResult = make([]int, 0)

	var lenSeed = cap(arrySeed)
	var lenVal = cap(arrWeight)

	for i := 0; i < lenSeed; i++ {
		for j := 1; j < lenVal; j++ {
			if arrySeed[i] == arrWeight[j] {
				fmt.Println("there is a seed, arrSee position: ", i, ", arrWeight Weight: ", j)
				//arrResult[arrResult.length] = j
				arrResult = append(arrResult, j)

			}
		}
	}

	

	totalSum := 0
	totalMult := 1
	
	for _, x := range arrResult {
		totalSum = totalSum + x //arrResult[i]
		if cap(arrResult) > 0 && x != 0 && totalSum != 0 {
			totalMult = x * totalMult
		}

	}
	//fmt.Println("total sum: ", totalSum)

	return totalSum + totalMult
}
/// By run the exercise it is necessary to open the project run the file:
"go to run app.go" in terminal.

*** Open an internet browser and go to the following URL "http://localhost:7777/"

*** Additional features allow adding a Json file, which has a POST service that can be consumed as an API.













