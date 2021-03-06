//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package docker

import (
	"fmt"
	"github.com/docker/docker/api"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/tlsconfig"
	"github.com/lastbackend/lastbackend/pkg/common/config"
	"net/http"
	"path/filepath"
)

type Runtime struct {
	client *client.Client
}

func New(cfg config.Docker) (*Runtime, error) {

	var (
		err error
		cli *http.Client
		r   = new(Runtime)
	)

	fmt.Println("Use docker CRI")

	if cfg.Certs != nil && *cfg.Certs != "" {

		fmt.Printf("Create Docker secure client: %s", cfg.Certs)

		options := tlsconfig.Options{
			CAFile:             filepath.Join(*cfg.Certs, "ca.pem"),
			CertFile:           filepath.Join(*cfg.Certs, "cert.pem"),
			KeyFile:            filepath.Join(*cfg.Certs, "key.pem"),
			InsecureSkipVerify: *cfg.TLS,
		}

		tlsc, err := tlsconfig.Client(options)
		if err != nil {
			return nil, err
		}

		cli = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsc,
			},
		}
	}

	host := client.DefaultDockerHost
	if cfg.Host != nil && *cfg.Host != "" {
		host = *cfg.Host
	}

	version := api.DefaultVersion
	if cfg.Host != nil && *cfg.Host != "" {
		version = *cfg.Version
	}

	r.client, err = client.NewClient(host, version, cli, nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}
