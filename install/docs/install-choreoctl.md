# Installing `choreoctl`
[//]: # (TODO: Refine this once we release the CLI as a binary.)

`choreoctl` is the command-line interface for Choreo. With that, you can seamlessly interact with Choreo and manage your resources.

## _Prerequisites_

1. Make sure you have installed [Go](https://golang.org/doc/install), version 1.23.5.
2. Make sure to clone the repository into your local machine.
   ```shell
   git clone https://github.com/choreo-idp/choreo.git
   ```


## Step 1 - Build `choreoctl`
From the root level of the repo, run:

```shell
make choreoctl-relase
```

Once this is completed, it will have a `dist` directory created in the project root directory.

## Step 2 - Install `choreoctl` into your host machine

Run the following command to install the `choreoctl` CLI into your host machine.

```shell
./dist/choreoctl/choreoctl-install.sh
````

To verify the installation, run:

```shell
choreoctl
```

You should see the following output:

```text
Welcome to Choreo CLI, the command-line interface for Open Source Internal Developer Platform

Usage:
  choreoctl [command]

Available Commands:
  apply       Apply Choreo resource configurations
  completion  Generate the autocompletion script for the specified shell
  config      Manage Choreo configuration contexts
  create      Create Choreo resources
  get         Get Choreo resources
  help        Help about any command
  logs        Get Choreo resource logs

Flags:
  -h, --help   help for choreoctl

Use "choreoctl [command] --help" for more information about a command.
```

Now it's all setup. It's time to deploy your first component in Choreo.

If you haven't installed Choreo yet, please follow the [installation guide](../README.md) to install it.
