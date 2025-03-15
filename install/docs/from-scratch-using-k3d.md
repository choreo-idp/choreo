# Installing from Scratch with a k3d (k3s in Docker) Cluster

In this section, you'll learn how to set up a [k3d](https://k3d.io/) cluster and install Choreo from scratch.

## _Prerequisites_

1. Make sure you have installed [k3d](https://k3d.io/stable/), version v5.8+.

   > We use k3d to quickly create a Kubernetes cluster, primarily for testing purposes.

   To verify the installation:

    ```shell
    k3d version
    ```
2. Make sure you have installed [Helm](https://helm.sh/docs/intro/install/), version v3.15+.

   > We use Helm as the package manager to install the required artifacts into the Kubernetes cluster.

   To verify the installation:

    ```shell
    helm version
    ```
3. Clone our repository and navigate to the `install` directory.

    ```shell
    git clone https://github.com/choreo-idp/choreo.git && cd choreo/install
    ```
All set up. Let's go ahead and install Choreo.

## Installing Choreo

### Step 1 - Create a k3d cluster

```shell
$ k3d cluster create --config k3d/k3d-config.yaml
```

### Step 2 - Install Cilium

The following helm chart provided by us installs Cilium with minimal configurations required for Choreo.

```shell
helm install cilium oci://ghcr.io/choreo-idp/helm-charts/cilium  --version 0.1.0 --namespace "choreo-system" --create-namespace --timeout 30m
```

### Step 3 - Install Choreo

```shell
helm install choreo oci://ghcr.io/choreo-idp/helm-charts/choreo  --version 0.1.0 --namespace "choreo-system" --create-namespace --timeout 30m
```

### Step 4 - Verify installation status

```shell
./check-status.sh
```

You should see the following output if the installation is successful.

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

All set! You have successfully installed Choreo using Kind. 🎉 You can now start using Choreo. Check out our [samples](../../samples) and [docs](../../docs) to get started.
