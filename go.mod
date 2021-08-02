require (
	cloud.google.com/go/firestore v1.3.0
	github.com/antonholmquist/jason v1.0.0
	github.com/aws/aws-sdk-go v1.34.5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/julienschmidt/httprouter v1.3.0
	github.com/logicful/models v0.0.0-00010101000000-000000000000
	github.com/logicful/service v0.0.0-00010101000000-000000000000
	github.com/robfig/cron/v3 v3.0.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/api v0.29.0
	google.golang.org/grpc v1.30.0

)

module api

go 1.14

replace github.com/logicful/service => ./service

replace github.com/logicful/models => ./models
