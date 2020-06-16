module github.com/xuyawen/go-gin-example

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.13
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/middleware  => ./middleware
	github.com/EDDYCJY/go-gin-example/models      => ./models
	github.com/EDDYCJY/go-gin-example/pkg/conf    => ./pkg/conf
	github.com/EDDYCJY/go-gin-example/pkg/e       => ./pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/setting => ./pkg/setting
	github.com/EDDYCJY/go-gin-example/routers     => ./routers
	github.com/EDDYCJY/go-gin-example/util        => ./util
)
