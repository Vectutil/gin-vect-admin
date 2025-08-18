package job

import (
	"gin-vect-admin/internal/config"
	"gin-vect-admin/pkg/logger"
	"gin-vect-admin/pkg/mysql"
	"testing"
)

func Test(t *testing.T) {
	xxxinit()
	addExampleJob()
}

func xxxinit() {
	config.InitConfig("E:\\workspace\\src\\jz-crawler\\config.yaml")
	mysql.InitMysql()
	logger.InitLogger()
}
