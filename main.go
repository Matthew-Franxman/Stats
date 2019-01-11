package main

import "fmt"

//since the files are both int the same package, they can share any functions or structs inbetween them
//here I will initialize the players
//note there are two ways to initialize structs, the way done the first two times and the way done the third time
//neither way is wrong, its just your preference :)
//LOL JK I COULDNT GET THE SECOND ONE TO WORK

func main() {

	farve := Football{"Bret Farve", "4", 58, 5343}
	jordan := Basketball{"Michael Jordan", "23", 32.4, 4.7}
	harper := Baseball{"Bryce Harper", "34", .343, 154}

	//here we are going to run each of the returnStats() methods and they will each return them for different sports
	farve.returnStats()
	fmt.Println()
	fmt.Println()
	jordan.returnStats()
	fmt.Println()
	fmt.Println()
	harper.returnStats()

	//“There's a fine line between genius and insanity. I have erased this line.”
	//Oscar Levant

}
