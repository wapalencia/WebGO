// forms.go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var path = "prueba.txt"

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

	http.HandleFunc("/upload", upload)
	http.HandleFunc("/uploadfile", uploadFile)
	http.HandleFunc("/Editar", EditPage)
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/recibir", RecibirPage2)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":7777", nil))

}

func upload(w http.ResponseWriter, r *http.Request) {
	//...................................
	//Reading into struct type from a JSON file
	//...................................

	var content, err = ioutil.ReadFile("Respuesta.json")
	if err != nil {
		log.Fatal(err)
	}
	user2 := PageResult{}

	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Id:%d, Name:%s, Password:%s, LoggedAt:%v", user2.ResultSeed, user2.Palabra, user2.ResultArry, user2.ResultSeed)

	HomePageVars := PageResult{
		Palabra:      user2.Palabra,
		Ispalidrome:  user2.Ispalidrome,
		ArryRead:     user2.ArryRead,
		ResultArry:   user2.ResultArry,
		WordStr:      user2.WordStr,
		PhraseStr:    user2.PhraseStr,
		ResponseWord: user2.ResponseWord,
		ArrySeed:     user2.ArrySeed,
		ArryWeight:   user2.ArryWeight,
		ResultSeed:   user2.ResultSeed,
	}

	t, err := template.ParseFiles("./view/upload.html") //parse the html file homepage.html
	if err != nil {                                     // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template xecuting error: ", err) //log it
	}
}
func uploadFile(w http.ResponseWriter, r *http.Request) {
	//...................................
	//Reading into struct type from a JSON file
	//...................................

	var content, err = ioutil.ReadFile("Respuesta.json")
	if err != nil {
		log.Fatal(err)
	}
	user2 := PageResult{}

	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Id:%d, Name:%s, Password:%s, LoggedAt:%v", user2.ResultSeed, user2.Palabra, user2.ResultArry, user2.ResultSeed)

	HomePageVars := PageResult{
		Palabra:      user2.Palabra,
		Ispalidrome:  user2.Ispalidrome,
		ArryRead:     user2.ArryRead,
		ResultArry:   user2.ResultArry,
		WordStr:      user2.WordStr,
		PhraseStr:    user2.PhraseStr,
		ResponseWord: user2.ResponseWord,
		ArrySeed:     user2.ArrySeed,
		ArryWeight:   user2.ArryWeight,
		ResultSeed:   user2.ResultSeed,
	}

	t, err := template.ParseFiles("./view/uploadfile.html") //parse the html file homepage.html
	if err != nil {                                         // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template xecuting error: ", err) //log it
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./view/index.html") //parse the html file homepage.html
	if err != nil {                                    // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePage) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {              // if there is an error
		log.Print("template executing error: ", err) //log it
	}

}
func RecibirPage(w http.ResponseWriter, r *http.Request) {

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

	t, err := template.ParseFiles("./view/Editar.html") //parse the html file homepage.html
	if err != nil {                                     // if there is an error
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

	myArrs := strings.Split(string(r.FormValue("Arry")), "")

	mArray := ReadArry(myArrs)

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

	content, err := json.Marshal(HomePageVars)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("Respuesta.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	crearArchivo()
	escribeArchivo(content)

	t, err := template.ParseFiles("./view/Editar.html") //parse the html file homepage.html
	if err != nil {                                     // if there is an error
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

	t, err := template.ParseFiles("./view/Editar.html") //parse the html file homepage.html
	if err != nil {                                     // if there is an error
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
func escribeArchivo(content []byte) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea

	_, err = file.Write(content)
	if existeError(err) {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
	fmt.Println("Archivo actualizado existosamente.")
}
func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
func crearArchivo() {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("Archivo creado exitosamente", path)
}
