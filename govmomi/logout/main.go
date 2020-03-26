package main

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/pelletier/go-toml"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vapi/rest"
)

type vcConfig struct {
	VCenter struct {
		Server   string
		User     string
		Password string
		Insecure bool
	}
}

type vsClient struct {
	govmomi *govmomi.Client
	rest    *rest.Client
}

func main() {
	ctx := context.Background()

	cfg, err := loadTomlCfg("./.vcconfig")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("no error from loadToml")
	}

	clt, err := login(ctx, cfg)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("no error from login")
	}

	sleepTime := 2 * time.Minute
	fmt.Println("sleep for", sleepTime)
	time.Sleep(sleepTime)

	if err := logout(ctx, clt); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("no error from logout")
	}

}

func loadTomlCfg(path string) (*vcConfig, error) {
	var cfg vcConfig

	secret, err := toml.LoadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load vcconfig.toml: %w", err)
	}

	err = secret.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal vcconfig.toml: %w", err)
	}

	fmt.Println(cfg)
	return &cfg, nil
}

func login(ctx context.Context, cfg *vcConfig) (*vsClient, error) {
	var clt vsClient

	u := url.URL{
		Scheme: "https",
		Host:   cfg.VCenter.Server,
		Path:   "sdk",
	}
	u.User = url.UserPassword(cfg.VCenter.User, cfg.VCenter.Password)
	insecure := cfg.VCenter.Insecure

	gClt, err := govmomi.NewClient(ctx, &u, insecure)
	if err != nil {
		return nil, fmt.Errorf("connecting to govmomi api failed: %w", err)
	}
	clt.govmomi = gClt

	clt.rest = rest.NewClient(clt.govmomi.Client)
	err = clt.rest.Login(ctx, u.User)
	if err != nil {
		return nil, fmt.Errorf("log in to rest api failed: %w", err)
	}

	return &clt, nil
}

func logout(ctx context.Context, clt *vsClient) error {
	err := clt.govmomi.Logout(ctx)
	if err != nil {
		return fmt.Errorf("govmomi api logout failed: %w", err)
	}
	fmt.Println("govmomi logged out")

	err = clt.rest.Logout(ctx)
	if err != nil {
		return fmt.Errorf("rest api logout failed: %w", err)
	}
	fmt.Println("rest logged out")

	return nil
}
