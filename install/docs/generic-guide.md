# Generic guide to Install Choreo using Helm Chart

## _Prerequisites_

- [Helm](https://helm.sh/docs/intro/install/) version v3.15+
  > Choreo use the Helm as the package manager to install the required artifacts into the kubernetes cluster.
- [Cilium](https://docs.cilium.io/en/stable/gettingstarted/k8s-install-default/#install-cilium) installed kubernetes cluster
    - Cilium version v1.15.10
    - kubernetes version v1.22.0+
  > Cilium is a dependency for choreo to operate. It uses the Cilium CNI plugin to manage the network policies and security for the pods in the cluster.


## Install

You can directly install Choreo using the Helm chart provided in our registry.

```shell
helm install choreo oci://ghcr.io/choreo-idp/helm-charts/choreo \
--version 0.1.0 --namespace "choreo-system" --create-namespace --timeout 30m
```

## Verify installation status

You can see the list of installed components and their status using the following command.

```shell
curl -sL https://raw.githubusercontent.com/choreo-idp/choreo/refs/heads/main/install/check-status.sh | bash
```
you should see the following output if the installation is successful.

```text
Choreo Installation Status:

Component                 Status         
------------------------  ---------------
cilium                    ✅ ready
vault                     ✅ ready
argo                      ✅ ready
cert_manager              ✅ ready
choreo_controller         ✅ ready
choreo_image_registry     ✅ ready
envoy_gateway             ✅ ready
redis                     ✅ ready
external_gateway          ✅ ready
internal_gateway          ✅ ready

Overall Status: ✅ READY
🎉 Choreo has been successfully installed and is ready to use! 🚀
```


