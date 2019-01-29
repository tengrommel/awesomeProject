package main

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"time"
)

func main() {
	user := "username"
	pass := "password"
	remotehost := "remotehost:22"
	knownhost := []byte("remotehost 	ecdsa-sha2-nisto256AAE...jZHN")
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
	defer conn.Close()
	client, _ := sftp.NewClient(conn)
	defer client.Close()
	_ = client.MkdirAll("./dir")
	dstfile, _ := client.Create("./dir/dest.txt")
	defer dstfile.Close()
	srcfile, _ := os.Open("./src.txt")
	_, _ = io.Copy(dstfile, srcfile)
}
