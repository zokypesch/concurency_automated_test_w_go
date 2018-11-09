package hello

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func Exec(c *gin.Context) {
	// exe_cmd("ls")
	// exe_cmd("workon env1")
	go exe_cmd("sh ./runCommand.sh")
	c.JSON(200, "---------------- running automation ------------------")
}

func exe_cmd(cmd string) ([]byte, error) {
	fmt.Println("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	// fmt.Printf("%s", out)
	// wg.Done() // Need to signal to waitgroup that this goroutine is done
	return out, err
}
