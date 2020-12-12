# fs-backuper

This is a simple golang based FS-Backuper with 3 different backup-modes

## Env

|Variable|description|
|---|---|
|BACKUP_MODE| Defines the mode of Backup <br>BACKUP_MODE=1 only copy of data from source to target dir <br>BACKUP_MODE=2 create a uncompressed tar of the source dir to the target dir<br>BACKUP_MODE=3 create a compressed tar.gz of the source dor to the target dir|
|SOURCE_DIR| Defines the source directory you want backup (be shure that its mounted correctly to your container)|
|TARGET_DIR| Defines the target directory you want backup (be shure that its mounted correctly to your container)|
|BACKUP_CLEANUP_DISABLE | default is false, set it to true to disable cleaning up your backups|
|BACKUP_RETENTION| Defines the retention of Backupfiles in days|


<br>
<br>

## docker

You can also use the prebuild docker-containers in https://hub.docker.com/repository/docker/pascalschoener/fs-backuper

## docker-compose
``` yaml
  fs-backuper:
    image: pascalschoener/fs-backuper
    container_name: fs-backuper
    restart: unless-stopped
    volumes:
      - ./backups:/tmp/backups
      - ./source:/tmp/source:ro
    environment:
      - BACKUP_MODE=1
      - SOURCE_DIR=/tmp/source
      - TARGET_DIR=/tmp/backups
```
