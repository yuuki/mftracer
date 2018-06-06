# mftracer

[license]: https://github.com/yuuki/lstf/blob/master/LICENSE

## Architecture Overview

```ascii
+-----------+
| mftracerd |----------+
+-----------+          | PutItem
                       V
+-----------+       +-----------------+      +-----------------+
| mftracerd |------>| Kinesis Streams | ---> | Lambda function |
+-----------+       +-----------------+      +-----------------+
                       ^                              |
+-----------+          |                              | INSERT or UPDATE
| mftracerd |----------+                              V
+-----------+           +----------+  SELECT +-----------------+
                        | mftracer |<--------| Postgres Aurora |
                        +----------+    ^    +-----------------+
                                        |
                                        | API
                                   +----------+
                                   | Mackerel |
                                   +----------+
```

## Usage

```shell
mftctl --level 2 --dest-ipv4 10.0.0.21
10.0.0.21
└<-- 10.0.0.22:many (connections:30)
└<-- 10.0.0.23:many (connections:30)
└<-- 10.0.0.24:many (connections:30)
	└<-- 10.0.0.30:many (connections:1)
	└<-- 10.0.0.31:many (connections:1)
└<-- 10.0.0.25:many (connections:30)
...
```

```shell
mftctl --dest-service blog --dest-roles redis --dest-roles memcached
blog:redis
└<-- 10.0.0.22:many (connections:30)
└<-- 10.0.0.23:many (connections:30)
└<-- 10.0.0.24:many (connections:30)
	└<-- 10.0.0.30:many (connections:1)
	└<-- 10.0.0.31:many (connections:1)
└<-- 10.0.0.25:many (connections:30)
blog:memcached
└<-- 10.0.0.23:many (connections:30)
└<-- 10.0.0.25:many (connections:30)
...
```

## License

[MIT][license] License.

## Author

[yuuki](https://github.com/yuuki)