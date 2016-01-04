package config

import (
	"fmt"

	"testing"
)

const (
	testManagerURL = "http://127.0.0.1:8080"
	testListenAddr = ":9000"
)

var (
	sampleConfig = fmt.Sprintf(`
managerURL = "%s"
listenAddr = "%s"
tlsCACert = "ca.pem"
tlsCert = "cert.pem"
tlsKey = "key.pem"
`, testManagerURL, testListenAddr)
)

func TestParseConfig(t *testing.T) {
	cfg, err := ParseConfig(sampleConfig)
	if err != nil {
		t.Fatalf("error parsing config: %s", err)
	}

	if cfg.ManagerURL != testManagerURL {
		t.Fatalf("expected manager url %s; received %s", testManagerURL, cfg.ManagerURL)
	}

	if cfg.ListenAddr != testListenAddr {
		t.Fatalf("expected listen addr %s; received %s", testListenAddr, cfg.ListenAddr)
	}

	if cfg.TLSCACert != "ca.pem" {
		t.Fatalf("expected ca cert ca.pem; received %s", cfg.TLSCACert)
	}

	if cfg.TLSCert != "cert.pem" {
		t.Fatalf("expected cert ca.pem; received %s", cfg.TLSCert)
	}

	if cfg.TLSKey != "key.pem" {
		t.Fatalf("expected key key.pem; received %s", cfg.TLSKey)
	}
}
