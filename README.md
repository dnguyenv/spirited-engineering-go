# Spirited Engineering: OpenShift 4 S2I and Webhook

In this tutorial, we will discuss about how to deploy a Golang web application on OpenShift 4 via source to image (aka s2i) framework. We will also talk about how to enable CI/CD using webhook

***Note:*** I use MacOS for this tutorial

**Prerequites**

- A OpenShift cluster
- Github account (https://github.com)
- OpenShift cli client (aka `oc`) (run `$ brew update && brew install openshift-cli` on your Mac) terminal

## Login to your OpenShift cluster 

```bash
$ oc login -u <your-user> -p "access-token" https://<your-cluster-api-url>
```

Example:

```bash
$ oc login -u kubeadmin -p "token" https://console-openshift-console.apps.se.spirited-engineering.os.fyre.se.com:6443
```

## Create new project

```bash
$ oc new-project <project-name>
$ oc project <project-name> #Switch to the newly created project
```

Example:

```bash
$ oc new-project spirited-engineering
$ oc project spirited-engineering 
```
## Create new application under the current project

```bash
$ oc new-app <your-git-hub-repo-uri> --name <your-app-name>
```
Example:

```bash
$ oc new-app https://github.com/dnguyenv/spirited-engineering-go.git --name spirited-engineering-go
```

## Configure OpenShift resources of the app to make the app accessible from outside

Expose the deployment config of the app as a service:

```bash
 $ oc expose dc <your-app-name> --port <a http port>
```

Example:

```bash
$ oc expose dc spirited-engineering-go --port 8080
```

Create a route with tls termination enabled

```bash
$ oc create route edge --service=<your-service-name> --port=<a http port>
```

Example:

```bash
$ oc create route edge --service=spirited-engineering-go --port=8080
```

## Access the app

Get the route information

```bash
$ oc get routes
NAME                      HOST/PORT                                                                     PATH   SERVICES                  PORT   TERMINATION   WILDCARD
spirited-engineering-go   spirited-engineering-go-spirited-engineering.apps.se.os.fyre.se.com          spirited-engineering-go   8080   edge          None
```

You now can access your app from a browser with this url (HOST/PORT value in the output of `oc get routes`): https://spirited-engineering-go-spirited-engineering.apps.se.os.fyre.se.com

## Enable CI/CD 

Now you have the application deployed from your source code on github all the way to your OpenShift cluster. Lets do 1 step further which is to make the application got deployed automatically whenever you've pushed changes to the code on github. There are different ways to do that but in this example, I'll walk you through how to do it with Webhook

