# AnyType Dockerized Backup Node

[AnyType](https://anytype.io/) is this incredible note-taking app that was released in the
opensource.

This project aims to spawn a AnyType self-hosted backup node for keeping your personal data private on your
own devices and network. Note that keeping the data on AnyType servers is fine in terms of security and privacy
because it is end-to-end encrypted. However, having a self-hosted backup node provides an extra layer of
security as well as a bigger space than 1GB for storing files but at the **non-negligible** cost of maintaining
the infrastructure.

Please note that, as easy as it is to spawn a backup node with this project, it is far from the only thing to do
to end up with a fully working self-hosted setup. I highly advise you not to use this project if you do not have
a deep understanding of the internals of anytype. Maintaining, i.e., fixing and upgrading this infrastructure is
not an easy thing. Moreover, you will still need to build the clients for your devices, which is definitely not a
no-brainer.

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

You can now run the following command to generate a default configuration file `config.yml` that you can edit before spawning
the infrastructure.

```bash
./anytype-backup-node_Linux_x86_64 init
```

Once the configuration file is ready, run the following command to bootstrap the services of the backup node.

```bash
./anytype-backup-node_Linux_x86_64 bootstrap
```

Now, you should have all services running but you need to build the applications for your devices with the proper
anytype configuration. This is not yet handled in this repository but one can refer to
https://tech.anytype.io/how-to/self-hosting in order to check how to do it and until it is handled in the repo or
simply made obsolete after some changes in the anytype codebase.

The anytype configuration files are accessible for your to review in `configurations/` unless you modified the path.

## Contributions and testing

This project has been tested on a Linux server. I've manually verified that the space is synced and that files are
stored in the file node by checking the content of MinIO and Redis.

Those tests are not automated yet so if you are willing to contribute, make sure you fully test
your changes yourself and prove it with some automated tests or at least some screenshots.

## Next

- Introduce commands to build the clients from the generated configuration. (Or maybe not if clients can be customized instead in the future).

## License

This project is licensed under the [MIT license](./LICENSE).