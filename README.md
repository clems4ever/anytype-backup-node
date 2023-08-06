# AnyType Dockerized Backup Node

AnyType is this incredible note taking app that was released in the opensource.
This repository will help you spawn a self-hosted backup node for keeping your data
private on your own machines. Keeping them on AnyType servers is fine since they are
encrypted, but having your own backup node provides a second layer of security in case
a vulnerability in the encryption done by AnyType.

## Getting Started

First make sure Docker and docker-compose are installed on your machine before anything
else.

Then, download the latest release of [anytype-backup-node](https://github.com/clems4ever/anytype-backup-node/releases).
On Linux and Mac, make sure you make the binary executable.

You can now run the following command:
```
./anytype-backup-node_Linux_x86_64 bootstrap -c config.yml
```

At this point, a configuration for your setup has been generated and used for
creating an ecosystem of services using docker-compose.

## License

This project is licensed under the [MIT license](./LICENSE).