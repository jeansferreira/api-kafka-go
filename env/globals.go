package env

import (
	"fmt"
	"os"
	"strings"
)

var KAFKA_BROKERS = []string{}
var PLAT_USER = os.Getenv("PLAT_USER")         //"jean.ferreira"
var PLAT_PASSWORD = os.Getenv("PLAT_PASSWORD") //"3qxBJMEZ9dI8JRV54HGubdzcV9omab6b"
var CLIENTS_URL = os.Getenv("CLIENTS_URL")     //"clients.k8s.platform.linximpulse.net"

func SetupEnv() {
	KAFKA_BROKERS = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	// var PLAT_USER = os.Getenv("PLAT_USER")
	// var PLAT_PASSWORD = os.Getenv("PLAT_PASSWORD")
	// var CLIENTS_URL = os.Getenv("CLIENTS_URL")

	fmt.Println(PLAT_USER)
	fmt.Println(PLAT_PASSWORD)
	fmt.Println(CLIENTS_URL)
	fmt.Println(KAFKA_BROKERS)
}
