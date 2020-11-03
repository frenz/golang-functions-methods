/*
Let us assume the following formula for 
displacement s 
as a 

function of 

time t, 
acceleration a, 
initial velocity v0, 
initial displacement s0.

s = 1/2 * a * t^2 + v0 * t + s0

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.*/
package main

import (
	"os"
	"fmt"
	"math"
)

func PrintError(err error){
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func insertFloatParameter(userPrint string) float64{
	var input float64
	fmt.Print(userPrint)
	_, err := fmt.Scanf("%f", &input)
	PrintError(err)
	return input
} 

func GenDisplaceFn (a, v0, s0 float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return a/2 * math.Pow(t, 2) + v0 * t + s0
	}
	return fn
}
func main() {
	fmt.Println("assignement week2 - displacement.go")
	a := insertFloatParameter("Insert acceleration: ")
	v0 := insertFloatParameter("Insert initial velocity: ")
	s0 := insertFloatParameter("Insert initial displacement: ")
	t := insertFloatParameter("Insert displacement time: ")
	
	DisplaceFn := GenDisplaceFn(a, v0, s0)
	fmt.Println(DisplaceFn(t))
	os.Exit(0)
}
