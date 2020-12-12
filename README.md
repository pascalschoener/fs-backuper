# fs-backuper

This is a simple golang based FS-Backuper with 3 different backup-modes

# Env
|Variable|description|
|---|---|
|BACKUP_MODE| Defines the mode of Backup <br>BACKUP_MODE=1 only copy of data from source to target dir <br>BACKUP_MODE=2 create a uncompressed tar of the source dir to the target dir<br>BACKUP_MODE=3 create a compressed tar.gz of the source dor to the target dir|
|SOURCE_DIR| Defines the source directory you want backup (be shure that its mounted correctly to your container)|
|TARGET_DIR| Defines the target directory you want backup (be shure that its mounted correctly to your container)|


# docker-compose
``` yaml

```