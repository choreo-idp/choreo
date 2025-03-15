# Install Choreo

The main installation method of Choreo is by using the Helm charts provided by us. This method is suitable for both development and production environments.
We have several installation methods. You can see the installation methods below.

- **[Generic guide to Install Choreo using Helm Chart](docs/generic-guide.md)** - This guide provides a generic way to install Choreo using the Helm chart.
- **[From scratch using Kind](docs/from-scratch-using-kind.md)** - This guide provides a way to install Choreo from scratch using [kind](https://kind.sigs.k8s.io/).
- **[From scratch using k3d](docs/from-scratch-using-k3d.md)** - This guide provides a way to install Choreo from scratch using [k3d](https://k3d.io/).

All these methods we are using our published helm charts.
- **[Cilium](https://github.com/choreo-idp/choreo/pkgs/container/helm-charts%2Fcilium)**
- **[Choreo](https://github.com/choreo-idp/choreo/pkgs/container/helm-charts%2Fchoreo)**

If you interested in installing from the source code, you can use the provided scripts in the `install` directory.

```shell
./install.sh
```
