// test

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains(`Hello, 
	playground sdfkj nsdjb kcnsdjk
	vnjksdn dmv ksd vdvdsvn3984u32yrewbyr7t
	748yt7443 n&^%$%*`, "&^%"))

	pkey := `-----BEGIN OPENSSH PRIVATE KEY-----
	b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	QyNTUxOQAAACC7nwuBhmo+H6VVNZSHI0in2iCU/yi67WfeFPfuyAdkBAAAAJh7nk7je55O
	4wAAAAtzc2gtZWQyNTUxOQAAACC7gIlwBMp+H6VVNZSHI0in2iCU/yi67WfeFPfuyAdkBA
	AAAEBHl+qBAosBAUIGuvdDR8gJN/PEhempLe4NtyKiO7hCPLuAiXAEyn4fpVU1lIcjSKfa
	IJT/KLrtZ94U9+7IB2QEAAAAEm1pa2VATWlrZS1zYW1zdW5nMQECAw==
	-----END OPENSSH PRIVATE KEY-----
	`

	fmt.Println(pkey)

}
