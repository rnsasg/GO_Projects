## Concepts

* The structure of Kubernetes APIs and Resources
* API versioning semantics
* Self-healing
* Garbage Collection and Finalizers
* Declarative vs Imperative APIs
* Level-Based vs Edge-Base APIs
* Resources vs Subresources


## Building services as Kubernetes APIs provides many advantages to plain old REST, including:

* Hosted API endpoints, storage, and validation.
* Rich tooling and CLIs such as kubectl and kustomize.
* Support for `AuthN and granular AuthZ`.
* Support for API evolution through API versioning and conversion.
* Facilitation of adaptive / self-healing APIs that continuously respond to changes in the system state without user intervention.
* Kubernetes as a hosting environment