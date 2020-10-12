# Status Monitor

Status Monitor is a lightweight client-server Go application that allows VMs to check in with a server to report the fact that they are still running.

It is intended to replace heavier ICMP or VSphere API means of keeping up with the status of a VM.

Advantages:

- client can be run after automated scripts that prepare a VM to be handed out for x purpose. pings and simple VM status cannot account for this.
- client is responsible for its own status. It sends out periodic small protobuf packets that the server uses to update the cached information.

TODO:

- server can request confirmation of availability from VM
- server can be queried for all available VMs.
- Status page
- Redis backend
- Multiple node support (with Redis)
- Helm Deployment
