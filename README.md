![Image of gopher](images/appenginegophercolor.jpg)

# airgab
Backup your data  from a remote-host periodically with airgab. It comes with a Prometheus endpoint for monitoring. Under the hood airgab runs the backuptask via rsync. It is packaged in a Dockerfile to bundle all dependencies and make it easier to deploy.

## Features

### Backup

- backup files or directory's
- useability of all rsync options

### Prometheus endpoint metrics

Defualt Prometheus port is 9100.

- airgab_counter - overall backup counter
- airgab_last_success - timestamp of last backup
- airgab_current_backup_size_megabyte - size of last backup in megabyte
- airgab_backup_duration_in_minutes - duration of last successful backup

## Requirements

- Docker
- private ssh key of remote host (id_rsa)

## Deploy

Option   |  Description | Required
 ---     |  ---         | ---
 user    |  User to login via ssh to remote host | true
 host    | Remote host, the source of the backup  | true
 options | Rsync options | true
 interval | interval of backup execution in the duration format (examples: 10s, 5m, 15h)| true
 destination | Destination of Backup, If you use the Dockerfile default set to "/home/pilot/backup" | false
 private-key | Location of private ssh-key. If you use the Dockerfile this is set to "/home/pilot/.ssh/id_rsa" | false

 ### Example
 ``` bash

 	docker run -d \
                -v /foo/storage/foo-backup:/home/pilot/backup \
                -v $HOME/.ssh/id_rsa:/home/pilot/.id_rsa \
                -p 9100:9100 \
                airgab \
                --user=superuser \
                --host=192.168.2.XXX \
                --source=/opt/ghost \
                --options=-a \
                --interval=10s
 ```

## Development

To build airgab from source use the Makefile

``` bash
# build and run
make

# cleanup local development environment
make cleanup
```