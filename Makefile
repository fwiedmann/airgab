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

all: install_cronjob

prepare:
        chmod +x airgab.sh
        mkdir -pv $(DESTIANTION_DATA_DIRECTORY)

test: prepare
        ./airgab.sh  $(SOURCE_DATA_DIRCECTORY) $(DESTIANTION_DATA_DIRECTORY) $(RESYNC_TEST_OPTIONS)

run: test
        ./airgab.sh  $(SOURCE_DATA_DIRCECTORY) $(DESTIANTION_DATA_DIRECTORY) $(RESYNC_OPTIONS)

install_cronjob: run
        echo "$(CRONJOB_SCHEDULE) $$PWD/airgab.sh $(SOURCE_DATA_DIRCECTORY) $(DESTIANTION_DATA_DIRECTORY) $(RESYNC_OPTIONS)" > $$PWD/temp_cronjob
        crontab -u "$$LOGNAME" temp_cronjob

