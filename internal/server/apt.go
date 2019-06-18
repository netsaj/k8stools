package ubuntu

import (
	"fmt"
	"os/exec"
)

func (e *Ubuntu) AptUpdate() {
	out, err := exec.Command("sh", "-c", `
	apt-get update
	`).Output()
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
}

func (e *Ubuntu) AptAddKey(key string) {
	out, err := exec.Command("sh", "-c", `
	apt-get update
	`).Output()
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
}
