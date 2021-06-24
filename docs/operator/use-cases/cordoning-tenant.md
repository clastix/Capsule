# Cordoning a Tenant

Bill needs to cordon a Tenant and its Namespaces for several reasons:

- Avoid accidental resource modification(s) including deletion during a Production Freeze Window
- During Kubernetes upgrade, to prevent any workload updates
- During incidents or outages
- During planned maintenance of a dedicated nodes pool in a BYOD scenario

With this said, the Tenant Owner and the related Service Account living into managed Namespaces, cannot proceed to any update, create or delete action.

This is possible just labelling the Tenant as follows:

```shell
$ kubectl label tenant oil capsule.clastix.io/cordon=enabled
tenant oil labeled
```

Any operation performed by Alice, the Tenant Owner, will be rejected.

```shell
$ kubectl --as alice --as-group capsule.clastix.io -n oil-dev create deployment nginx --image nginx
error: failed to create deployment: admission webhook "cordoning.tenant.capsule.clastix.io" denied the request: tenant oil is freezed: please, reach out to the system administrator

$ kubectl --as alice --as-group capsule.clastix.io -n oil-dev delete ingress,deployment,serviceaccount --all
error: failed to create deployment: admission webhook "cordoning.tenant.capsule.clastix.io" denied the request: tenant oil is freezed: please, reach out to the system administrator
```

Uncordoning can be done removing the said label:

```shell
$ kubectl label tenant oil capsule.clastix.io/cordon-
tenant.capsule.clastix.io/oil labeled

$ kubectl --as alice --as-group capsule.clastix.io -n oil-dev create deployment nginx --image nginx
deployment.apps/nginx created
```

# What’s next

This end our tour in Capsule use cases.
As we improve Capsule, more use cases about multi-tenancy, policy admission control, and cluster governance will be covered in the future.
Stay tuned!
