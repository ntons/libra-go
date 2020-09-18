package sdk

import (
	"flag"
	"fmt"
	"net"
	"path/filepath"
	"sync"

	"github.com/ghodss/yaml"
	log "github.com/ntons/log-go"
	"github.com/ntons/tongo/template"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	sdkpb "github.com/ntons/libra-go/api/sdk/v1"
)

type CommandLineOptions struct {
	// [-c] configuration file path
	ConfigFilePath string
}

func (clopts *CommandLineOptions) SetFlag() {
	flag.StringVar(&clopts.ConfigFilePath, "c", "", "config file path")
}

func loadConfig(fp string) (_ *sdkpb.ServerConfig, err error) {
	b, err := template.RenderFile(fp, nil)
	if err != nil {
		return
	}
	switch ext := filepath.Ext(fp); ext {
	case ".json":
		return
	case ".yml", ".yaml":
		if b, err = yaml.YAMLToJSON(b); err != nil {
			return
		}
	default:
		return nil, fmt.Errorf("unknown file extension: %v", ext)
	}
	cfg := &sdkpb.ServerConfig{}
	if err = protojson.Unmarshal(b, cfg); err != nil {
		return
	}
	return cfg, nil
}

func Serve(factory ServiceFactory) (err error) {
	clopts := &CommandLineOptions{}
	clopts.SetFlag()
	flag.Parse()

	zc := zap.NewProductionConfig()
	zc.Sampling = nil
	zl, err := zc.Build(zap.AddCaller())
	if err != nil {
		return
	}
	log.SetZapLogger(zl)

	cfg, err := loadConfig(clopts.ConfigFilePath)
	if err != nil {
		return
	}
	rawAppCfg, err := protojson.Marshal(cfg.App)
	if err != nil {
	}

	lis, err := net.Listen("tcp", cfg.Bind)
	if err != nil {
		return
	}
	defer lis.Close()

	var wg sync.WaitGroup
	defer wg.Wait()

	srv := NewServer()
	defer srv.Shutdown()

	svc, err := factory(rawAppCfg)
	if err != nil {
		return
	}
	defer svc.Stop()

	srv.RegisterService(svc)

	if err = srv.Dial(
		cfg.AccessPoint.Address,
		WithGrpcDialOption(
			grpc.WithAuthority(cfg.AccessPoint.Authority),
			grpc.WithInsecure(),
		),
	); err != nil {
		return
	}
	wg.Add(1)
	go func() { defer wg.Done(); srv.Serve(lis) }()

	srv.WaitForTerm()
	return
}
