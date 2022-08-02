package repository

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type RepositoryRest struct {
	URL        string
	ClientRest *resty.Client
}

func NewRepositoryRest(URL, session string) *RepositoryRest {
	clientRest := resty.New()
	clientRest.SetBaseURL(URL).
		SetHeader("Cookie", session)

	return &RepositoryRest{
		URL:        URL,
		ClientRest: clientRest,
	}
}

func (r *RepositoryRest) GetDiaries(params map[string]string) ([]byte, error) {
	resp, err := r.ClientRest.
		SetTimeout(120 * time.Second).
		R().
		Get("fusion/services/custom/cvj/consulta/getsessoes?tipoSessao=1,4,5&dataInicio=2022-05-01&dataFinal=2022-06-13")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		switch resp.StatusCode() {
		case 400:
			return nil, fmt.Errorf("Erro ao buscar arquivo - httpCode %d, body: %s", resp.StatusCode(), resp.Body())
		case 401:
			return nil, fmt.Errorf("Sessão está incorreta - httpCode %d, body: %s", resp.StatusCode(), resp.Body())
		case 403:
			return nil, fmt.Errorf("Sessão não autorizada - httpCode %d, body %s", resp.StatusCode(), resp.Body())
		case 404:
			return nil, fmt.Errorf("Diários não encontrados - httpCode %d, body %s", resp.StatusCode(), resp.Body())
		}
	}

	return resp.Body(), nil
}
