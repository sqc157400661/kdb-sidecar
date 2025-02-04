package backup

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

type FullBackupExecutor struct {
	version       string // xtrabackup 版本
	baseDir       string
	xtraBackupBin string
}

// NewXtraBackupCmd 创建一个新的 XtraBackupCmd
func NewXtraBackupCmd(baseDir string) (*FullBackupExecutor, error) {
	if baseDir == "" {
		baseDir = os.Getenv("XTRA_BACKUP_BASE_DIR")
	}
	if baseDir == "" {
		baseDir = "/backup"
	}
	xtraBackupBin := "/usr/bin/xtrabackup"
	// 自动检测 xtrabackup 版本
	version, err := detectXtraBackupVersion(xtraBackupBin)
	if err != nil {
		return nil, err
	}

	return &FullBackupExecutor{
		version:       version,
		baseDir:       baseDir,
		xtraBackupBin: xtraBackupBin,
	}, nil
}

// detectXtraBackupVersion 自动检测 xtrabackup 版本
func detectXtraBackupVersion(xtrabackupPath string) (string, error) {
	// 执行 `xtrabackup --version` 命令来获取版本信息
	cmd := exec.Command(xtrabackupPath, "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get xtrabackup version: %v", err)
	}

	// 解析版本号，假设版本信息格式为: "xtrabackup version 8.0.22-23"
	re := regexp.MustCompile(`xtrabackup version (\d+\.\d+)`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		return "", fmt.Errorf("failed to parse xtrabackup version")
	}

	// 返回版本号，如 "8"
	return matches[1], nil
}

func (e *FullBackupExecutor) Backup() error {
	var backupArgs []string
	backupArgs = append(backupArgs, fmt.Sprintf("--target-dir=%s", e.baseDir))
	backupArgs = append(backupArgs, "--backup")
	if e.version == "2.4" {
		// xtrabackup 2.4 版本命令
		backupArgs = append(backupArgs, "--no-lock") // 2.4 版本不支持 --lock --no-lock
	} else if e.version == "8" {
		// xtrabackup 8 版本命令
		backupArgs = append(backupArgs, "--lock-ddl") // 8 版本使用该参数来锁定 DDL
	}
	// 执行备份命令
	cmdExec := exec.Command(e.xtraBackupBin, backupArgs...)
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	if err := cmdExec.Run(); err != nil {
		return fmt.Errorf("backup failed: %v", err)
	}
	return nil
}

// StreamBackup 是否启用流式备份
func (e *FullBackupExecutor) StreamBackup() error {
	var backupArgs []string
	backupArgs = append(backupArgs, fmt.Sprintf("--target-dir=%s", e.baseDir))
	backupArgs = append(backupArgs, "--stream=xbstream") // 启用流式备份
	if e.version == "2.4" {
		// xtrabackup 2.4 版本命令
		backupArgs = append(backupArgs, "--no-lock") // 2.4 版本不支持 --lock --no-lock
	} else if e.version == "8" {
		// xtrabackup 8 版本命令
		backupArgs = append(backupArgs, "--lock-ddl") // 8 版本使用该参数来锁定 DDL
	}
	// 执行备份命令
	cmdExec := exec.Command(e.xtraBackupBin, backupArgs...)
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	if err := cmdExec.Run(); err != nil {
		return fmt.Errorf("backup failed: %v", err)
	}
	return nil
}

func (e *FullBackupExecutor) Clear() {

}

func (e *FullBackupExecutor) Restore(restoreTime string, binlog string) error {
	var restoreArgs []string
	restoreArgs = append(restoreArgs, "--prepare")

	// 根据 xtrabackup 版本构建命令参数
	if e.version == "2.4" {
		restoreArgs = append(restoreArgs, fmt.Sprintf("--target-dir=%s", e.baseDir))
	} else if e.version == "8" {
		restoreArgs = append(restoreArgs, fmt.Sprintf("--target-dir=%s", e.baseDir))
	}

	// 如果设置了时间点恢复
	if restoreTime != "" {
		restoreArgs = append(restoreArgs, fmt.Sprintf("--apply-log-only"))
		restoreArgs = append(restoreArgs, fmt.Sprintf("--stop-never"))
		restoreArgs = append(restoreArgs, fmt.Sprintf("--restore-time=%s", restoreTime))
	}

	// 如果设置了 binlog 位点恢复
	if binlog != "" {
		restoreArgs = append(restoreArgs, fmt.Sprintf("--start-position=%s", binlog))
	}

	// 执行恢复命令
	restoreCmd := exec.Command(e.xtraBackupBin, restoreArgs...)
	restoreCmd.Stdout = os.Stdout
	restoreCmd.Stderr = os.Stderr
	if err := restoreCmd.Run(); err != nil {
		return fmt.Errorf("restore failed: %v", err)
	}

	// 执行 copy-back 命令，恢复数据到 MySQL 数据目录
	copyBackArgs := []string{"--copy-back", fmt.Sprintf("--target-dir=%s", e.baseDir)}
	copyBackCmd := exec.Command(e.xtraBackupBin, copyBackArgs...)
	copyBackCmd.Stdout = os.Stdout
	copyBackCmd.Stderr = os.Stderr
	if err := copyBackCmd.Run(); err != nil {
		return fmt.Errorf("copy-back failed: %v", err)
	}

	// 恢复完成后启动 MySQL
	// 使用 systemctl 或其他方式启动 MySQL
	fmt.Println("MySQL 恢复完成，启动 MySQL 服务")
	return nil
}
