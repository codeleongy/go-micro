module github.com/micro/go-micro/v2

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/bwmarrin/discordgo v0.20.2
	github.com/caddyserver/certmagic v0.10.6
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ef-ds/deque v1.0.4-0.20190904040645-54cb57c252a1
	github.com/evanphx/json-patch/v5 v5.0.0
	github.com/forestgiant/sliceutil v0.0.0-20160425183142-94783f95db6c
	github.com/fsnotify/fsnotify v1.4.7
	github.com/fsouza/go-dockerclient v1.6.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-acme/lego/v3 v3.4.0
	github.com/go-git/go-git/v5 v5.1.0
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee
	github.com/gobwas/pool v0.2.0 // indirect
	github.com/gobwas/ws v1.0.3
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/hashicorp/hcl v1.0.0
	github.com/hpcloud/tail v1.0.0
	github.com/imdario/mergo v0.3.9
	github.com/kr/pretty v0.2.0
	github.com/lib/pq v1.3.0
	github.com/lucas-clemente/quic-go v0.14.1
	github.com/micro/cli/v2 v2.1.2
	github.com/miekg/dns v1.1.27
	github.com/nats-io/nats-server/v2 v2.1.6 // indirect
	github.com/nats-io/nats.go v1.9.2
	github.com/nlopes/slack v0.6.1-0.20191106133607-d06c2a2b3249
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.1
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	go.etcd.io/bbolt v1.3.4
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa
	golang.org/x/net v0.5.0
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
	google.golang.org/grpc v1.52.0
	google.golang.org/grpc/examples v0.0.0-20230201212035-3151e834fa25
	google.golang.org/protobuf v1.28.1
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
	gopkg.in/telegram-bot-api.v4 v4.6.4
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.8
	google.golang.org/grpc => google.golang.org/grpc v1.49.0
)
