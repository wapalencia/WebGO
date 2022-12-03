// forms.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"strings"
	"time"
	//"os"
)

type Tpalindrome struct {
	Palabra     string
	Ispalidrome string
}
type TsearchWord struct {
	Palabra   string
	Frasworde string
	ExisWord  bool
}

type TreadArry struct {
	Arry       []string
	ResultArry [][]string
}

type TcreateSeed struct {
	Arry       []string
	ResultArry []string
	ResultSeed string
}

type Person struct {
	UserName string
}

type PageVariables struct {
	Date     string
	Time     string
	UserName string
	Name     string
	Correo   string
	Message  [][]string
	Age      string
}

type PageResult struct {
	Palabra      string
	Ispalidrome  string
	WordStr      string
	PhraseStr    string
	ResponseWord string
	ArryRead     []string
	ResultArry   [][]string
	ArrySeed     []string
	ArryWeight   []string
	ResultSeed   int
}

func main() {

	http.HandleFunc("/Editar", EditPage)
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/recibir", RecibirPage2)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":7777", nil))

	/* 	tmpl := template.Must(template.ParseFiles("forms.html"))

	   	t := template.New("fieldname example")
	   	t, _ = t.Parse("welcome {{.UserName}}!")
	   	p := Person{UserName: "Wilfredo"}
	   	t.Execute(os.Stdout, p)

	   	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	   		if r.Method != http.MethodPost {
	   			tmpl.Execute(w, nil)

	   			return
	   		}
	   		details := ContactDetails{
	   			//name:    r.FormValue("email"),
	   			name:    r.FormValue("name"),
	   			correo:  r.FormValue("subject"),
	   			Message: r.FormValue("message"),
	   		}

	   		// do something with details
	   		_ = details

	   		tmpl.Execute(w, struct{ Success bool }{true})
	   		tmpl.Execute(w, p)

	   	})

	   	http.ListenAndServe(":8000", nil)
	   	fmt.Println("run server: http://localhost:8000/") */

}

func HomePage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("index.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePage) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {              // if there is an error
		log.Print("template executing error: ", err) //log it
	}

}
func RecibirPage(w http.ResponseWriter, r *http.Request) {
	//x := []byte("conact page.")
	//w.Write(x)

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println(r.Form["UserName"])
		fmt.Println(r.Form["Age"])
		fmt.Println(r.Form["mensage"])
	}

	var n string

	if IsPalindrome(string(r.FormValue("UserName"))) {
		n = "Verdadro"
	} else {
		n = "False"
	}
	//arr := ReadArry(MySplit(string(r.FormValue("mensage"))))
	myArrs := strings.Split(string(r.FormValue("mensage")), "")

	p := string(r.FormValue("name"))
	//n := IsPalindrome(string(r.FormValue("UserName")))> 0 ? 1 : 0
	c := string(r.FormValue("escribir"))
	m := ReadArry(myArrs)
	fmt.Println("my split ", myArrs)
	e := string(r.FormValue("Age"))

	now := time.Now()              // find the time right now
	HomePageVars := PageVariables{ //store the date and time in a struct
		Date:     now.Format("02-01-2006"),
		Time:     now.Format("15:04:05"),
		UserName: p,
		Name:     n,
		Correo:   c,
		Message:  m,
		Age:      e,
	}

	t, err := template.ParseFiles("Editar.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}

}

func RecibirPage2(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println(r.Form["Palindrome"])
		fmt.Println(r.Form["word"])
		fmt.Println(r.Form["phrase"])
		fmt.Println(r.Form["Arry"])
		fmt.Println(r.Form["seed"])
		fmt.Println(r.Form["weigth"])
	}

	var vPalindrome, exitWord string

	if IsPalindrome(string(r.FormValue("Palindrome"))) {
		vPalindrome = "Verdadro"
	} else {
		vPalindrome = "False"
	}
	////////////////////////////////////////// este el segundo ejercicio ///////////////////////////////////
	myWord := string(r.FormValue("word"))
	myPhrase := string(r.FormValue("phrase"))

	if ExistWord(myPhrase, myWord) {
		exitWord = "Verdadro"
	} else {
		exitWord = "False"
	}

	//arr := ReadArry(MySplit(string(r.FormValue("mensage"))))
	myArrs := strings.Split(string(r.FormValue("Arry")), "")

	//p := string(r.FormValue("weigth"))
	//n := IsPalindrome(string(r.FormValue("UserName")))> 0 ? 1 : 0
	//c := string(r.FormValue("seed"))
	mArray := ReadArry(myArrs)
	//fmt.Println("my split ", myArrs)
	//e := string(r.FormValue("word"))
	//now := time.Now()              // find the time right now

	//////////////////////////////////////// Cuarto ejercicio ///////////////////////////
	//var arrySeed, arryWeight []string

	arrySeed := strings.Split(string(r.FormValue("seed")), "")
	arryWeight := strings.Split(string(r.FormValue("weigth")), "")
	totSeed := CountSeed(arrySeed, arryWeight)

	HomePageVars := PageResult{
		Palabra:      string(r.FormValue("Palindrome")),
		Ispalidrome:  vPalindrome,
		ArryRead:     myArrs,
		ResultArry:   mArray,
		WordStr:      myWord,
		PhraseStr:    myPhrase,
		ResponseWord: exitWord,
		ArrySeed:     arrySeed,
		ArryWeight:   arryWeight,
		ResultSeed:   totSeed,
	}

	t, err := template.ParseFiles("Editar.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template xecuting error: ", err) //log it
	}

}

func EditPage(w http.ResponseWriter, r *http.Request) {

	p := string("wpalencia")
	n := string("Wilfredo")
	c := string("wpalencia@gmail.com")
	m := make([][]string, 0)
	e := string("1")

	now := time.Now()              // find the time right now
	HomePageVars := PageVariables{ //store the date and time in a struct
		Date:     now.Format("02-01-2006"),
		Time:     now.Format("15:04:05"),
		UserName: p,
		Name:     n,
		Correo:   c,
		Message:  m,
		Age:      e,
	}

	t, err := template.ParseFiles("Editar.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("password:", r.Form["mensage"])
	}
}

func MySplit(str string) []string {
	var res = make([]string, 0)

	for i := 0; i < len(str); i++ {
		res = append(res, string(str[i]))
	}

	return res
}

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

func ExistWord(str, sub string) bool {
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

	//fmt.Println("Arreglo resultado: ", arrResult)

	totalSum := 0
	totalMult := 1
	//for i := 0; i < cap(arrResult); i++ {
	for _, x := range arrResult {
		totalSum = totalSum + x //arrResult[i]
		if cap(arrResult) > 0 && x != 0 && totalSum != 0 {
			totalMult = x * totalMult
		}

	}
	//fmt.Println("total sum: ", totalSum)

	return totalSum + totalMult
}
func IsPalindrome(w string) bool {

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
