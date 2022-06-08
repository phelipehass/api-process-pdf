build-mocks:
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen
	~/go/bin/mockgen -source=extract_data_from_pdf/service.go -destination=extract_data_from_pdf/mock/service.go -package=mock

dependencies:
	go mod tidy