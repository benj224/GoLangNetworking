package networking

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/perlin-network/noise"
)

func Init() {
	ip := GetPublicIp()
	selfNode, err := noise.NewNode(noise.WithNodeBindHost(net.ParseIP(ip)), noise.WithNodeBindPort(49972))
	if err != nil {
		panic(err)
	}

	if _, err := selfNode.Ping(context.TODO(), alice.Addr()); err != nil {
		panic(err)
	}

}

func GetPublicIp() string {
	url := "https://api.ipify.org?format=text"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(ip)
}
