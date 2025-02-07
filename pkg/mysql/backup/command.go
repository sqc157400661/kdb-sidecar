package backup

type XtraBackupCommand struct {
}

func NewXtraBackupCommand() *XtraBackupCommand {
	return &XtraBackupCommand{}
}

func (c *XtraBackupCommand) Backup() {

}

func (c *XtraBackupCommand) Restore() {

}
