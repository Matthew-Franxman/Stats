package main

import "fmt"

//since the files are both int the same package, they can share any functions or structs inbetween them
//here I will initialize the players
//note there are two ways to initialize structs, the way done the first two times and the way done the third time
//neither way is wrong, its just your preference :)
//LOL JK I COULDNT GET THE SECOND ONE TO WORK

func main() {
	//so now imma try to make this into a slice so that it can be a little more efficient
	var farve, jordan, harper Player

	//ite being a little shithole about it so please ignore the green lines it works I promise
	farve = &Football{"Bret Farve", "4", 58, 5343}
	jordan = &Basketball{"Michael Jordan", "23", 32.4, 4.7}
	harper = &Baseball{"Bryce Harper", "34", .343, 154}

	//here I make it into a slice list that can be added to
	s := []Player{farve, jordan, harper}

	//here we are going to run each of the returnStats() methods and they will each return them for different sports
	//I made it into a for loop to make it much less code then hardcoding every single one of these out
	for i := 0; i < len(s); i++ {
		fmt.Println()
		s[i].returnStats()
		fmt.Println()
	}

	//“There's a fine line between genius and insanity. I have erased this line.”
	//Oscar Levant

}
