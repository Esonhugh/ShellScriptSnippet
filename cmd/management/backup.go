package management

import (
	"github.com/Esonhugh/ShellScriptSnippet/cmd"
	"github.com/Esonhugh/ShellScriptSnippet/utils/Ask"
	"github.com/Esonhugh/ShellScriptSnippet/utils/Database"
	"github.com/Esonhugh/ShellScriptSnippet/utils/File"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var BackupRecoverCmdOpts struct {
	BackupPath  string
	RecoverMode bool
}

func init() {
	BackupCmd.Flags().StringVarP(&BackupRecoverCmdOpts.BackupPath, "backup-path", "p", "", "Backup path")
	BackupCmd.Flags().BoolVarP(&BackupRecoverCmdOpts.RecoverMode, "recover", "r", false, "Recover mode")
	cmd.RootCmd.AddCommand(BackupCmd)
}

var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup or recover snippets",
	Run: func(cmd *cobra.Command, args []string) {
		realbackupPath := BackupRecoverCmdOpts.BackupPath
		if BackupRecoverCmdOpts.RecoverMode {
			// Recover
			if BackupRecoverCmdOpts.BackupPath == "" {
				backupFp := Ask.ForName("Where is the backup path?")
				if !File.PathExists(backupFp) {
					log.Errorf("Can not found %s", backupFp)
					return
				}
				realbackupPath = backupFp
			}
			Database.Close()
			err := File.Copy(realbackupPath, Database.DBLocation)
			if err != nil {
				log.Errorf("Copy file error: %v", err)
				return
			}
			log.Infof("Recover DB from %s to %s success ", realbackupPath, Database.DBLocation)
		} else {
			// Backup
			if BackupRecoverCmdOpts.BackupPath == "" {
				backupFp := Ask.ForName("Where is the backup path?")
				if File.PathExists(backupFp) {
					log.Errorf("Backup path %s is exists", backupFp)
					return
				}
				realbackupPath = backupFp
			}
			err := File.Copy(Database.DBLocation, realbackupPath)
			if err != nil {
				log.Errorf("Copy file error: %v", err)
				return
			}
			log.Infof("Backup DB from %s to %s success ", Database.DBLocation, realbackupPath)
		}
	},
}
