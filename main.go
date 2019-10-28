package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	userName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAddress(p *Person, i int) string {
	return p.Email[i].Address
}

func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

//func GetPersonInterest(p *Person) []

func Writemessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("\n********************************\n")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}
}

func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func main() {
	person := Person{
		ID:        9,
		FirstName: "Murat",
		LastName:  "OĞUZ",
		userName:  "murat.oguz",
		Gender:    "m",
		Name:      Name{Family: "smdebug", Personal: "Murat"},
		Email: []Email{
			Email{ID: 1, Kind: "Personal", Address: "murat.oguz@smdebug.com"},
			Email{ID: 2, Kind: "Job", Address: "murat.oguz@metglobal.com"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "Go"},
			Interest{ID: 2, Name: "Python"},
			Interest{ID: 3, Name: "PHP"},
		},
	}

	Writemessage("Started...")
	WriteStarLine()
	res := GetPerson(&person)
	Writemessage(res)
	WriteStarLine()
	Writemessage("Emails:")
	resEmails := GetPersonEmailAddress(&person, 0)
	Writemessage(resEmails)
	WriteStarLine()
	Writemessage("Personal Email:")
	resEmail := GetPersonEmail(&person, 1)
	fmt.Println(resEmail)
	WriteStarLine()
	SaveJSON("person.JSON", person)
	Writemessage("End...")
}
