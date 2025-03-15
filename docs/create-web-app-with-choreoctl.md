# Deploy a Web Application in Choreo using choreoctl

This section guides you through creating, deploy and accessing a sample Web Application using `choreoctl`.

If you haven't installed Choreo & `choreoctl` yet, please follow the [installation guide](../install/README.md) to install them.

## Step 1 - Create the sample Web Application component

For this, you will be using a sample Web Application component from the [awesome-compose](https://github.com/docker/awesome-compose).

Run the following command to create a sample Web Application component in Choreo.

```shell
choreoctl create component --name hello-world --type WebApplication --git-repository-url https://github.com/docker/awesome-compose --branch master --path /react-nginx --buildpack-name React --buildpack-version 18.20.6
```

You will see the following output:

```text
Component 'hello-world' created successfully in project 'default-project' of organization 'default-org'
```

## Step 2 - Build the created sample component

Create a build resource for hello-world component using Choreo CLI interactive mode.

```shell
choreoctl create build -i
```
Use the build name as 'build1' and keep other inputs as defaults.
```shell
./choreoctl create build -i                                                           config-context  ✭ ✱
Selected inputs:
- organization: default-org
- project: default-project
- component: hello-world
- deployment track: default
- name: build1
- revision: latest
Enter git revision (optional, press Enter to use latest):
Build 'build1' created successfully in project 'default-project' of organization 'default-org'

```

## Step 3 - View build logs and status
```shell
choreoctl logs --type build --component hello-world --build build1 --follow
```
> Note: The build step will take around 5 minutes to get all the dependencies and complete the build.

See the build status using get build resource command.
```shell
choreoctl get build --component hello-world  build1
```
> Note: Proceed to the next step after build  is in 'Completed' status.

## Step 4 - View the generated deployable artifact

As part of the successful build, a deployment artifact resource is created to trigger a deployment.
```shell
choreoctl get deployableartifact --component hello-world
```
## Step 5 - Create a deployment resource

For this option, we will explore the interactive mode which guide you to create the deployment resource.
```shell
choreoctl create deployment -i
```
Name the deployment as 'dev-deployment'. Following is a sample CLI output.
```shell
choreoctl create deployment -i                                                        config-context  ✭ ✱
Selected resources:
- organization: default-org
- project: default-project
- component: hello-world
- deployment track: default
- environment: development
- deployable artifact: default-org-default-project-hello-world-default-foo-0c5ff1ee
- name: dev-deployment
Enter deployment name:
Deployment 'dev-deployment' created successfully in component 'hello-world' of project 'default-project' in organization 'default-org'
```

## Step 6 - View the generated endpoint resource

```shell
choreoctl get endpoint --component hello-world
```
You should see a similar output as follows.
``` shell
NAME     ADDRESS                                                                                 AGE   ORGANIZATION
webapp   https://default-org-default-project-hello-world-ea384b50-development.choreo.localhost   14h   default-org
```
## Step 7 - Test the deployed WebApp

You have two options to test your WebApp component.

1. Option 1: Access the WebApp by exposing the external-gateway as a LoadBalancer to your host machine.
2. Option 2: port-forward from your host machine to external-gateway service.

### Option 1: Expose the external-gateway as a LoadBalancer

The following steps will guide you through exposing the external-gateway service as a LoadBalancer to your host machine.
In this you will be using the [cloud-provider-kind](https://github.com/kubernetes-sigs/cloud-provider-kind/tree/main) to
expose the LoadBalancer service(external-gateway) to your host machine.

First, [install](https://github.com/kubernetes-sigs/cloud-provider-kind/tree/main?tab=readme-ov-file#install) the cloud-provider-kind tool to your host machine.

Then, run this tool in sudo mode, and it will automatically assign LoadBalancer IP to your external-gateway service.

```shell
# run this command in a separate terminal and keep it running.
$ sudo $(which cloud-provider-kind)
```

Then you could find the load balancer IP for your external-gateway service as follows.

```shell
# to find the external-gateway service name
$ kubectl get svc -n choreo-system | grep gateway-external
```

```shell
# to find the LoadBalancer-IP
# <name> should be replaced with the service name found in the previous step.
$ kubectl get svc/<name> -n choreo-system -o=jsonpath='{.status.loadBalancer.ingress[0].ip}'
```

Then add this IP to your /etc/hosts file as follows.

```text
<LoadBalancer-IP> react-starter-development.choreo.localhost
```

Now you can access the WebApp using following URL.

https://default-org-default-project-hello-world-ea384b50-development.choreo.localhost

### Option 2: Port-forward the external-gateway service

The following steps will guide you through port-forwarding from your host machine to the external-gateway service.

First, find the external-gateway service using the following command.

```shell
kubectl get svc -n choreo-system | grep gateway-external
```

Then port-forward the service to your host machine using the following command.

```shell
# <name> should be replaces with the service name found in the previous step.
kubectl port-forward svc/<name> -n choreo-system 443:443
```

Then add the following entry to your /etc/hosts file.

```
127.0.0.1 default-org-default-project-hello-world-ea384b50-development.choreo.localhost
```

Now you can access the WebApp using the following URL.

https://default-org-default-project-hello-world-ea384b50-development.choreo.localhost

## Step 8 - View deployment logs

```shell
choreoctl logs --type deployment --component hello-world --deployment dev-deployment --follow
```