# WhoIP
`whoip` is a CLI tool to query IP address metadata from ipinfo.io

```
$ whoip
Get IP address metadata from IPInfo.io
Requires at least one IP address to lookup

Usage:
        whoip {-p} {ip-address} {ip-address}..
Options:
        -p      print with a pipe delimiter instead of a table

```

```
$ whoip 8.8.8.8
+---------+---------+------------+---------------+--------------------+---------------------+--------+
|   IP    | COUNTRY |   REGION   |     CITY      |    ORGANIZATION    |      TIMEZONE       | POSTAL |
+---------+---------+------------+---------------+--------------------+---------------------+--------+
| 8.8.8.8 |   US    | California | Mountain View | AS15169 Google LLC | America/Los_Angeles | 94043  |
+---------+---------+------------+---------------+--------------------+---------------------+--------+
```

```
$ whoip 54.239.31.91 104.16.182.15 172.217.12.78
+---------------+---------+------------+----------------+--------------------------+---------------------+--------+
|      IP       | COUNTRY |   REGION   |      CITY      |       ORGANIZATION       |      TIMEZONE       | POSTAL |
+---------------+---------+------------+----------------+--------------------------+---------------------+--------+
| 54.239.31.91  |   US    |  Virginia  | Virginia Beach | AS16509 Amazon.com, Inc. |  America/New_York   | 23452  |
| 172.217.12.78 |   US    | California | Mountain View  |    AS15169 Google LLC    | America/Los_Angeles | 94043  |
| 104.16.182.15 |   US    | California | San Francisco  | AS13335 Cloudflare, Inc. | America/Los_Angeles | 94107  |
+---------------+---------+------------+----------------+--------------------------+---------------------+--------+
```

Use the `-p` flag to output the data with a pipe delimiter instead of as a table

```
$ whoip -p 104.16.182.15
104.16.182.15|US|California|San Francisco|AS13335 Cloudflare, Inc.|America/Los_Angeles|94107
```