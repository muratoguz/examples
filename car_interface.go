package main

import (
	"fmt"
	"strconv"
)

// Tüm otomobil struct'ları için ortak özellikleri barındıran interface.
type Carface interface {
	Run() bool
	Stop() bool
	Information() string
	Hebele() bool
}

// Carface interface'i için tasarladığımız ve ortak kullanılan metot.
func CarExecute(c Carface) {

	fmt.Println("\n")
	fmt.Println("Araç Bilgi; \n" + c.Information())
	fmt.Println("\n")

	msg := ""

	isRun := c.Run()
	if isRun {
		msg = "çalışıyor"
	} else {
		msg = "çalışmıyor"
	}
	fmt.Println("Araç " + msg + ".")

	isStop := c.Stop()
	if isStop {
		msg = "durdu"
	} else {
		msg = "durmuyor, fren tutmuyor"
	}
	fmt.Println("Araç " + msg + ".")

	isHebele := c.Hebele()
	if isHebele {
		fmt.Println("Hebele")
	}
}

// Struct'lar

// Base Struct
// Tüm otomobil struct'ları için ortak base struct'tır.
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

// Base Struct
// Bu "base struct'ı" sadece özel üretilen markaların(Ferrari, Lamborghini) struct'ları miras alabilir.
type SpecialProduction struct {
	Special bool
}

// Ferrari - Start

type Ferrari struct {
	Car // Struct'ı anonim olarak tanımlayarak kalıtım sağladık.
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : " + x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special : " + add
	}
	return ret
}

// Ferrari - End

// Lamborghini - Start

type Lamborghini struct {
	Car // Struct'ı anonim olarak tanımlayarak kalıtım sağladık.
	SpecialProduction
}

func (_ Lamborghini) Run() bool {
	return true
}

func (_ Lamborghini) Stop() bool {
	return true
}

func (x *Lamborghini) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : " + x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special : " + add
	}
	return ret
}

// Lamborghini - End

// Mercedes - Start

type Mercedes struct {
	Car // Struct'ı anonim olarak tanımlayarak kalıtım sağladık.
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return false
}

func (_ Mercedes) Hebele() bool {
	return true
}

func (x *Mercedes) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : " + x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	return ret
}

// Mercedes - End

func main() {

	// Ferrari
	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "F50"
	ferr.Color = "Red"
	ferr.Speed = 340
	ferr.Price = 1.4
	ferr.Special = true
	// fmt.Println(ferr.Information())

	// Özel üretimi bulunmayan marka örneği(Non SpecialProduction)
	merc := new(Mercedes)
	merc.Brand = "Mercedes"
	merc.Model = "CLX"
	merc.Color = "Black"
	merc.Speed = 290
	merc.Price = 0.4
	// fmt.Println(merc.Information())

	//CarExecute(ferr)
	CarExecute(merc)
}
