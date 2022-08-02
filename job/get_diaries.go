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
	GetDiariesUrls    *service.GetDiariesFromCouncilService
	Redis             redis.Cmdable
	InitialDateFilter string
	FinalDateFilter   string
}

const LayoutDate = "2006-01-02"

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
		log.Errorf("GetDiaries - Lock: %s", err.Error())
	}

	defer func(lock *redislock.Lock, ctx context.Context) {
		err := lock.Release(ctx)
		if err != nil {
			log.Errorf("[GetDiaries] - Erro ao desbloquear chave no Redis: %s", err.Error())
		}
	}(lock, ctx)

	gd.dateFilterPeriod()
	err = gd.GetDiariesUrls.ProcessDiariesJSON(gd.buildParams())
	if err != nil {
		log.Errorf("[GetDiariesFromCouncilService] - Ocorreu um erro ao buscar url's dos diários - %s", err.Error())
	}
}

func (gd *GetDiaries) dateFilterPeriod() {
	if len(gd.InitialDateFilter) == 0 && len(gd.FinalDateFilter) == 0 {
		initialDate := time.Now().AddDate(0, 0, -1)
		finalDate := time.Now()
		gd.InitialDateFilter = initialDate.Format(LayoutDate)
		gd.FinalDateFilter = finalDate.Format(LayoutDate)
	}
}

func (gd *GetDiaries) buildParams() string {
	return "?tipoSessao=1,4,5&dataInicio=" + gd.InitialDateFilter + "&dataFinal=" + gd.FinalDateFilter
}
