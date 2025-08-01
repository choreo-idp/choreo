[//]: # (TODO: Update Ballerina builder image with the correct run image)

# Patient Management Service (Mediflow)

## Overview

The **MediFlow** service provides functionalities to manage patient data, including:

- Adding a new patient
- Retrieving patient details by name
- Listing all patients

The service exposes several REST endpoints for performing these operations.

### Health Check

**Endpoint:** `/health`  
**Functionality:** Ensures the service is running.

### Add a new patient

**Endpoint:** `/patients`  
**Method:** `POST`  
**Functionality:** Adds a new patient by sending a JSON payload.

### Retrieve a patient by name

**Endpoint:** `/patients/{name}`  
**Method:** `GET`  
**Functionality:** Retrieves patient details by their name.

### List all patients

**Endpoint:** `/patients`  
**Method:** `GET`  
**Functionality:** Retrieves all patients.

The source code is available at:
https://github.com/wso2/choreo-samples/tree/main/patient-management-service


## Step 1: Deploy the Application

The following command will create the relevant resources in OpenChoreo. It will also trigger a build by creating a build resource.

```bash
kubectl apply -f new-design-sample/from-source/services/ballerina-buildpack-patient-management/patient-management-service.yaml
```

> [!NOTE]
> The build will take around 8 minutes depending on the network speed.

## Step 2: Port-forward the OpenChoreo Gateway

Port forward the OpenChoreo gateway service to access the frontend locally:

```bash
kubectl port-forward -n openchoreo-data-plane svc/gateway-external 8443:443 &
```

## Step 3: Test the Application

   Health check
   ```bash
    curl -k https://dev.openchoreoapis.localhost:8443/default/patient-management-service/mediflow/health
   ```

   Add a new patient
   ```bash
   curl -k -X POST https://dev.openchoreoapis.localhost:8443/default/patient-management-service/mediflow/patients \
   -H "Content-Type: application/json" \
   -d '{
   "name": "Alice",
   "age": 30,
   "condition": "Healthy"
   }'
   ```

   Retrieve a patient by name
   ```bash
    curl -k https://dev.openchoreoapis.localhost:8443/default/patient-management-service/mediflow/patients/Alice
   ```

   List all patients
   ```bash
    curl -k https://dev.openchoreoapis.localhost:8443/default/patient-management-service/mediflow/patients
   ```

## Clean Up

Stop the port forwarding and remove all resources:

```bash
# Find and stop the specific port-forward process
pkill -f "port-forward.*gateway-external.*8443:443"

# Remove all resources
kubectl delete -f new-design-sample/from-source/services/ballerina-buildpack-patient-management/patient-management-service.yaml
```
