[![Actions Status](https://github.com/jfbramlett/nwp-platform-go/workflows/Go/badge.svg)](https://github.com/jfbramlett/nwp-platform-go/actions)

# NWP-Platform-Gp
Sample version of the nwp-platform written in Go. There are two components, a platform and a sample financial institution. 
There is a single endpoint for accountlist with the platform taking a protocol-specific endpoint (dda10), converting that
request to EEL and then calling the fi server. The fi server processes the request as an EEL request and returns
back to the platform who in turn translares the EEL response to the protocol (dda10) response.