package main

import ()

func main() {
	// Parse FLags
	username := flag.String("username", "", "check services with given username")
	flag.Parse()

	if *username == "" {
		// Read Username, if flags is empty
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\033[37;1mUsername:\033[0m ")
		*username, _ = reader.ReadString('\n')
	}

	*username = strings.ToLower(strings.Replace(strings.Trim(*username, " \r\n"), " ", "", -1))

	sherlock(*username)
}
