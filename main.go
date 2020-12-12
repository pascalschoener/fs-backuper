package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

//doBackup backups a fs-dir in several modes
func doBackup(sourceDir string, targetDir string, mode int) error {

	var cmd = exec.Command("echo", "nothing to do")

	switch {
	case mode == 1:
		log.Info("BACKUP_MODE=1 Simple copy")
		cmd = exec.Command("/bin/sh", "-c", "cp -r "+sourceDir+" "+targetDir+"/backup_"+time.Now().Format("20060102150405")+"")

	case mode == 2:
		log.Info("BACKUP_MODE=2 tar without compression")
		cmd = exec.Command("/bin/sh", "-c", "tar -cf "+targetDir+"/backup_"+time.Now().Format("20060102150405")+".tar "+sourceDir)
	case mode == 3:
		log.Info("BACKUP_MODE=3 tar with gzip compression")
		cmd = exec.Command("/bin/sh", "-c", "tar -czf "+targetDir+"/backup_"+time.Now().Format("20060102150405")+".tar.gz "+sourceDir)
	}

	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func main() {

	source := os.Getenv("SOURCE_DIR")
	target := os.Getenv("TARGET_DIR")
	mode := os.Getenv("BACKUP_MODE")

	if source == "" || target == "" || mode == "" {
		log.Errorf("Error while getting needed Env-Variables: source=" + source + " target=" + target + " mode=" + mode)
	}

	modeInt, err := strconv.Atoi(mode)

	if err != nil {
		log.Error(err)
	}

	for {
		var retryErr error

		for i := 0; i < 3; i++ {
			err := doBackup(source, target, modeInt)
			if err == nil {
				break
			}
			retryErr = err
			log.Warnf("Error while processing Backup (Retry "+fmt.Sprint(i+1)+" of 3): %s", err)
			time.Sleep(3 * time.Second)
		}

		if retryErr != nil {
			log.Errorf("No backup was done after 3 retries because of Error: %s", retryErr)
		}

		if retryErr == nil {
			log.Info("Backup Sucessfully. Sleep while waiting for next run in 24 hour")
		}
		time.Sleep(24 * time.Hour)
	}

}
