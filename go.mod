module nlp_worker

go 1.15

replace binhct/common => ../common

require (
	binhct/common v0.0.0-00010101000000-000000000000
	github.com/rs/xid v1.2.1
	github.com/spf13/viper v1.7.1
	go.mongodb.org/mongo-driver v1.4.2
)
