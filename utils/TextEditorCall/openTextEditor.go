package TextEditorCall

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func TextEdit(vi string) string {
	tmpDir := os.TempDir()
	tmpFile, tmpFileErr := os.CreateTemp(tmpDir, "sss-EDITING-SCRIPT")
	if tmpFileErr != nil {
		log.Errorf("Error: %s while creating tempFile\n", tmpFileErr)
		return ""
	}
	defer os.Remove(tmpFile.Name())
	path, err := exec.LookPath(vi)
	if err != nil {
		log.Errorf("Error: %s while looking up for %s!!\n", path, vi)
		return ""
	}
	log.Debugln("Open File", tmpFile.Name())
	cmd := exec.Command(path, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Errorf("Error: Start failed: %s\n", err)
		return ""
	}
	log.Infoln("Waiting for command to finish.")
	err = cmd.Wait()
	log.Debugln("Complete Editing")
	if err != nil {
		log.Errorf("Command finished with error: %v\n", err)
		return ""
	}
	defer tmpFile.Close()
	var readByte []byte
	readByte, err = os.ReadFile(tmpFile.Name())
	if err != nil {
		log.Errorf("Failed to get user input: %v \n", err)
	}
	log.Debugln("read file data from editor", readByte)
	return string(readByte)
}
