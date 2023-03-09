package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:  0,
	}
	return b
}

// format the bill using receiver function
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.item {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}
	//tips
	fs += fmt.Sprintf("%-25v ... $%v \n", "tip:", b.tip)
	//total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total +b.tip)

	return fs
}

//receiver fuction with pointers
//update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
} 
//add an item to the bill
func (b bill) addItem(name string, price float64){
	b.item[name] = price
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created bill - ", b.name)

	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item: ", reader)
		price, _ := getInput("Price: ", reader)

		p, err := strconv.ParseFloat(price, 64) //convert to float 64
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added - ", name, price)
		promptOptions(b)
		break
	case "s":
		b.save()
		fmt.Println("You save the file", b.name)
		break
	case "t":
		tip, _ := getInput("Tip amount ($)", reader)
		t, err := strconv.ParseFloat(tip, 64) //convert to float 64
		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip Added - ", t)
		promptOptions(b)
		break
	default:
		fmt.Println("Invalid option selected...")
		promptOptions(b)
	}
	// fmt.Println(opt)
}

//save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")

}