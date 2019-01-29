package main

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	user := "username"
	pass := []byte("password")
	remotehost := "remotehost:22"
	cmd := "hostname"
	privkeyfile := "/User/username/.ssh/id_rsa"
	knownhost := []byte("remotehost ecdsa-sha2-nistp256 AAE...jZHN")
	_, _, hostkey, _, _, _ := ssh.ParseKnownHosts(knownhost)
	privkeydata, _ := ioutil.ReadFile(privkeyfile)
	parsekey := ssh.ParsePrivateKeyWithPassphrase
	privkey, _ := parsekey(privkeydata, pass)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privkey),
		},
		HostKeyCallback: ssh.FixedHostKey(hostkey),
		Timeout:         5 * time.Second,
	}
	conn, _ := ssh.Dial("tcp", remotehost, config)
	sess, _ := conn.NewSession()
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr
	_ = sess.Run(cmd)
	sess.Close()
}
