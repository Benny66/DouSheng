// @Author: Ciusyan 2023/2/8
package impl_test

import (
	"context"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	// 驱动加载所有需要放入IOC的实例
	_ "github.com/Go-To-Byte/DouSheng/video_service/apps/all"
)

var (
	service video.ServiceServer
)

func TestPublishVideo(t *testing.T) {
	should := assert.New(t)

	req := video.NewPublishVideoRequest()
	req.Token = "xxx"
	req.Title = "sss"

	_, err := service.PublishVideo(context.Background(), req)

	if should.NoError(err) {
		t.Log("保存成功")
	}
}

func init() {

	// 加载配置文件
	if err := conf.LoadConfigFromToml("../../../etc/config.toml"); err != nil {
		panic(err)
	}

	// 初始化全局Logger
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}

	// 初始化IOC容器
	if err := ioc.InitAllDependencies(); err != nil {
		panic(err)
	}

	// 从IOC中获取接口实现
	service = ioc.GetGrpcDependency(video.AppName).(video.ServiceServer)
}
