package docker

import (
	"fmt"
	"os/exec"
)

func add_apt_keys(){
	out, err := exec.Command("sh","-c",`
	apt-get update
	`).Output()
	if err != nil{
		panic(err)
	}
	fmt.Print(out)
}
