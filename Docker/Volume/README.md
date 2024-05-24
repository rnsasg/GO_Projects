# Docker Volume

A volume is a mount point on the container’s directory tree where a portion of the host directory tree has been mounted

Semantically, a volume is a tool for segmenting and sharing data that has a scope or life cycle that’s independent of a single container.That makes volumes an important part of any containerized system design that shares or writes files.

- [X] Database software versus database data
- [X] Web application versus log data
- [X] Data processing application versus input and output data
- [X] Web server versus static content
- [X] Products versus support tools

* Volumes enable separation of concerns and create modularity for architectural components. That modularity helps you understand, build, support, and reuse parts of larger systems more easily.

* This separation of relatively static and dynamic file space allows application or image authors to implement
advanced patterns such as polymorphic and composable tools.