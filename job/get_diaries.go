package job

import (
	"api/getDiariesFromCouncil/service"
	"context"
	"github.com/apex/log"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"runtime/debug"
	"time"
)

type GetDiaries struct {
	GetDiariesUrls *service.GetDiariesFromCouncilService
	Redis          redis.Cmdable
}

func (gd *GetDiaries) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("[GetDiariesFromCouncilService] - Recuperando do panico no job - erro: %v - stacktrace: %v", r, string(debug.Stack()))
		}
	}()

	ctx := context.Background()
	locker := redislock.New(gd.Redis)
	lock, err := locker.Obtain(ctx, "GetDiaries", 15*time.Minute, nil)
	if err == redislock.ErrNotObtained {
		return
	} else if err != nil {
		log.WithError(err).Error("GetDiaries - Lock")
	}

	defer lock.Release(ctx)

	err = gd.GetDiariesUrls.ProcessDiariesJSON()
	if err != nil {
		log.Errorf("[GetDiariesFromCouncilService] - Ocorreu um erro ao buscar url's dos diários - %s", err.Error())
	}
}
