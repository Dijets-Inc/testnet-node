<div align="center">
  <img src="resources/DijetsNode.png?raw=true">
</div>

---

A Golang Node implementation for Dijets network.

## Installation

Dijets is an incredibly lightweight protocol, so the minimum computer requirements are quite modest.
Note that as network usage increases, hardware requirements may change.

The minimum recommended hardware specification for nodes connected to Mainnet is:

- CPU: Equivalent of 8 AWS vCPU
- RAM: 16 GiB
- Storage: 1 TiB
- OS: Ubuntu 20.04/22.04 or macOS >= 12
- Network: Reliable IPv4 or IPv6 network connection, with an open public port.

If you plan to build DijetsNodeGo from source, you will also need the following software:

- [Go](https://golang.org/doc/install) version >= 1.18.1
- [gcc](https://gcc.gnu.org/)
- g++

### Native Install

Clone the DijetsNodeGo repository:

```sh
git clone git@github.com:Dijets-Inc/dijetsnodego.git
cd dijetsnodego
```

This will clone and checkout to `master` branch.

#### Building the Dijets Executable

Build Dijets by running the build script:

```sh
./scripts/build.sh
```

The output of the script will be the Dijets binary named `dijetsnodego`. It is located in the build directory:

```sh
./build/dijetsnodego
```

### Binary Install

Download the [latest build](https://github.com/Dijets-Inc/dijetsnodego/releases/latest) for your operating system and architecture.

The Dijets binary to be executed is named `dijetsnodego`.

### Docker Install

Make sure docker is installed on the machine - so commands like `docker run` etc. are available.

Building the docker image of latest dijetsnodego branch can be done by running:

```sh
./scripts/build_image.sh
```

To check the built image, run:

```sh
docker image ls
```

The image should be tagged as `hyphenesc/dijetsnodego:xxxxxxxx`, where `xxxxxxxx` is the shortened commit of the Dijets source it was built from. To run the dijets node, run:

```sh
docker run -ti -p 9650:9650 -p 9651:9651 hyphenesc/dijetsnodego:xxxxxxxx /dijetsnodego/build/dijetsnodego
```

## Running Dijets

### Connecting to Mainnet

To connect to the Dijets Mainnet, run:

```sh
./build/dijetsnodego
```

You should see some pretty ASCII art and log messages.

You can use `Ctrl+C` to kill the node.

### Connecting to Fuji

To connect to the Fuji Testnet, run:

```sh
./build/dijetsnodego --network-id=fuji
```

## Bootstrapping

A node needs to catch up to the latest network state before it can participate in consensus and serve API calls. This process, called bootstrapping, currently takes several days for a new node connected to Mainnet.

A node will not [report healthy](https://docs.djtx.network/build/dijetsnodego-apis/health) until it is done bootstrapping.

Improvements that reduce the amount of time it takes to bootstrap are under development.

The bottleneck during bootstrapping is typically database IO. Using a more powerful CPU or increasing the database IOPS on the computer running a node will decrease the amount of time bootstrapping takes.

## Generating Code

Dijets Node Binary uses multiple tools to generate efficient and boilerplate code.

### Running protobuf codegen

To regenerate the protobuf go code, run `scripts/protobuf_codegen.sh` from the root of the repo.

This should only be necessary when upgrading protobuf versions or modifying .proto definition files.

To use this script, you must have [buf](https://docs.buf.build/installation) (v1.9.0), protoc-gen-go (v1.28.0) and protoc-gen-go-grpc (v1.2.0) installed.

To install the buf dependencies:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

If you have not already, you may need to add `$GOPATH/bin` to your `$PATH`:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

If you extract buf to ~/software/buf/bin, the following should work:

```sh
export PATH=$PATH:~/software/buf/bin/:~/go/bin
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/protobuf/cmd/protoc-gen-go-grpc
scripts/protobuf_codegen.sh
```

For more information, refer to the [GRPC Golang Quick Start Guide](https://grpc.io/docs/languages/go/quickstart/).

### Running protobuf codegen from docker

```sh
docker build -t dijetsnode:protobuf_codegen -f api/Dockerfile.buf .
docker run -t -i -v $(pwd):/opt/dijetsnode -w/opt/dijetsnode dijetsnode:protobuf_codegen bash -c "scripts/protobuf_codegen.sh"
```

### Running mock codegen

To regenerate the [gomock](https://github.com/golang/mock) code, run `scripts/mock.gen.sh` from the root of the repo.

This should only be necessary when modifying exported interfaces or after modifying `scripts/mock.mockgen.txt`.

## Versioning

### Library Compatibility Guarantees

The release version for each DijetsNodeGo binary is essentially also the version of the network itself. It is expected that interfaces exported by DijetsNodeGo's packages may change in `Patch` version updates.

### API Compatibility Guarantees

APIs exposed when running DijetsNodeGo will maintain backwards compatibility, unless the functionality is explicitly deprecated and announced when removed.

## Supported Platforms

DijetsNodeGo can run on different platforms, with different levels of stress testing achieved through its development:

The following table lists currently supported platforms:

| Architecture | Operating system |
| :----------: | :--------------: | 
|    amd64     |      Linux       |    
|    arm64     |      Linux       |     
|    amd64     |      Darwin      |   
|    amd64     |     Windows      |  (Windows OS is not yet qualified as fully stress tested)

DijetsNodeGo is a tweaked fork of AvalancheGo which maintains upstream changes.

## Bugs / Vulnerabilities
--

### Versioning

DijetsNodeGo is first and foremost a client for the Dijets network. The versioning of DijetsNodeGo follows that of the Dijets network.

- `v0.x.x` indicates a development network version.
- `v1.x.x` indicates a production network version.
- `vx.[Upgrade].x` indicates the number of network upgrades that have occurred.
- `vx.x.[Patch]` indicates the number of client upgrades that have occurred since the last network upgrade.