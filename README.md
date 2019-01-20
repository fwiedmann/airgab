# airgab

airgab is a wrapper for rsync & crontab to backup directorys from other hosts periodically.

## requirements

To install airgab you need the "make" utility. You can easaly install it via your package-manager. For Ubunut users:

```bash
sudo apt install make
```
## install airgab
Before you can run airgab, you have to configure the variables  in the "Makefile" for your needs.

### 1. open the Makefile with your favourite editor
```bash
vim Makefile
```

### 2. configure the varibales for your needs
```bash
# where your  data will be backuped from (syntax: sshuser@ip:/path)
SOURCE_DATA_DIRCECTORY =

# where your backup will be stored (syntax: /path/foo/bar)
DESTIANTION_DATA_DIRECTORY =

# Default is "-a". "-c" will check increment with checksum and "-P" prints the progress
RESYNC_OPTIONS = -acP

# test the rsync connection with a dry-run
RESYNC_TEST_OPTIONS = --dry-run

#  cron schedulue expression (for every hour: 0 * * * *)
CRONJOB_SCHEDULE = 0 * * * *
```

### 3. run
Prepare the environment of your destination host
```bash
make prepare
```

airgab dry-run
```bash
make test
```

on-time airgab run
```bash
make run
```

install airgab with the cronjob
```bash
make
```

