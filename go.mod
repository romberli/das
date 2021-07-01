module github.com/romberli/das

go 1.16

replace (
	github.com/ClickHouse/clickhouse-go v1.4.3 => github.com/romberli/clickhouse-go v1.4.4-0.20210422094559-b05fc8c4dbe9
	github.com/go-mysql-org/go-mysql v1.1.2 => github.com/romberli/go-mysql v1.1.3-0.20210425112236-f28342fcf6bc
	github.com/spf13/pflag v1.0.5 => github.com/romberli/pflag v1.0.6-alpha
)

require (
	github.com/XiaoMi/soar v0.8.2-0.20210526115350-9a47ddc1089f
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/asaskevich/govalidator v0.0.0-20200819183940-29e1ff8eb0bb
	github.com/gin-gonic/gin v1.6.3
	github.com/hashicorp/go-multierror v1.1.0
	github.com/jinzhu/now v1.1.2
	github.com/romberli/go-util v0.3.9-0.20210621075221-c920fb17a852
	github.com/romberli/log v1.0.20
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	go.uber.org/zap v1.16.0
)
