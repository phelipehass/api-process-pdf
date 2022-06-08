package service

import (
	"github.com/tj/assert"
	"os"
	"testing"
)

func indication() *string {
	indication := "6118/2022 - Henrique Deckmann - MDB - " +
		"Notificar o Proprietário do imóvel situado na Rua Marechal Hermes,, " +
		"sobre a inscrição cadastral: 13.20.31.23.0704 32 de propriedade de MATILDES LOPES AREVALO, " +
		"Bairro Glória, para que o mesmo providencia a execução de calçada em seu imóvel. Conforme já " +
		"solicitado através da Indicação No 19054/2021, feita em 9 de Novembro de 2021"

	return &indication
}

func indications() *string {
	indications := "745/2022 - Wilian Tonezi - PATRIOTA - Instalaçao de nova boca de lobo na Rua " +
		"Augusto Borinelli, em frente ao número 104, no bairro Vila Nova. |||750/2022 - Wilian Tonezi - " +
		"PATRIOTA - Operação tapa-buraco na Rua Jaguaruna, nas proximidades do no 13, no Bairro Centro. " +
		"|||760/2022 - Diego Machado - PSDB - Ensaibramento e patrolamento da Rua Alfredo Nielson, em toda " +
		"a sua extensão, no Distrito de Pirabeiraba. |||807/2022 - Pastor Ascendino Batista - PSD - Ensaibramento " +
		"e patrolamento da Rua Sculptor , em toda a sua extensão, no Bairro Jardim Paraiso. |||733/2022 - " +
		"Kiko do Restaurante - PSD - Refazer boca de lobo na Rua Maria de Lourdes Goularte, nas proximidades " +
		"do no 621, no Bairro Parque Guarani. |||695/2022 - Nado - PROS - Colocação de abrigo de ônibus na Rua " +
		"Guaíra, nas proximidades do no 1560, no Bairro Iririú."

	return &indications
}

func setEnvs() {
	os.Setenv("INIT_EXTRACTION", "INDICAÇÕES")
	os.Setenv("FINISH_EXTRACTION", "MATÉRIA DA ORDEM DO DIA")
	os.Setenv("TITLE", "CÂMARA DE VEREADORES DE JOINVILLE")
	os.Setenv("SUBTITLE", "ESTADO DE SANTA CATARINA")
}

func Test_ProcessData(t *testing.T) {
	setEnvs()
	processDataService := NewService()
	data := processDataService.ProcessData(indication())

	assert.Equal(t, "6118/2022", data[0].NumberIndication)
	assert.Equal(t, "Henrique Deckmann", data[0].NamePersonResponsible)
	assert.Equal(t, "MDB", data[0].Entourage)
	assert.Equal(t, "Notificar o Proprietário do imóvel situado na Rua Marechal Hermes,, sobre a inscrição cadastral: 13.20.31.23.0704 32 de propriedade de MATILDES LOPES AREVALO, Bairro Glória, para que o mesmo providencia a execução de calçada em seu imóvel. Conforme já solicitado através da Indicação No 19054/2021, feita em 9 de Novembro de 2021", data[0].Description)
	assert.Equal(t, "Glória", data[0].District.String)
	assert.Equal(t, "Rua Marechal Hermes", data[0].Street.String)
}

func Test_ProcessData_List(t *testing.T) {
	setEnvs()
	processDataService := NewService()
	data := processDataService.ProcessData(indications())

	assert.Equal(t, "745/2022", data[0].NumberIndication)
	assert.Equal(t, "Wilian Tonezi", data[0].NamePersonResponsible)
	assert.Equal(t, "PATRIOTA", data[0].Entourage)
	assert.Equal(t, "Instalaçao de nova boca de lobo na Rua "+
		"Augusto Borinelli, em frente ao número 104, no bairro Vila Nova.", data[0].Description)
	assert.Equal(t, "Vila Nova", data[0].District.String)
	assert.Equal(t, "Rua Augusto Borinelli", data[0].Street.String)

	assert.Equal(t, "750/2022", data[1].NumberIndication)
	assert.Equal(t, "Wilian Tonezi", data[1].NamePersonResponsible)
	assert.Equal(t, "PATRIOTA", data[1].Entourage)
	assert.Equal(t, "Operação tapa-buraco na Rua Jaguaruna, nas proximidades do no 13, no Bairro Centro.", data[1].Description)
	assert.Equal(t, "Centro", data[1].District.String)
	assert.Equal(t, "Rua Jaguaruna", data[1].Street.String)

	assert.Equal(t, "760/2022", data[2].NumberIndication)
	assert.Equal(t, "Diego Machado", data[2].NamePersonResponsible)
	assert.Equal(t, "PSDB", data[2].Entourage)
	assert.Equal(t, "Ensaibramento e patrolamento da Rua Alfredo Nielson, em toda "+
		"a sua extensão, no Distrito de Pirabeiraba.", data[2].Description)
	assert.Equal(t, "de Pirabeiraba", data[2].District.String)
	assert.Equal(t, "Rua Alfredo Nielson", data[2].Street.String)

	assert.Equal(t, "807/2022", data[3].NumberIndication)
	assert.Equal(t, "Pastor Ascendino Batista", data[3].NamePersonResponsible)
	assert.Equal(t, "PSD", data[3].Entourage)
	assert.Equal(t, "Ensaibramento "+
		"e patrolamento da Rua Sculptor , em toda a sua extensão, no Bairro Jardim Paraiso.", data[3].Description)
	assert.Equal(t, "Jardim Paraiso", data[3].District.String)
	assert.Equal(t, "Rua Sculptor", data[3].Street.String)

	assert.Equal(t, "733/2022", data[4].NumberIndication)
	assert.Equal(t, "Kiko do Restaurante", data[4].NamePersonResponsible)
	assert.Equal(t, "PSD", data[4].Entourage)
	assert.Equal(t, "Refazer boca de lobo na Rua Maria de Lourdes Goularte, nas proximidades "+
		"do no 621, no Bairro Parque Guarani.", data[4].Description)
	assert.Equal(t, "Parque Guarani", data[4].District.String)
	assert.Equal(t, "Rua Maria de Lourdes Goularte", data[4].Street.String)

	assert.Equal(t, "695/2022", data[5].NumberIndication)
	assert.Equal(t, "Nado", data[5].NamePersonResponsible)
	assert.Equal(t, "PROS", data[5].Entourage)
	assert.Equal(t, "Colocação de abrigo de ônibus na Rua "+
		"Guaíra, nas proximidades do no 1560, no Bairro Iririú.", data[5].Description)
	assert.Equal(t, "Iririú", data[5].District.String)
	assert.Equal(t, "Rua Guaíra", data[5].Street.String)
}

func Test_ProcessData_Empty(t *testing.T) {
	setEnvs()
	processDataService := NewService()
	values := ""
	data := processDataService.ProcessData(&values)

	assert.Empty(t, data)
}
