package backup

import (
	"backup-automation/logs"
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func Backup(sourcePath, destinationPath string) {
	sourceFiles := getNameFilesFromSourcePath(sourcePath)
	destinationFiles := getNameFilesFromSourcePath(destinationPath)

	backupFiles(sourcePath, destinationPath, sourceFiles, destinationFiles)
}

func checkBackUp(sourceFile, destinationFile string) error {
	backupData, err := os.ReadFile(destinationFile)
	if err != nil {
		return err
	}

	sourceData, err := os.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	if !bytes.Equal(backupData, sourceData) {
		return errors.New("An error Occured while backing up the file:")
	}

	return nil
}

func checkModifiedFiles(sourceFiles, destinationFiles []os.DirEntry) []os.DirEntry {
	var files []os.DirEntry
	for _, sFile := range sourceFiles {
		for _, dFile := range destinationFiles {
			if sFile.Name() == dFile.Name() {
				sInfo, err := sFile.Info()
				if err != nil {
					logs.SaveLogs(logrus.Fatal, fmt.Sprintf("%s", err))
				}
				dInfo, err := dFile.Info()
				if err != nil {
					logs.SaveLogs(logrus.Fatal, fmt.Sprintf("%s", err))
				}

				if sInfo.ModTime().Unix() > dInfo.ModTime().Unix() {
					files = append(files, sFile)
					logs.SaveLogs(logrus.Info, fmt.Sprintf("Modified file found: %s", sFile.Name()))
				}
			}
		}
	}

	if len(files) == 0 {
		return sourceFiles
	}

	return files
}

func backupFiles(sourcePath, destinationPath string, sourceFiles, destinationFiles []os.DirEntry) {

	sourceFiles = checkModifiedFiles(sourceFiles, destinationFiles)

	for _, file := range sourceFiles {
		if file.IsDir() {
			source := sourcePath + "/" + file.Name() + "/"
			destination := destinationPath + "/" + file.Name() + "/"
			err := os.Mkdir(destination, 0750)
			if err != nil && os.IsNotExist(err) {
				logs.SaveLogs(logrus.Fatal, fmt.Sprintf("%s %s", err, file.Name()))
			}

			backupFiles(source, destination, getNameFilesFromSourcePath(source), getNameFilesFromSourcePath(destination))
			continue
		}

		sourceFile := fmt.Sprintf("%s/%s", sourcePath, file.Name())
		destinationFile := fmt.Sprintf("%s/%s", destinationPath, file.Name())
		data, err := os.ReadFile(sourceFile)
		if err != nil {
			logs.SaveLogs(logrus.Error, fmt.Sprintf("Unable to read file content: %s", file.Name()))
			continue
		}

		err = os.WriteFile(destinationFile, data, 0644)
		if err != nil {
			logs.SaveLogs(logrus.Error, fmt.Sprintf("Unable to backup file: %s", file.Name()))
			continue
		}

		logrus.Info("Checking file...")
		if err = checkBackUp(sourceFile, destinationFile); err != nil {
			logs.SaveLogs(logrus.Error, fmt.Sprintf("%s", err))
			continue
		}
		logrus.Info(fmt.Sprintf("File (%s) verified successfully", file.Name()))
	}
}

func getNameFilesFromSourcePath(sourcePath string) []os.DirEntry {
	nameFiles, err := os.ReadDir(sourcePath)
	if err != nil {
		logs.SaveLogs(logrus.Fatal, fmt.Sprintf("%s", err))
	}
	return nameFiles
}
