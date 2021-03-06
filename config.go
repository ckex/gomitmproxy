package main

import (
	"crypto/tls"
	"time"
)

const (
	ONE_DAY   = 24 * time.Hour
	TWO_WEEKS = ONE_DAY * 14
	ONE_MONTH = 1
	ONE_YEAR  = 1
)

type Cfg struct {
	Port    *string
	Raddr   *string
	Log     *string
	Monitor *bool
	Tls     *bool
}

type TlsConfig struct {
	PrivateKeyFile  string
	CertFile        string
	Organization    string
	CommonName      string
	ServerTLSConfig *tls.Config
}

func NewTlsConfig(pk, cert, org, cn string) *TlsConfig {
	return &TlsConfig{
		PrivateKeyFile: pk,
		CertFile:       cert,
		Organization:   org,
		CommonName:     cn,
		ServerTLSConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
				tls.TLS_RSA_WITH_RC4_128_SHA,
				tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			},
			PreferServerCipherSuites: true,
		},
	}
}
