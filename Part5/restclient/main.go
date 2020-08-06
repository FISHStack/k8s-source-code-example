package main

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", *"/root/.kube/config")
	if err != nil {
		panic(err)
	}
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restclient,err := rest.RESTClientFor(config)
	if err!=nil{
		panic(err)
	}
	result := &corev1.PodList{}

	err = restclient.Get()
	.Namespace("default")
	.Resource("pods")
	.VersionedParams(&metav1.ListOptions{Limit:500},scheme.ParameterCodec)
	.Do()
	.Into(result)

	if err!=nil{
		panic(err)
	}

	for _,d := range result.Items{
		fmt.Printf("NAMESPACE:%v \t NAME:%v \t STATU:%+v \n ",d.Namespace,d.Name,d.Status.Phase)
	}

}
