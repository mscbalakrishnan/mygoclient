package main

import (
	"fmt"

	mycrypto "github.com/mscbalakrishnan/mygoclient/crypto"
	"github.com/mscbalakrishnan/mygoclient/k8sclient"
	"golang.org/x/text/message"
)

func main() {

	cluster := k8sclient.NewCluster("default")

	k8sResourceList := cluster.GetAllDeployments()

	fmt.Println("deployments listssss: \n")
	for _, k8sRes := range k8sResourceList {
		k8sRes.ToString()
	}

	p := message.NewPrinter(message.MatchLanguage("ta"))
	p.Println("tamil")

	emsg := mycrypto.EncryptMessage("This is a Balakrishnan")
	mycrypto.DecryptMessage(emsg)

	fmt.Println("Md5 Sum: {}", mycrypto.GetMD5Hash("Balakrishnan"))

}
