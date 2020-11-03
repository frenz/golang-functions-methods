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
type Animal interface {
	Eat()
	Move()
	Speak()			
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

func readRequest() (string,string,string){
	var inputString string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString = scanner.Text()
	err := scanner.Err()
	PrintError(err)
	words := strings.Split(strings.ToLower(inputString)," ")
	return words[0], words[1], words[2]
}


func main() {
	fmt.Println("assignement week3 - animal.go")					
	animals := make(map[string]animal)
	animals ["cow"] = animal{"grass","walk","moo"}
	animals	["bird"] = animal{"worms","fly","peep"}
	animals ["snake"] =  animal{"mice","slither","hsss"}
	for{
		fmt.Print("\n> ")	
		mode, name, action := readRequest()
		if mode == "query" {

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
		if mode == "newanimal" {
			animals[name] = animals[action]
		}
	}
}
