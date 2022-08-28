package exec

import (
	"os/exec";
	"fmt";
)


func ExecuteLinux(input string) error {
	cmd := exec.Command(input)
	stdOut, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(stdOut))
	return nil
	

}