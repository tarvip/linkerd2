package installsp

// Template provides the base template for the `linkerd install-sp` command.
const Template = `apiVersion: linkerd.io/v1alpha2
kind: ServiceProfile
metadata:
  name: linkerd-dst.{{.Namespace}}.svc.{{.ClusterDomain}}
  namespace: {{.Namespace}}
spec:
  routes:
  - name: POST /io.linkerd.proxy.destination.Destination/Get
    condition:
      method: POST
      pathRegex: /io\.linkerd\.proxy\.destination\.Destination/Get
  - name: POST /io.linkerd.proxy.destination.Destination/GetProfile
    condition:
      method: POST
      pathRegex: /io\.linkerd\.proxy\.destination\.Destination/GetProfile
`
