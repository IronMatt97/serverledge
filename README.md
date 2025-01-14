![ServerlEdge](docs/logo.png)

Serverledge is a Function-as-a-Service (FaaS) framework designed to
work in Edge-Cloud environments.

Serverledge allows user to define and execute functions across
distributed nodes located in Edge locations or in Cloud data centers.
Each Serverledge node serves function invocation requests using a local
container pool, so that functions are executed within containers for isolation purposes.
When Edge nodes are overloaded, they can to offload computation
to neighbor Edge nodes or to the Cloud.

Serverledge has been first described in a paper presented at [IEEE PerCom 2023](https://www.percom.org/)
(yet to appear):

> G. Russo Russo, T. Mannucci, V. Cardellini, and F. Lo Presti. Serverledge: Decentralized function-
> as-a-service for the edge-cloud continuum. In *Proceedings of IEEE International Conference on
> Pervasive Computing and Communications, PerCom 2023, Atlanta, GA, USA, March 13-17,
> 2023*. IEEE, 2023 (To appear)

Note that Serverledge is a research prototype and it is not (yet) meant to
be used for production use, as it may lack critical security features.


## Building from sources

1. Check that Golang is correctly installed on your machine.

1. Download a copy of the source code.

1. Build the project:

```
$ make
```

You will find executables in `./bin/`.

## Running (single-node deployment)

As functions are executed within Docker containers, you need Docker to
be installed on the host. Furthermore, the Serverledge node needs
permissions to create containers.

You also need an **etcd** server to run Serverledge. To quickly start a local
server:

	$ ./scripts/start-etcd.sh   # stop it with ./scripts/stop-etcd.sh

Start a Serverledge node:

	$ bin/serverledge

### Creating and invoking functions

Register a function `func` from example code:

	$ bin/serverledge-cli create -f func --memory 600 --src examples/hello.py --runtime python310 --handler "hello.handler" 

Invoke `func` with arguments `a=2` and `b=3`:

	$ bin/serverledge-cli invoke -f func -p "a:2" -p "b:3"

For non-trivial inputs, it is recommended to specify function arguments through a
JSON file, instead of using the `-p` flag, as follows:

	$ bin/serverledge-cli invoke -f func --params_file input.json

where `input.json` may contain:

	{
		"a": 2,
		"b": 3
	}

Functions can be also invoked asynchronously using the `--async` flag:

	$ bin/serverledge-cli invoke -f func --async

The server will reply with a `requestID`, which can be used by the client to
poll for the execution result:

	$ bin/serverledge-cli poll --request <requestID>


## Distributed Deployment

[This repository](https://github.com/grussorusso/serverledge-deploy) provides an
Ansible playbook to deploy Serverledge in a distributed configuration.

In this case, you can instruct `serverledge-cli` to
connect to a node other than `localhost` or use a non-default port
by means of environment variables or command-line options:

- Use `--host <HOST>` (or `-H <HOST>`) and/or `--port <PORT>` (or, `-P <PORT>`)
to specify the server
host and port on the command line
- Alternatively, you can set the environment variables
`SERVERLEDGE_HOST` and/or `SERVERLEDGE_PORT`, which are read by the client.

Example:
 
    $ bin/serverledge-cli status -H <host ip-address> -P <port number>

## Configuration

You can provide a configuration file using YAML or TOML syntax. Depending on the
chosen format, the default file name will be `serverledge-conf.yaml` or
`serverledge-conf.toml`. The file can be either placed in `/etc/serverledge`,
in the user `$HOME` directory, or in the working directory where the server is
started.

Alternatively, you can indicate a specific configuration file when starting the
server:

	$ bin/serverledge <config file>

Further information about supported configuration options is available [here](./docs/configuration.md).

### Example

The configuration file may look like this:

	container.pool.memory: 4096
	etcd.address: "1.2.3.4:2379"
	scheduler.queue.capacity: 0
	metrics.enabled: true


## Additional Documentation


 - [Writing functions](./docs/writing-functions.md)
 - [Metrics](./docs/metrics.md)
 - [Serverledge Internals: Executor](./docs/executor.md)


## License

Serverledge is distributed under the terms of the [MIT
license](https://github.com/grussorusso/serverledge/blob/master/LICENSE.txt).
