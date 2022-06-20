package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	v1 "k8s.io/api/core/v1"
)

func main() {

	file := flag.String("fileName", "samplePodList.json", "Full file name with path")
	flag.Parse()

	var podList v1.PodList

	byts, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal("could not open file")
	}

	err = json.Unmarshal([]byte(byts), &podList)
	if err != nil {
		log.Fatal("could not unmarshall the data")
	}

	fmt.Printf("%s,%s,%s,%s,%s,%s\n", "Namespace", "NodeName", "PodName", "ContainerName", "MemoryRequest", "MemoryLimit")
	for _, pod := range podList.Items {
		for _, container := range pod.Spec.Containers {
			fmt.Printf("%s,%s,%s,%s,%s,%s\n", pod.Namespace, pod.Spec.NodeName, pod.Name, container.Name, container.Resources.Requests.Memory().String(), container.Resources.Limits.Memory().String())
		}

	}
}
