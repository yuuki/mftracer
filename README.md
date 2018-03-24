# mftracer

[license]: https://github.com/yuuki/lstf/blob/master/LICENSE

## Architecture Overview

```
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


## License

[MIT][license] License.

## Author

[yuuki](https://github.com/yuuki)