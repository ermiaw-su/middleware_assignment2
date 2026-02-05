package main

import "fmt"

const quotas = 3
const waitingList = 2

type Participant struct {
	name string
}

type Waiting struct {
	name string
}

//Slice
var participant []Participant
var waiting []Waiting

func register(){
	//p and w become temporary variable
	var p Participant
	var w Waiting

	for (len(waiting) <= waitingList) {
		fmt.Println("Enter the name: ")
		
		if(len(participant) < quotas) {
			fmt.Scan(&p.name)
			fmt.Printf("%s%s\n", p.name, " is registered")
			participant = append(participant, p)
		} else if (len(waiting) < waitingList) {
			fmt.Scan(&w.name)
			fmt.Printf("%s%s\n", w.name, " in waiting list")
			waiting = append(waiting, w)
		} else {
			fmt.Scan(&w.name)
			fmt.Printf("%s%s\n", w.name, " rejected")
			break
		}
	}
}

func main() {
	register()
}