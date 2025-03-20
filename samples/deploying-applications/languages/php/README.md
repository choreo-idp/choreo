# Reading List Php web app

## Overview
This is a simple hello world web application built with PHP.

The source code is available at:
https://github.com/wso2/choreo-samples/tree/main/hello-world-php-webapp

## Deploy in Choreo

```bash
choreoctl apply -f https://raw.githubusercontent.com/choreo-idp/choreo/main/samples/deploying-applications/languages/php/hello-world-web-app.yaml
``` 

## Check the Build Workflow Status
You can check the logs of the workflow by running the following command.

```bash
choreoctl logs --type build --build hello-world-web-application-php-build-01 --organization default-org --project default-project --component hello-world-web-application-php
```

## Check the Deployment Status
You can check the deployment logs by running the following command.

```bash
choreoctl logs --type deployment --deployment hello-world-web-application-php-development-deployment-01 --organization default-org --project default-project --component hello-world-web-application-php
```

Note: You should see a k8s namespace created for your org, project and environment combination.

## Access the web application
For this sample, we will use kubectl port-forward to access the web application.

1. Run the following command to port-forward the gateway.

    ```bash
    kubectl port-forward svc/choreo-external-gateway -n choreo-system 8443:443 &
    ```

2. Access the web application from your browser using the following URL.
    - https://hello-world-web-application-php-development.choreoapps.localhost:8443/

## Cleanup

To clean up the resources created by this sample, you can run the following command:

```bash
choreoctl delete -f https://raw.githubusercontent.com/choreo-idp/choreo/main/samples/deploying-applications/languages/php/hello-world-web-app.yaml
```
