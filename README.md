# Spirited Engineering: OpenShift 4 S2I and Webhook

In this tutorial, we will discuss about how to deploy a Golang web application on OpenShift 4 via source to image (aka s2i) framework. We will also talk about how to enable CI/CD using webhook

***Note:*** I use MacOS for this tutorial

**Prerequites**
- A OpenShift cluster
- Github account (https://github.com)
- OpenShift cli client (aka `oc`) 

## Login to your OpenShift cluster 
## Create new application under the current project

```bash
$ oc new-app --strategy=source --context-dir=spiritedengineering/theapp --name spirited-engineering-go centos/go-toolset-7-centos7~https://<personal-git-token>@github.ibm.com/dnguyenv/s2i-go
```

Expose the app with a route:

```bash
 $ oc expose dc spirited-engineering-go --port 8080
 $ oc expose service spirited-engineering-go
```

Get the route information

```bash
$ oc get routes
NAME                      HOST/PORT                                                            PATH   SERVICES                  PORT   TERMINATION   WILDCARD
spirited-engineering-go   spirited-engineering-go-duyhard.apps.mcm-aiops-dev.os.fyre.ibm.com          spirited-engineering-go   8080                 None
```

Your app is here: https://spirited-engineering-go-duyhard.apps.mcm-aiops-dev.os.fyre.ibm.com