# Logmetro

__NOTE:__ still in development, this README describes the initial goals of Logmetro. No implementation yet.

## Features

- Monitor log streams (such as HTTP access logs) in realtime
  - EG: requests per second, average response time, counting status codes...
- Awesome dashboard in browser
- Useful HTTP API

## Things we don't intend to implement

- higher accuracy
  - data might not be perfect.
- persistant data
  - export to other datastores using an Output feature.
- complex event stream processing
  - We do only simple aggregation (count something...)
  - Consider using https://github.com/norikra/norikra for such cases

## Usage

```
$ logmetro -config config.yml
```

## Configuration

here's a plan...

``` yaml
# vim: ft=yaml

inputs:
  - type: tcp
    options:
      listen: [::]:9755
      format: ltsv

web:
  listen: [::]:9754
  cors: '*.local'
  # Specify directory of web ui, if you use awesome web ui
  ui: /path/to/ui/directory

services_grouping:
  - key: server_name
  # - key: some_key
  #   default: _another_template

services:
  _default:
    items:
      reqtime:
        type: numeric
        options:
          key: reqtime
          multiply: 1000
      requests:
        type: count
      rps:
        type: count
        options:
          round: 1s
      status:
        type: valuecount
        options:
          key: status

    # Use time in log?
    time: time
    # or
    # realtime: true

    dashboard:
      type: generic_access_log

    window:
      length: 5s
      retention: 12h

    # monitor in sub groups (e.g. status for each server separately)
    groups:
      - key: server
        filter: proxy
      - key: server
        filter: application
        exclude_from_parent: true

outputs:
  - type: zabbix
  - type: fluentd
  - type: exec
    command: your-favorite-output
```
