# Amadeus-go
This projects aims to serve as a wrapper for [Amadeus](https://developers.amadeus.com) API, which gives flight and hotel information. The hotel however, is not yet ready but the rest of it (overalls to 12 API calls total) is ready. This project was implemented using the following technologies and libraries.
- [go-kit](https://gokit.io/) as a microsevice toolkit
- [gRPC](https://grpc.io/) as the transport layer
- [protobuf](https://developers.google.com/protocol-buffers/) for serialization
- [consul](https://www.consul.io/) for service discovery
- [redis](https://redis.io/) for caching some stuff such as token

## Quickstart
Firstly start by cloning the project into your local computer using the following command:
```bash
git clone github.com:meysam81/amadeus-go.git
```

After doing that, change your directory into the porject's root directory:
```bash
cd amadeus-go
```

Doing the above, you'll be able to `make` some stuff that is in the [Makefile](Makefile). From the entries that already exists, you surely need to do the following as they are mandatory.
```bash
make proto
```
This command will generate the compiled proto files for you in their right directory. **this command has to be entered before going any futher**

You can now run the server using the following command:
```bash
make dev_run
```

Also if you want to test the server, there's a built-in client in [cmd/cli/cli.go](cmd/cli/cli.go). To run it use the following command:
```bash
make dev_cli
```

You can also run it in container as there's a [Dockerfile](Dockerfile) in the project's root directory. So runnning the following commands will build and run it.
```bash
make build
make up
```

If you want to stop the container you can enter the following command into your terminal:
```bash
make down
```

## Contribute
Feel free to contribute at anytime, PR's are welcomed with pleasure :smiley:

