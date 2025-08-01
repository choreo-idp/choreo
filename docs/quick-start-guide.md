## Quick Start Guide

Setting up OpenChoreo in a Kubernetes cluster involves multiple steps and tools. This guide provides a fast and simple way to install a fully functional OpenChoreo instance on your local machine with minimal prerequisites and effort by using a pre-configured dev container.
This dev container has all the necessary tools installed for setting up OpenChoreo and is ready to be used. Once the installation is complete, you can explore the underlying setup to understand how it works.
When you're done, you can fully clean up the setup, leaving your machine clutter-free.

#### Prerequisites

To get started, you’ll need:

- **Docker** – Just have it installed on your machine, and you're good to go.
    - We recommend using [Docker Engine version 26.0+](https://docs.docker.com/engine/release-notes/26.0/).
- **5–10 minutes of your time** – Setup is quick and depends on your internet connection speed.

#### Start the Dev Container

Run the following command to start the dev container and launch a terminal session within it:

```shell
docker run --rm -it --name openchoreo-quick-start \
-v /var/run/docker.sock:/var/run/docker.sock \
-v openchoreo-state:/state \
-v tf-state:/app/terraform \
--network bridge \
-p 8443:8443 \
-p 7007:7007 \
ghcr.io/openchoreo/quick-start:latest

```

#### Install OpenChoreo
This process sets up a [KinD](https://kind.sigs.k8s.io/) (Kubernetes-in-Docker) cluster in your Docker environment and installs OpenChoreo along with its dependencies.

To begin the installation, run:

```shell
./install.sh
```

> [!TIP]
> If you previously used this setup and encounter errors during installation, ensure you perform the proper cleanup as outlined in the [Cleaning up](#cleaning-up) section before starting again.

Once the installation is complete, you will see the following confirmation message:
```text
======================================================================
                     OpenChoreo Component Status                     
======================================================================

+- Networking (Infrastructure) -----------------------------------+
| cilium                    [READY]                            |
+-----------------------------------------------------------------+

+- Control Plane (Core) ------------------------------------------+
| cert_manager_cp           [READY]                            |
| controller_manager        [READY]                            |
| api_server                [READY]                            |
+-----------------------------------------------------------------+

+- Data Plane (Core) ---------------------------------------------+
| vault                     [READY]                            |
| registry                  [READY]                            |
| redis                     [READY]                            |
| envoy_gateway             [READY]                            |
| external_gateway          [PENDING]                          |
| internal_gateway          [PENDING]                          |
| fluent_bit_dp             [PENDING]                          |
+-----------------------------------------------------------------+

+- Build Plane (Optional) ----------------------------------------+
| build_plane               [READY]                            |
+-----------------------------------------------------------------+

+- Identity Provider (Optional) ----------------------------------+
| identity_provider         [READY]                            |
+-----------------------------------------------------------------+

+- Observability Plane (Optional) --------------------------------+
| opensearch                [NOT_INSTALLED]                    |
| opensearch_dashboard      [NOT_INSTALLED]                    |
| observer                  [NOT_INSTALLED]                    |
+-----------------------------------------------------------------+

+- Backstage Demo (Optional) -------------------------------------+
| backstage                 [READY]                            |
+-----------------------------------------------------------------+

======================================================================

Setting up OpenChoreo DataPlane

Running in single-cluster mode using context 'kind-openchoreo-quick-start'
kind-openchoreo-quick-start

Applying DataPlane to context: kind-openchoreo-quick-start
dataplane.openchoreo.dev/default created

DataPlane applied to 'default' successfully!

Setting up OpenChoreo BuildPlane

kind-openchoreo-quick-start

Applying BuildPlane to context: kind-openchoreo-quick-start
buildplane.openchoreo.dev/default created

BuildPlane applied to 'default' successfully!
```

> [!NOTE]
> If you see any components are still in the `pending` state, you can check the status again in few minutes. 
> Run the following command to check the installation status of components
> ```shell
> ./check-status.sh
> ```

#### Deploying a Web Application with OpenChoreo

You now have OpenChoreo fully setup in your docker environment.
Next, lets deploy a sample Web Application by running the following command:

```shell
./deploy_web_application.sh
```

Once the deployment is complete, you will receive the following message together with a URL to access the application:

```text
✅ Endpoint is ready!
🌍 You can now access the Sample Web Application at: https://react-starter-image-development.choreoapps.localhost:8443
```

### Understanding What Happens Behind the Scenes
By following the install and deploy web application commands, you first, setup OpenChoreo and then, successfully deployed and accessed a fully functional Web Application.

Let’s now explore what happens after each command.

#### 1. The Install Command
- A dev container with all the necessary tools for OpenChoreo to run is set up in a local Docker environment.
- A KinD Kubernetes cluster is created, where the OpenChoreo IDP and its dependencies were installed using Helm charts.

#### Foundation Resources Created by OpenChoreo

The installation process, by default, sets up several essential abstractions. These are:
- Organization
- [Dataplane](https://github.com/openchoreo/openchoreo/tree/main/docs#dataplane)
- [Environments](https://github.com/openchoreo/openchoreo/tree/main/docs#environment) (e.g., Development, Staging, Production)
- [Deployment Pipeline](https://github.com/openchoreo/openchoreo/tree/main/docs#deploymentpipeline) for the environments
- [Project](https://github.com/openchoreo/openchoreo/tree/main/docs#project)

To access the artifacts created in OpenChoreo you can use choreoctl as shown in the following commands:

First you can get the current context
```shell
choreoctl config current-context
```
Next you can check each resource type

```shell
choreoctl get organizations
```
```shell
choreoctl get dataplanes
```
```shell
choreoctl get environments
```
```shell
choreoctl get projects
```
To get more details on any of these abstractions, you can use a similar command to the following command:

```shell
choreoctl get project default-project -oyaml
```

#### 2. The Deploy Web Application Command
The deploy script creates a sample Web Application [Component](https://github.com/openchoreo/openchoreo/tree/main/docs#component), along with a [Deployment](https://github.com/openchoreo/openchoreo/tree/main/docs#component) for the sample web application.

To inspect these resources in more detail, run the following commands:

```shell
choreoctl get components
```
```shell
choreoctl get deployment --component react-starter-image
```

OpenChoreo generates a [DeployableArtifact](https://github.com/openchoreo/openchoreo/tree/main/docs#deployableartifact) and an [Endpoint](https://github.com/openchoreo/openchoreo/tree/main/docs#endpoint) to access the running application:

```shell
choreoctl get deployableartifact --component react-starter-image
```
```shell
choreoctl get endpoint --component react-starter-image
```

You can also check out the logs for the sample Web Application Deployment with the following command:
```shell
choreoctl logs --type deployment --component react-starter-image --deployment react-starter-image-deployment --follow
```

> [!TIP]
> You can try out more samples on this setup. Check out our [Samples](../samples/README.md) for more details.
> The samples are also available in the dev container at `/app/samples`.

> [!IMPORTANT]
> If you try out samples in the dev container, you don't need to expose the OpenChoreo Gateway to your host machine. The OpenChoreo gateway is already exposed to host machine via dev container port 8443.

### Cleaning up
After finishing your work, you have two options:

1. **Exit and return later**: If you plan to return, simply exit the dev container by running:
    ```shell
    exit
    ```
2. **Full cleanup**: To remove all resources and clean up the environment completely, run:
    ```shell
    ./uninstall.sh
    ```
    ```shell
    exit
    ```
    ```shell
    docker volume rm choreo-state tf-state
    ```
> [!Note]
> The `tf-state` Docker volume is used to persist the installation state, so if you exit the dev container before completing the uninstallation and return later, your installation progress will still be there.
> The `choreo-state` volume is used to store the kubeconfig of the created KinD cluster, ensuring that the cluster configuration remains available even after restarting the container.

That's it!

Now you understand how OpenChoreo simplifies the deployment and management of cloud-native applications.
