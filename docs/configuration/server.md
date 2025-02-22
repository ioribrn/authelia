---
layout: default
title: Server
parent: Configuration
nav_order: 11
---

# Server

The server section configures and tunes the http server module Authelia uses.

## Configuration

```yaml
server:
  host: 0.0.0.0
  port: 9091
  path: ""
  read_buffer_size: 4096
  write_buffer_size: 4096
  enable_pprof: false
  enable_expvars: false
  disable_healthcheck: false
  tls:
    key: ""
    certificate: ""
```

## Options

## host
<div markdown="1">
type: string
{: .label .label-config .label-purple } 
default: 0.0.0.0
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Defines the address to listen on. See also [port](#port). Should typically be `0.0.0.0` or `127.0.0.1`, the former for
containerized environments and the later for daemonized environments like init.d and systemd.

Note: If utilising an IPv6 literal address it must be enclosed by square brackets and quoted:

```yaml
host: "[fd00:1111:2222:3333::1]"
```

### port
<div markdown="1">
type: integer
{: .label .label-config .label-purple } 
default: 9091
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Defines the port to listen on. See also [host](#host).

### path
<div markdown="1">
type: string 
{: .label .label-config .label-purple } 
default: ""
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Authelia by default is served from the root `/` location, either via its own domain or subdomain.

Modifying this setting will allow you to serve Authelia out from a specified base path. Please note
that currently only a single level path is supported meaning slashes are not allowed, and only
alphanumeric characters are supported.

Example: https://auth.example.com/, https://example.com/
```yaml
server:
  path: ""
```

Example: https://auth.example.com/authelia/, https://example.com/authelia/
```yaml
server:
  path: authelia
```

### read_buffer_size
<div markdown="1">
type: integer 
{: .label .label-config .label-purple } 
default: 4096
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Configures the maximum request size. The default of 4096 is generally sufficient for most use cases.

### write_buffer_size
<div markdown="1">
type: integer 
{: .label .label-config .label-purple } 
default: 4096
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Configures the maximum response size. The default of 4096 is generally sufficient for most use cases.

### enable_pprof
<div markdown="1">
type: boolean
{: .label .label-config .label-purple } 
default: false
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Enables the go pprof endpoints.

### enable_expvars
<div markdown="1">
type: boolean
{: .label .label-config .label-purple } 
default: false
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Enables the go expvars endpoints.

### disable_healthcheck
<div markdown="1">
type: boolean
{: .label .label-config .label-purple } 
default: false
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

On startup Authelia checks for the existence of /app/healthcheck.sh and /app/.healthcheck.env and if both of these exist
it writes the configuration vars for the healthcheck to the /app/.healthcheck.env file. In instances where this is not
desirable it's possible to disable these interactions entirely.

An example situation where this is the case is in Kubernetes when set security policies that prevent writing to the
ephemeral storage of a container or just don't want to enable the internal health check.

### tls

Authelia typically listens for plain unencrypted connections. This is by design as most environments allow to
security on lower areas of the OSI model. However it required, if you specify both the [tls key](#key) and 
[tls certificate](#certificate) options, Authelia will listen for TLS connections.

#### key
<div markdown="1">
type: string (path)
{: .label .label-config .label-purple } 
default: ""
{: .label .label-config .label-blue }
required: situational
{: .label .label-config .label-yellow }
</div>

The path to the private key for TLS connections. Must be in DER base64/PEM format.

#### certificate
<div markdown="1">
type: string (path)
{: .label .label-config .label-purple } 
default: ""
{: .label .label-config .label-blue }
required: situational
{: .label .label-config .label-yellow }
</div>

The path to the public certificate for TLS connections. Must be in DER base64/PEM format.

## Additional Notes

### Buffer Sizes

The read and write buffer sizes generally should be the same. This is because when Authelia verifies
if the user is authorized to visit a URL, it also sends back nearly the same size response as the request. However
you're able to tune these individually depending on your needs.
