# AnyType Dockerized Backup Node

AnyType is this incredible note taking app that was released in the opensource.
This repository will help you spawn a self-hosted backup node for keeping your data
private on your own machines. Keeping them on AnyType servers is fine since they are
encrypted, but having your own backup node provides a second layer of security in case
a vulnerability in the encryption done by AnyType.

## Usage

Make sure Docker and docker-compose are installed on your machine before running the
next command.

```
./start_backup_node
```

At this point, a configuration for your setup has been generated and used for
creating an ecosystem of services using docker-compose.

## License

This project is licensed under the [MIT license](./LICENSE).