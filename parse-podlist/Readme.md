# Parse Pod List

This folder has code sample supporting the [textParsing the output of "kubectl get pods -A -o json" from a file](https://medium.com/@maninder.bindra/parsing-the-output-of-kubectl-get-pods-a-o-json-from-a-file-81350fa1d757)

## Parsing Sample using JQ

```bash
jq -r '["Namespace,NodeName, PodName, ContainerName","MemoreRequest","MemoryLimit"], (.items[] | [.metadata.namespace , .spec.nodeName , .metadata.name ] + (.spec.containers[] | [.name, .resources.requests.memory, .resources.limits.memory]) ) | @csv' samplePodList.json
```

## Parse using go

```bash
go run main.go --fileName ./samplePodList.json
```
