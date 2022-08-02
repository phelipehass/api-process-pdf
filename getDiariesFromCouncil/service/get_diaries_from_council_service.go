package service

import (
	"api/config"
	"api/getDiariesFromCouncil/repository"
	"api/models"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"strconv"
)

type GetDiariesFromCouncilService struct {
	UrlBase         string
	DiaryRepository *repository.PostgresRepository
	RepositoryRest  *repository.RepositoryRest
}

func NewService(diaryRepo *repository.PostgresRepository, repoRest *repository.RepositoryRest) *GetDiariesFromCouncilService {
	return &GetDiariesFromCouncilService{
		UrlBase:         config.GetURLConsultDiaries(),
		DiaryRepository: diaryRepo,
		RepositoryRest:  repoRest,
	}
}

func (s *GetDiariesFromCouncilService) ProcessDiariesJSON(pathParams string) error {
	body, err := s.RepositoryRest.GetDiaries(pathParams)
	if err != nil {
		return err
	}

	processData, err := parseJSONDiaries(body)
	if err != nil {
		return err
	}

	diaryModels := s.processDiariesData(processData)
	errors := s.saveDiaryModels(diaryModels)
	if errors > 0 {
		return fmt.Errorf("Ocorreu erro ao salvar %d diários", errors)
	}

	return nil
}

func parseJSONDiaries(diariesJSON []byte) (*models.ProcessData, error) {
	var processData models.ProcessData
	if err := json.Unmarshal(diariesJSON, &processData); err != nil {
		return nil, err
	}

	return &processData, nil
}

func (s *GetDiariesFromCouncilService) processDiariesData(processData *models.ProcessData) []models.Diary {
	diaryModels := []models.Diary{}
	diaryModel := models.Diary{}

	for _, diary := range processData.Diaries {
		diaryModel = models.Diary{
			UrlArchive: s.UrlBase + "fusion/services/CVJ/customService/downloadPDF/" + strconv.FormatInt(diary.Diary, 10),
			Cod:        diary.Diary,
		}

		diaryModels = append(diaryModels, diaryModel)
	}

	return diaryModels
}

func (s *GetDiariesFromCouncilService) saveDiaryModels(diaryModels []models.Diary) int {
	var errors int

	for _, diary := range diaryModels {
		err := s.DiaryRepository.SaveDiary(&diary)
		if err != nil {
			log.Errorf("[saveDiaryModels] - Erro ao salvar o diário %d - Erro: %s", diary.Cod, err.Error())
			errors += 1
		}
	}

	return errors
}
