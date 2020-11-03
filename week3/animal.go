package main
import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

func PrintError(err error){
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)						    
	}
}

type animal struct{
	food, locomotion, noise string
}

func (a animal) Move() string {
		return a.locomotion
}

func (a animal) Speak() string {
		return a.noise
}

func (a animal) Eat() string {
		return a.food
}

func readRequest() (string,string){
	var inputString string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString = scanner.Text()
	err := scanner.Err()
	PrintError(err)
	array := strings.Split(strings.ToLower(inputString)," ")
	return array[0], array[1]
}

func main() {
	fmt.Println("assignement week3 - animal.go")					
	animals := make(map[string]animal)
	animals ["cow"] = animal{"grass","walk","moo"}
	animals	["bird"] = animal{"worms","fly","peep"}
	animals ["snake"] =  animal{"mice","slither","hsss"}
	fmt.Println(animals)

	for{
		fmt.Print("\n> ")
		name, action := readRequest()
		var response string
		switch action {
			case "move":
				response = animals[name].Move()
			case "eat":
				response = animals[name].Eat()
			case "speak":
				response = animals[name].Speak()
		}
		fmt.Print(response)

	}
}
