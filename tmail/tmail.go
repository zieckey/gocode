package tmail

import (
	"os/exec"
	"log"
	"runtime"
)

// SendMail sends an email to the addresses using 'mail' command on *nux platform.
func SendMail(title, message string, email ...string) error {
	if runtime.GOOS == "windows" {
		log.Printf("TODO: cannot send email on windows title=[%v] messagebody=[%v]", title, message)
		return nil
	}
	mailCommand := exec.Command("mail", "-s", title)
	mailCommand.Args = append(mailCommand.Args, email...)
	stdin, err := mailCommand.StdinPipe()
	if err != nil {
		log.Printf("StdinPipe failed to perform: %s (Command: %s, Arguments: %s)", err, mailCommand.Path, mailCommand.Args)
		return err
	}
	stdin.Write([]byte(message))
	stdin.Close()
	_, err = mailCommand.Output()
	if err != nil || !mailCommand.ProcessState.Success() {
		log.Printf("send email ERROR : <%v> title=[%v] messagebody=[%v]", err.Error(), title, message)
		return err
	}

	return nil
}