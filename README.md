# wgrest

[![Go Report Card](https://goreportcard.com/badge/github.com/rawmind0/wgrest)](https://goreportcard.com/report/github.com/rawmind0/wgrest)

WireGuard REST API

WireGuard is an simple and modern VPN. It is cross-platform (Windows, macOS, BSD, iOS, Android).

👉 Looking for Vue/React js developer to do UI interface for wgrest. Since there is no UI the first initiator could choose technology.
For more details get in touch with me.


## Install

```shell
curl -L https://github.com/suquant/wgrest/releases/download/v0.0.1/wgrest-linux-amd64 -o wgrest

chmod +x wgrest
```

## Build

To build the WireGuard REST API, 2 steps may be needed:

* generate the API code with `go-swagger`, executing `make generate`
* build the generated API code, executing `make build` . It will generate `./dist/wgrest` bianry
* (Optional) cross build the generated API code, executing `make all`. It will generate amd64 and arm64 binaries at `./dist` folder

## Run

WireGuard REST API Server provides these executing options:

```shell
wgrest -h
Usage:
  wgrest [OPTIONS]

Manage WireGuard VPN tunnels by RESTful manner

Application Options:
      --scheme=                 the listeners to enable, this can be repeated and defaults to the schemes in the swagger spec
      --cleanup-timeout=        grace period for which to wait before killing idle connections (default: 10s)
      --graceful-timeout=       grace period for which to wait before shutting down the server (default: 15s)
      --max-header-size=        controls the maximum number of bytes the server will read parsing the request header's keys and values, including the request line. It does not limit the size of the request body. (default: 1MiB)
      --socket-path=            the unix socket to listen on (default: /var/run/wgrest.sock)
      --host=                   the IP to listen on (default: localhost) [$HOST]
      --port=                   the port to listen on for insecure connections, defaults to a random value [$PORT]
      --listen-limit=           limit the number of outstanding requests
      --keep-alive=             sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download) (default: 3m)
      --read-timeout=           maximum duration before timing out read of the request (default: 30s)
      --write-timeout=          maximum duration before timing out write of the response (default: 60s)
      --tls-host=               the IP to listen on for tls, when not specified it's the same as --host [$TLS_HOST]
      --tls-port=               the port to listen on for secure connections, defaults to a random value [$TLS_PORT]
      --tls-certificate=        the certificate to use for secure connections [$TLS_CERTIFICATE]
      --tls-key=                the private key to use for secure connections [$TLS_PRIVATE_KEY]
      --tls-ca=                 the certificate authority file to be used with mutual tls auth [$TLS_CA_CERTIFICATE]
      --tls-listen-limit=       limit the number of outstanding requests
      --tls-keep-alive=         sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download)
      --tls-read-timeout=       maximum duration before timing out read of the request
      --tls-write-timeout=      maximum duration before timing out write of the response

Auth flags:
      --token=                  authentication token

Cors Options:
      --cors-allow-origins=     Configure cors Access-Control-Allow-Origin header (default: [*])
      --cors-allow-credentials  Enable cors Access-Control-Allow-Credentials header
      --cors-debug              Enable cors debug

Help Options:
  -h, --help                    Show this help message
```

### Execute server 

```shell
wgrest --token=secret --scheme=http --port=8000
```

## Create **wg0** device

```shell
curl -v -g \
    -H "Accept: */*" \
    -H "Content-Type: application/json" \
    -H "Token: secret" \
    -X POST \
    -d '{
        "name": "wg0", 
        "listen_port":51820, 
        "private_key": "cLmxIyJx/PGWrQlevBGr2LQNOqmBGYbVfu4XcRO2SEo=", 
        "network": "10.10.1.1/24"
    }' \
    http://127.0.0.1:8000/devices/
```

## Get devices

```shell
curl -v -g \
    -H "Accept: */*" \
    -H "Content-Type: application/json" \
    -H "Token: secret" \
    -X GET \
    http://127.0.0.1:8000/devices/
```

## Add peer

```shell
curl -v -g \
    -H "Accept: */*" \
    -H "Content-Type: application/json" \
    -H "Token: secret" \
    -X POST \
    -d '{
        "public_key": "hQ1yeyFy+bZn/5jpQNNrZ8MTIGaimZxT6LbWAkvmKjA=", 
        "allowed_ips": ["10.10.1.2/32"], 
        "preshared_key": "uhFI9c9rInyxqgZfeejte6apHWbewoiy32+Bo34xRFs="
    }' \
    http://127.0.0.1:8000/devices/wg0/peers/
```

## Get peers

```shell
curl -v -g \
    -H "Accept: */*" \
    -H "Content-Type: application/json" \
    -H "Token: secret" \
    -X GET \
    http://127.0.0.1:8000/devices/wg0/peers/
```

## Delete peer

Since the wireguard public key is the standard base64 encoded string, it is not safe to use in URI schema,
is that reason peer_id contains the same public key of the peer but encoded with URL safe base64 encoder.

peer_id can be retrieved either by `peer_id` field from peer list endpoint or by this rule

```shell
python3 -c "import base64; \
    print(\
        base64.urlsafe_b64encode(\
            base64.b64decode('hQ1yeyFy+bZn/5jpQNNrZ8MTIGaimZxT6LbWAkvmKjA=')\
        ).decode()\
    )"
```

delete peer request

```shell
curl -v -g \
    -H "Accept: */*" \
    -H "Content-Type: application/json" \
    -H "Token: secret" \
    -X DELETE \
    http://127.0.0.1:8000/devices/wg0/peers/hQ1yeyFy-bZn_5jpQNNrZ8MTIGaimZxT6LbWAkvmKjA=/
```
