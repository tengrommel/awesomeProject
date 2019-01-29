package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"time"
)

func main() {
	user := "username"
	pass := "password"
	remotehost := "remotehost:22"
	knownhost := []byte("remotehost		ecdsa-sha2-nistp256 AAE...jZHN")
	_, _, hostkey, _, _, _ := ssh.ParseKnownHosts(knownhost)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.FixedHostKey(hostkey),
		Timeout:         5 * time.Second,
	}
	conn, _ := ssh.Dial("tcp", remotehost, config)
	sess, _ := conn.NewSession()
	stdin, _ := sess.StdinPipe()
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr
	_ = sess.Shell()
	fmt.Fprintf(stdin, "%s\n%s\n%s\n", "pwd", "ls", "exit")
	_ = sess.Wait()
	sess.Close()
}
