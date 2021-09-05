module github.com/netw-device-driver/ndd-runtime

go 1.16

require (
	github.com/go-logr/logr v0.4.0
	github.com/karimra/gnmic v0.18.0
	github.com/netw-device-driver/ndd-grpc v0.1.32
	github.com/openconfig/gnmi v0.0.0-20210707145734-c69a5df04b53
	github.com/pkg/errors v0.9.1
	github.com/spf13/afero v1.6.0
	github.com/yndd/ndd-yang v0.1.62
	golang.org/x/time v0.0.0-20210611083556-38a9dc6acbc6
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
	sigs.k8s.io/controller-runtime v0.9.3
)
