# ConoHa builder for Packer

A builder plugin of Packer for building images with ConoHa.  
This plugin is a forked version of [OpenStack Builder](https://www.packer.io/docs/builders/openstack.html) of Packer.

## Install

Download the binary from the [Releases](https://github.com/summerwind/packer-builder-conoha/releases) page and place it in one of the following places:

- The directory where packer is, or the executable directory
- `~/.packer.d/plugins` on Unix systems or `%APPDATA%/packer.d/plugins` on Windows
- The current working directory

## Build plugin

You can build this plugin with the following command. Please install [Glide](https://github.com/Masterminds/glide) before the build.

```
$ make
```

## Configuration Reference

The same configuration option can be used as the original OpenStack Builder. An example configuration file is in the `example` directory.

## License

Mozilla Public License 2.0
