package main

import (
    "flag"
    "fmt"
    v1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
    "os"
    "path/filepath"
)

func main() {
    clientset := createClent()

    pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }

    for _, pod := range pods.Items {
        fmt.Printf("getting logs for [%s:%s]\n", pod.Namespace, pod.Name)
        printContainerLog(clientset, pod.Namespace, pod.Name, "l1")
        printContainerLog(clientset, pod.Namespace, pod.Name, "l2")
    }
}

func printContainerLog(client *kubernetes.Clientset, namespace string, pod string, container string) {
    r, err := client.CoreV1().Pods(namespace).GetLogs(pod, &v1.PodLogOptions{Container: container, Follow: true}).Stream()
    if err != nil {
        panic(err.Error())
    }

    fmt.Printf("%s", r)

    buffer := make([]byte, 100)
    for {
        n, err := r.Read(buffer)
        if err != nil {
            panic(err)
        }
        if n <= 0 {
            fmt.Println("end of log")
            break
        }
        fmt.Printf("got [%s]", buffer)
    }
}

func createInClient() *kubernetes.Clientset {
    config, err := rest.InClusterConfig()
    if err != nil {
        panic(err.Error())
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    return clientset
}

func createClent() *kubernetes.Clientset {
    var kubeconfig *string
    if home := homeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    // use the current context in kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err.Error())
    }

    // create the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    return clientset
}

func homeDir() string {
    if h := os.Getenv("HOME"); h != "" {
        return h
    }
    return os.Getenv("USERPROFILE") // windows
}
