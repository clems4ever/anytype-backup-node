# AnyType Dockerized Backup Node

[AnyType](https://anytype.io/) is this incredible note-taking app that was released in the
opensource.

This project aims to spawn a AnyType self-hosted dockerized backup node for keeping your personal data private on
your own devices and network. Note that keeping the data on AnyType servers is fine in terms of security and
privacy because it is end-to-end encrypted. However, having a self-hosted backup node provides an extra layer of
security as well as a bigger space than 1GB for storing files but at the **non-negligible** cost of maintaining
the infrastructure.

Please note that, as easy as it is to spawn a backup node with this project, it is not the only thing to do
to end up with a fully working self-hosted setup including clients. I would highly advise not to use this project if
you do not have a deep understanding of the internals of anytype because fixing and upgrading this infrastructure is
not an easy thing. Moreover, you will still need to build the clients for your devices, which is definitely not a
no-brainer.

*This project is my own independant contribution to the Anytype project, therefore it is not maintained by the
Anytype team.*

## Disclaimer of Liability

Please be advised that by using this project, you agree to do so at your own risk. While every effort has been
made to ensure the integrity and security of the code, I cannot be held liable for any damage, loss, or
corruption of data that may occur as a result of using this project. It is strongly recommended that you back
up all important data and thoroughly review the code and documentation before implementation. By proceeding with
the use of this project, you acknowledge and accept full responsibility for any and all potential consequences.

## Getting Started

With the Liability disclaimer out of the way, let's first make sure Docker and docker-compose are installed on your
machine before anything else.

Then, download the latest release of [anytype-backup-node](https://github.com/clems4ever/anytype-backup-node/releases).
On Linux and Mac, make sure you make the binary executable with

```bash
chmod +x anytype-backup-node_Linux_x86_64
```

You can now run the following command to generate a [default configuration](./internal/backupnode/config.yml) that you
can edit before spawning the infrastructure.

```bash
# generate the default config.yml file.
./anytype-backup-node_Linux_x86_64 init
```

Once the configuration file is ready, run the following command to bootstrap the services of the backup node.

```bash
# spawn the services.
./anytype-backup-node_Linux_x86_64 bootstrap
```

Now, you should have all services running. However, as documented above you need to build the clients for your devices with
the proper anytype configuration. The anytype configuration files including `heart.yml` are accessible for your to review in
`configurations/` unless you modified the path in your configuration.

This project does not build the clients because the Anytype team is working on a way to  [customize the clients without rebuilding
them](https://github.com/orgs/anyproto/discussions/17#discussioncomment-6651683) and it will be released soon enough.

## Where is the Dockerfile and docker-compose.yml?

The files are embedded in the go application that you can download from the
[release tab](https://github.com/clems4ever/anytype-backup-node/releases) but if you want you can check the source of the
[Dockerfile](./internal/backupnode/Dockerfile) and the [docker-compose.yml](./internal/backupnode/docker-compose.yml) file.

## Contributions and testing

This project has been tested on a Linux server. I've manually verified that the space is synced and that files are
stored in the file node by checking the content of MinIO and Redis.

Those tests are not automated yet so if you are willing to contribute, make sure you fully test
your changes yourself and prove it with some automated tests or at least some screenshots.

## Next

- Authenticate redis

## Sponsorship

You can thank me by sponsoring me on [Github](https://github.com/sponsors/clems4ever) 

or

You can buy me a beer in the crypto sphere on Ethereum or Polygon at [0x92a9C9e6a418712252fB5996CfE3391a7baBBffB](https://etherscan.io/address/0x92a9C9e6a418712252fB5996CfE3391a7baBBffB).

## License

This project is licensed under the [MIT license](./LICENSE).
