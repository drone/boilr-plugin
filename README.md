# boilr-plugin

This is a [boilr template](https://github.com/tmrts/boilr) for creating a pipeline plugin. Use a pipeline plugin to create and share re-usable pipeline steps. Get started by installing the template:

```console
$ boilr download drone/boilr-plugin drone-plugin
```

create a project in directory my-plugin:

```console
$ boilr template use drone-plugin my-plugin
```

enter the docker registry name for this project:

```text
[?] Please choose a value for "DockerRepository" [default: "owner/name"]:
```

enter the go module name:

```text
[?] Please choose a value for "GoModule" [default: "github.com/owner/name":
```
