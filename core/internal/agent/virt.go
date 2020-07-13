package agent

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/zcalusic/sysinfo"
)

// CheckProduct check machine details
func CheckProduct() (virt string) {
	var si sysinfo.SysInfo
	si.GetSysInfo() // read sysinfo
	virt = fmt.Sprintf("%s (%s) by %s", si.Product.Name, si.Product.Version, si.Product.Vendor)
	return
}

// CheckContainer are we in a container? what container is it?
func CheckContainer() (product string) {
	product = "None"
	data, err := ioutil.ReadFile("/proc/1/cgroup")
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, "freezer") {
			log.Println("Checking if we are in a container...")
			fields := strings.Split(line, ":")
			if len(fields) > 1 &&
				fields[len(fields)-1] != "/" {
				product = strings.Split(fields[2], "/")[1]
				log.Println("Inside a container: ", product)
				return
			}
		}
	}
	log.Println("no, we are not")

	return
}
