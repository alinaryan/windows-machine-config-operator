# Creating an AWS Windows MachineSet

_\<windows_container_ami\>_ can be used as it is without any modification.
You must use Windows Server 2019 with a version 10.0.17763.1457 or earlier to work
around Windows containers behind a Kubernetes load balancer becoming unreachable
[issue](https://github.com/microsoft/Windows-Containers/issues/78).
                                                                           
_\<infrastructureID\>_ should be replaced with the output of:
```shell script
 oc get -o jsonpath='{.status.infrastructureName}{"\n"}' infrastructure cluster
```
_\<region\>_ should be replaced with a valid AWS region like `us-east-1`.
_\<zone\>_ should be replaced with a valid AWS availability zone like `us-east-1a`.

```
apiVersion: machine.openshift.io/v1beta1
kind: MachineSet
metadata:
  labels:
    machine.openshift.io/cluster-api-cluster: <infrastructureID> 
  name: <infrastructureID>-windows-worker-<zone>
  namespace: openshift-machine-api
spec:
  replicas: 1
  selector:
    matchLabels:
      machine.openshift.io/cluster-api-cluster: <infrastructureID> 
      machine.openshift.io/cluster-api-machineset: <infrastructureID>-windows-worker-<zone>
  template:
    metadata:
      labels:
        machine.openshift.io/cluster-api-cluster: <infrastructureID> 
        machine.openshift.io/cluster-api-machine-role: worker
        machine.openshift.io/cluster-api-machine-type: worker
        machine.openshift.io/cluster-api-machineset: <infrastructureID>-windows-worker-<zone>
        machine.openshift.io/os-id: Windows
    spec:
      metadata:
        labels:
          node-role.kubernetes.io/worker: ""
      providerSpec:
        value:
          ami:
            id: <windows_container_ami>
          apiVersion: awsproviderconfig.openshift.io/v1beta1
          blockDevices:
            - ebs:
                iops: 0
                volumeSize: 120
                volumeType: gp2
          credentialsSecret:
            name: aws-cloud-credentials
          deviceIndex: 0
          iamInstanceProfile:
            id: <infrastructureID>-worker-profile 
          instanceType: m5a.large
          kind: AWSMachineProviderConfig
          placement:
            availabilityZone: <zone>
            region: <region>
          securityGroups:
            - filters:
                - name: tag:Name
                  values:
                    - <infrastructureID>-worker-sg 
          subnet:
            filters:
              - name: tag:Name
                values:
                  - <infrastructureID>-private-<zone>
          tags:
            - name: kubernetes.io/cluster/<infrastructureID> 
              value: owned
          userDataSecret:
            name: windows-user-data
            namespace: openshift-machine-api
```
