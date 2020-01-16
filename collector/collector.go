package main

import (
    "fmt"
    v1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)

func main() {
    config, err := rest.InClusterConfig()
    if err != nil {
        panic(err.Error())
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("%+v\n", pods.Items)

    r, err := clientset.CoreV1().Pods("default").GetLogs("somelogname", &v1.PodLogOptions{}).Do().Get()
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("got [%s] - [%+v]\n", r.GetObjectKind(), r)
}
