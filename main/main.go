package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Machine struct {
	coffee int
	milk int
	sugar int
	water bool
	cup int
	cash1 int
	cash2 int
	cash5 int
	cash10 int
	cash50 int
	cash100 int
}

type Settings struct{
	colors bool
	clearScr bool
}

var settings Settings

type Coffee struct{
	name string
	price int
	milkUse int
	coffeeUse int
	sugarUse int
}
var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
//colorBlue := "\033[34m"
//colorPurple := "\033[35m"
//colorCyan := "\033[36m"
//var colorWhite = "\033[37m"

func saveAll(machine Machine){
	out := bytes.Buffer{}
	out.WriteString("coffee = " + strconv.Itoa(machine.coffee) + "\n")
	out.WriteString("milk = " +strconv.Itoa(machine.milk) + "\n")
	out.WriteString("sugar = " +strconv.Itoa(machine.sugar) + "\n")
	out.WriteString("water = " +strconv.FormatBool(machine.water) + "\n")
	out.WriteString("cup = " +strconv.Itoa(machine.cup) + "\n")
	out.WriteString("cash1 = " +strconv.Itoa(machine.cash1) + "\n")
	out.WriteString("cash2 = " +strconv.Itoa(machine.cash2) + "\n")
	out.WriteString("cash5 = " +strconv.Itoa(machine.cash5) + "\n")
	out.WriteString("cash10 = " +strconv.Itoa(machine.cash10) + "\n")
	out.WriteString("cash50 = " +strconv.Itoa(machine.cash50) + "\n")
	out.WriteString("cash100 = " +strconv.Itoa(machine.cash100) + "\n")

	f, err := os.Open("Machine")
	if err != nil {
		panic(err)
	}
	f, err = os.OpenFile("Machine", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	io.WriteString(f, out.String())
	io.WriteString(f, "              ")
	f.Close()
} 									//Запись ресурсов кофе-машины в файл
func getAll() (Machine){
	f, err := os.Open("Machine")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(c),"\n")
	coffee, err := strconv.Atoi(strings.Replace(arr[0],"coffee = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	milk, err := strconv.Atoi(strings.Replace(arr[1],"milk = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	sugar, err := strconv.Atoi(strings.Replace(arr[2],"sugar = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	water, err :=strconv.ParseBool(strings.Replace(arr[3],"water = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	cup, err := strconv.Atoi(strings.Replace(arr[4],"cup = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	cash1, err := strconv.Atoi(strings.Replace(arr[5],"cash1 = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	cash2, err := strconv.Atoi(strings.Replace(arr[6],"cash2 = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	cash5, err := strconv.Atoi(strings.Replace(arr[7],"cash5 = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	cash10, err := strconv.Atoi(strings.Replace(arr[8],"cash10 = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	cash50, err := strconv.Atoi(strings.Replace(arr[9],"cash50 = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	cash100, err := strconv.Atoi(strings.Replace(arr[10],"cash100 = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	var m Machine
	m = Machine{coffee,milk,sugar,water,cup,cash1,cash2,cash5,cash10,cash50,cash100}
	f.Close()
	return m
} 										//Чтение ресурсов кофе-машины из файла
func canChange(m Machine, bal int) (bool){
	var can = true
	if m.cash1+m.cash2*2+m.cash5*5+m.cash10*10 < bal {
		can = false
	}
	return can
} 						//может ли кофе машина выдать сдачу
func buyCoffee(m Machine,c Coffee, bal int) (Machine, int) {
	if m.water == true {
		if m.coffee - c.coffeeUse >= 0 {
			if m.milk - c.milkUse >= 0 {
				if m.sugar - c.sugarUse >= 0 {
					if m.cup > 0 {
						if bal >= c.price {
							m.coffee -= c.coffeeUse
							m.sugar -= c.sugarUse
							m.milk -= c.milkUse
							m.cup -= 1
							//temp := 0
							bal -= c.price
							clearScr()
							str := "Готовка [_______________]"
							for i := 0;i < 15;i++{
								time.Sleep(time.Millisecond*300)
								str = strings.Replace(str,"_","█",1)
								fmt.Printf("\r"+str)
							}
							addHistory(c)
							fmt.Printf("\r##### Ваш " + string(colorYellow)+c.name +string(colorReset)+ " готов #####")
						} else
						{ fmt.Println("Недостаточно денег") }
					} else
					{ fmt.Println("Закончились стаканчики") }
				} else
				{ fmt.Println("Недостаточно сахара") }
			} else
			{ fmt.Println("Недостаточно молока") }
		} else
		{ fmt.Println("Недостаточно кофе") }
	} else
	{ fmt.Println("Нет подлкючения к источнику воды") }
	saveAll(m)
	time.Sleep(time.Second)
	return m,bal
} 	//Покупка кофе
func getMenu() ([6]Coffee){
	var menu [6]Coffee
	menu[0] = Coffee{name: "Американо", price: 30, milkUse: 0, coffeeUse: 10, sugarUse: 10}
	menu[1] = Coffee{name: "Латте", price: 35, milkUse: 30, coffeeUse: 10, sugarUse: 20}
	menu[2] = Coffee{name: "Макиато", price: 40, milkUse: 10, coffeeUse: 15, sugarUse: 15}
	menu[3] = Coffee{name: "Эспрессо", price: 30, milkUse: 0, coffeeUse: 10, sugarUse: 10}
	menu[4] = Coffee{name: "Доппио", price: 40, milkUse: 0, coffeeUse: 20, sugarUse: 20}
	menu[5] = Coffee{name: "Латте макиато", price: 45, milkUse: 35, coffeeUse: 10, sugarUse: 25}
	return menu
} 									//Получение списка кофе, которые может приготовить кофе-машина
func getChange(bal int, m Machine) (int, Machine){
	for bal > 9 {
		bal -= 10
		m.cash10 -=1
	}
	for bal > 4 {
		bal -= 5
		m.cash5 -=1
	}
	for bal > 1 {
		bal -= 2
		m.cash2 -=1
	}
	for bal > 0 {
		bal -= 1
		m.cash1 -=1
	}
	return bal, m
} 				//Получить сдачу
func getNumber(x int) (string){
	if (settings.colors == false){
		return strconv.Itoa(x)
	}
	str := ""
	if x == 0{
		str = string(colorRed) + strconv.Itoa(x) + string(colorReset)
	} else {
		str = string(colorGreen) + strconv.Itoa(x) + string(colorReset)
	}

	return str
} 								//Цветные цифры
func getBool(x bool) (string){
	if (settings.colors == false){
		return strconv.FormatBool(x)
	}
	str := ""
	if x == false{
		str = string(colorRed) + strconv.FormatBool(x) + string(colorReset)
	} else {
		str = string(colorGreen) + strconv.FormatBool(x) + string(colorReset)
	}
	return str
}
func setSettings(){
	out := bytes.Buffer{}
	out.WriteString("colors = " + strconv.FormatBool(settings.colors) + "\n")
	out.WriteString("clearScr = " + strconv.FormatBool(settings.clearScr) + "\n")
	f, err := os.Open("settings")
	if err != nil {
		panic(err)
	}
	f, err = os.OpenFile("settings", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	io.WriteString(f, out.String())
	io.WriteString(f, "      ")
	f.Close()
}
func getSettings()  {
	f, err := os.Open("settings")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(c),"\n")
	colors, err :=strconv.ParseBool(strings.Replace(arr[0],"colors = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	clear, err :=strconv.ParseBool(strings.Replace(arr[1],"clearScr = ", "",1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	settings = Settings{colors: colors, clearScr: clear}
}
func clearScr() {
	if (settings.clearScr) {
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
	}
}
func addHistory(c Coffee){
	out := bytes.Buffer{}
	out.WriteString(strconv.Itoa(time.Now().Year()) + "-" + time.Now().Month().String() +"-"+ strconv.Itoa(time.Now().Day()) + " " + c.name + " " + strconv.Itoa(c.price) + "\n")
	f, err := os.OpenFile("history", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	io.WriteString(f, out.String())
	f.Close()
}
func getHistory(){
	f, err := os.Open("history")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(c),"\n")
	buyed := len(arr)-1
	cashAmount := 0
	for i:= 0;i < len(arr)-1;i++{
		str := strings.Fields(arr[i])
		temp, err := strconv.Atoi(str[2])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		cashAmount += temp
	}
	fmt.Println(string(c))
	fmt.Println("Всего куплено: " + getNumber(buyed) + " кофе на сумму " + getNumber(cashAmount) + "p\n" +
		"Нажмите Enter для продолжения...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main(){
	var machine Machine
	machine = getAll()
	getSettings()
	fmt.Println("##### Кофе машина RadCoffee v1.5 #####")
	time.Sleep(time.Second)
	bal := 0
	for start := true; start == true;{
		fmt.Println("##### Баланс: " + getNumber(bal))
		clearScr()
		fmt.Println("1) Купить кофе\n2) Тех. обслуживание\n3) Пополнить баланс\n4) Получить сдачу\n5) Настройки\n0) Выход\n#################################")
		var choice int
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1: // КУПИТЬ КОФЕ
			for buing := true;buing; {
				clearScr()
				fmt.Println("#############################\n" +
					"##### Баланс: " + getNumber(bal) +
					"\n##### Выберите вид кофе #####")
				var menu [6]Coffee
				menu = getMenu()
				for i := 0; i < 6; i++ {
					fmt.Println(strconv.Itoa(i+1) + ") " + menu[i].name + "  " + strconv.Itoa(menu[i].price) + "p")
				}
				fmt.Println("0) Выход в главное меню")
				fmt.Scanf("%d\n", &choice)
				switch choice {
				case 0:
					buing = false
					break
				case 1:
					machine, bal = buyCoffee(machine, menu[0], bal)
					break
				case 2:
					machine, bal = buyCoffee(machine, menu[1], bal)
					break
				case 3:
					machine, bal = buyCoffee(machine, menu[2], bal)
					break
				case 4:
					machine, bal = buyCoffee(machine, menu[3], bal)
					break
				case 5:
					machine, bal = buyCoffee(machine, menu[4], bal)
					break
				case 6:
					machine, bal = buyCoffee(machine, menu[5], bal)
					break
				}
				saveAll(machine)
			}
			break
		case 2: // ТЕХ ОБСЛУЖИВАНИЕ
			clearScr()
			fmt.Println("1) Статус\n" +
							"2) Забрать деньги\n" +
							"3) Пополнить запасы кофе\n" +
							"4) Пополнить запасы молока\n" +
							"5) Пополнить запасы сахара\n" +
							"6) Пополнить запасы стаканчиков\n" +
							"7) История покупок")
			fmt.Scanf("%d\n", &choice)
			switch choice {
			case 1:
				clearScr()
				fmt.Println("Кофе:", getNumber(machine.coffee)+"гр")
				fmt.Println("Сахар:", getNumber(machine.sugar)+"гр")
				fmt.Println("Молоко:", getNumber(machine.milk)+"мл")
				fmt.Println("Источник воды:", getBool(machine.water))
				fmt.Println("Стаканчики:", getNumber(machine.cup)+"шт")
				fmt.Println("1р:", getNumber(machine.cash1)+"шт")
				fmt.Println("2р:", getNumber(machine.cash2)+"шт")
				fmt.Println("5р:", getNumber(machine.cash5)+"шт")
				fmt.Println("10р:", getNumber(machine.cash10)+"шт")
				fmt.Println("50р:", getNumber(machine.cash50)+"шт")
				fmt.Println("100р:", getNumber(machine.cash100)+"шт")
				fmt.Println("Всего:", getNumber(machine.cash1+machine.cash2*2+machine.cash5*5+machine.cash10*10+machine.cash50*50+machine.cash100*100)+"р")
				fmt.Println("#################################\nНажмите Enter для продолжения...")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				break
			case 2:
				clearScr()
				fmt.Println("1) Забрать монеты по 1р\n" +
								"2) Забрать монеты по 2р\n" +
								"3) Забрать монеты по 5р\n" +
								"4) Забрать монеты по 10р\n" +
								"5) Забрать купюры по 50р\n" +
								"6) Забрать купюры по 100р")
				fmt.Scanf("%d\n", &choice)
				fmt.Print("Кол-во: ")
				amount := 0
				fmt.Scanf("%d\n", &amount)
				switch choice {
				case 1:
					if amount <= machine.cash1 {
						machine.cash1 -= amount
					} else {
						fmt.Println("В машине нет столько монет")
						time.Sleep(time.Second)
					}
					break
				case 2:
					if amount <= machine.cash2 {
						machine.cash2 -= amount
					} else {
						fmt.Println("В машине нет столько монет")
						time.Sleep(time.Second)
					}
					break
				case 3:
					if amount <= machine.cash5 {
						machine.cash5 -= amount
					} else {
						fmt.Println("В машине нет столько монет")
						time.Sleep(time.Second)
					}
					break
				case 4:
					if amount <= machine.cash10 {
						machine.cash10 -= amount
					} else {
						fmt.Println("В машине нет столько монет")
						time.Sleep(time.Second)
					}
					break
				case 5:
					amount := 0
					fmt.Scanf("%d\n", &amount)
					if amount <= machine.cash50 {
						machine.cash50 -= amount
					} else {
						fmt.Println("В машине нет столько купюр")
						time.Sleep(time.Second)
					}
					break
				case 6:
					amount := 0
					fmt.Scanf("%d\n", &amount)
					if amount <= machine.cash100 {
						machine.cash100 -= amount
					} else {
						fmt.Println("В машине нет столько купюр")
						time.Sleep(time.Second)
					}
					break
				}
				break
			case 3:
				machine.coffee += 1000
				if machine.coffee > 2000{
					machine.coffee = 2000
				}
				break
			case 4:
				machine.milk += 500
				if machine.milk > 1000{
					machine.milk = 1000
				}
				break
			case 5:
				machine.sugar += 1000
				if machine.sugar > 5000{
					machine.sugar = 2000
				}
				break
			case 6:
				machine.cup += 50
				if machine.cup > 100{
					machine.cup = 100
				}
			case 7:
				getHistory()
			}
			break

		case 3: // ПОПОЛНЕНИЕ БАЛАНСА
			for addbal := true; addbal;{
				clearScr()
				fmt.Println("##### Пополнение баланса ######\n" +
					"##### Баланс: " + getNumber(bal) +
					"\n1) 1р\n2) 2p\n3) 5p\n4) 10p\n5) 50p\n6) 100p\n0) Назад")
				fmt.Scanf("%d\n", &choice)
				switch choice {
				case 0:
					addbal = false
					break
				case 1:
					machine.cash1 += 1
					bal += 1
					break
				case 2:
					machine.cash2 += 1
					bal += 2
					break
				case 3:
					machine.cash5 += 1
					bal += 5
					break
				case 4:
					if canChange(machine, bal) {
						machine.cash10 += 1
						bal += 10
					} else {
						fmt.Println("В аппарате недостаточно монет для сдачи")
						time.Sleep(time.Second)
					}
					break
				case 5:
					if canChange(machine, bal) {
						machine.cash50 += 1
						bal += 50
					} else {
						fmt.Println("В аппарате недостаточно монет для сдачи")
						time.Sleep(time.Second)
					}
					break
				case 6:
					if canChange(machine, bal) {
						machine.cash100 += 1
						bal += 100
					} else {
						fmt.Println("В аппарате недостаточно монет для сдачи")
						time.Sleep(time.Second)
					}
					break
				}
			}
		case 4:
			if bal > 0 {
				fmt.Println("##### Ваша сдача: " + getNumber(bal) + "p #####")
				bal, machine = getChange(bal, machine)
				time.Sleep(time.Second)
			}
		case 5:
			for choice != 0 {
				clearScr()
				fmt.Println("##### Настройки #####\n1) Цвета текста: " + getBool(settings.colors) +
					"\n2) Очистка экрана: " + getBool(settings.clearScr) + "\n0) Назад")
				fmt.Scanf("%d\n", &choice)
				switch choice {
				case 1:
					settings.colors = !settings.colors
					setSettings()
					break
				case 2:
					settings.clearScr = !settings.clearScr
					setSettings()
					break
				case 0:
					choice = 0
				}
			}
		case 0:
			start = false
			break
		}
	}
	saveAll(machine)
}