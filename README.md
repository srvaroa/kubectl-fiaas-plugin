Kubectl FIAAS plugins
=====================

This repo contains a PoC for some [Kubectl
plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/)
that simplify accessing Kubernetes resources created by
[FIAAS](https://fiaas.github.io).  FIAAS is a PaaS like abstraction on
top of Kubernetes used by several online marketplaces in
[Schibsted](https://schibsted.com) and [Adevinta](https://adevinta.com).

At the moment the available plugins are:

- `kubectl fiaas pods $app_name`: return all pods for the given FIAAS
  application
- `kubectl fiaas logs $app_name`: return logs for all pods of the given
  FIAAS application.

Both commands use the `--namespace` (-n) flag to scope the query to a
specific namespace.

Installation
------------

Build the binaries with

    make

Then get then in your PATH.  In my case I use a ~/.bin folder:

    ln -s $(pwd)/bin/* ~/.bin/

Usage
-----

    $ kubectl fiaas pods badger
    NAME                      READY   STATUS    RESTARTS   AGE
    badger-569f4f686c-s8xk7   1/1     Running   0          2d2h
    badger-569f4f686c-wzsdb   1/1     Running   0          2d2h

    $ kubectl fiaas logs badger
    ... a bunch of logs ...


In both cases the raw Kubernetes command would be simple:

    $ kubectl get pods -lapp=badger
    $ kubectl logs -lapp=badger

But the plugins let the user deal with a higher level of abstraction
(the app) without needing to know how FIAAS tags resources in
Kubernetes.  The intent is to provide more plugins that solve more
complex use cases.
