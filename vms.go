package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var dhcpgo ="vms.go"

type Server struct {
	Vmname string
	Description string 
}

type Vm struct {
	Id string
	IPAddress string
	Vmname string
	Description string
	Sku string
	Datacenter string
	Username string
}

func VmsHandler(c *gin.Context) {
	query := "SELECT vmname, ipaddress, description FROM vms WHERE username = $1"

	username, _ := c.Cookie("username")

	qrows, err := Db.Query(query, username)
	if err != nil {
		fmt.Println(err)
	}
	defer qrows.Close()

	var vmlist []Vm
	for qrows.Next() {
		var vm Vm
		if err := qrows.Scan(&vm.Vmname, &vm.IPAddress, &vm.Description); err != nil {
			fmt.Println(err)
			continue
		}
		// Append the shortcut to the slice corresponding to its category
		vmlist = append(vmlist, vm)
	}
	if err = qrows.Err(); err != nil {
		fmt.Println(err)
	}

	data := gin.H{
		"Title": "Vos machines virtuelles",
		"Vms": vmlist,
	}

	ShowPage(c, "vms", data)
}

func CreateVmsHandler(c *gin.Context) {

	if c.Request.Method == "POST" {
		fmt.Println("C'est POST")
		username, _ := c.Cookie("username")
		res := CreateServer(c.PostForm("vmname"), c.PostForm("description"), c.PostForm("password"))
		fmt.Println(res)
		SetVmInDb(c.PostForm("vmname"), c.PostForm("description"), c.PostForm("sku"), "192.168.1.201", username)
		SetResources(c.PostForm("sku"), c.PostForm("vmname"))	
	}

	data := gin.H{
		"Title": "Créer une machine virtuelle",
	}

	ShowPage(c, "create-vm", data)
}

func GetServer(server Server) {

}

func CreateServer(vmname, description, password string) string {
	// Workaround for certificate
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	patchString := fmt.Sprintf("https://192.168.1.85:10000/virtual-server/remote.cgi?program=create-domain&json=1&domain=%v&desc=%v&unix&dns&dir&mail&web&webmin&allocate-ip&generate-ssh-key&mysql&pass=%v", vmname, description, password)

	req, cerr := http.NewRequest("GET", patchString, nil)
	NonFatal(cerr, dhcpgo, "Preparing request to create vmname" + vmname)
	req.SetBasicAuth("root", os.Getenv("API_PWD"))

	client := &http.Client{}
	resp, err := client.Do(req)
	NonFatal(err, dhcpgo, "Executing GET request for site " + vmname)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	// Convert the body to a string or process it
	return string(body)
}

func SetResources(sku, vmname string) {
	// Workaround for certificate
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var cpu int
	var mem int

	switch sku {
	case "1":
		cpu = 1
		mem = 1073741824
	case "2":
		cpu = 2
		mem = 1073741824
	case "3":
		cpu = 2
		mem = 2147483648
	case "4":
		cpu = 4
		mem = 2147483648
	case "5":
		cpu = 4
		mem = 4294967296
	case "6":
		cpu = 6
		mem = 2147483648
	case "7":
		cpu = 6
		mem = 4294967296
	}

	patchString := fmt.Sprintf("https://192.168.1.85:10000/virtual-server/remote.cgi?program=modify-resources&json=1&domain=%v&max-procs=%v&max-mem=%v", vmname, cpu, mem)

	req, cerr := http.NewRequest("GET", patchString, nil)
	NonFatal(cerr, dhcpgo, "Preparing request to set resources" + sku)
	req.SetBasicAuth("root", os.Getenv("API_PWD"))

	client := &http.Client{}
	resp, err := client.Do(req)
	NonFatal(err, dhcpgo, "Executing GET request for SKU " + sku)
	defer resp.Body.Close()	
}

func SetVmInDb(name, description, sku, ipaddress, username string) {
	// Insert user into the database
	Db.Exec("INSERT INTO vms (vmname, description, ipaddress, sku, datacenter, username) VALUES (?, ?, ?, ?, ?, ?)", name, description, ipaddress, sku, "1", username)
}