# Logmetro

__NOTE:__ still in development, this README describes the first goal of Logmetro. No implementation yet.

## Features

- Monitor log stream (such as http access logs) in realtime
  - requests per second, average response time, counting status codes...
- Awesome dashboard in browser
- Useful HTTP API

## What we don't

- higher accuracy
  - data might not be perfect.
- persist data
  - export to other datastore using Output feature.
- complex event stream processing
  - We do only simple aggregation (count something...)
  - Consider using https://github.com/norikra/norikra for such case

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
