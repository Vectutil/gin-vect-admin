package logic

import (
	"context"
	"gin-vect-admin/internal/app/dao"
	"gin-vect-admin/internal/app/model"
	"gin-vect-admin/internal/app/types"
	"gin-vect-admin/pkg/mysql"
)

type (
	JobMQLogic struct {
		jobMqDao dao.IJobMQDao
	}
	IJobMQLogic interface {
		SearchOne(ctx context.Context, req *types.GetJobMqOneReq) (model *model.JobMq, err error)
	}
)

func NewJobMQLogic() IJobMQLogic {
	return &JobMQLogic{}
}

func (l *JobMQLogic) Create(ctx context.Context, model *model.JobMq) {

}

func (l *JobMQLogic) SearchOne(ctx context.Context, req *types.GetJobMqOneReq) (model *model.JobMq, err error) {
	l.jobMqDao = dao.NewJobMQDao(mysql.GetDB())
	return l.jobMqDao.SearchOne(ctx, req)
}
