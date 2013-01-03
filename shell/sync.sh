#!/usr/bin/env sh
rsync -ravz --exclude=.hg  wwwuser@hostname:/data/backups /data/backups
