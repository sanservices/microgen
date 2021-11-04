package tlsutil

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// TLS is the struct representation of the tls configuration
type TLS struct {
	CACertPEM   string `yaml:"ca_cert_pem"`
	CertPEM     string `yaml:"cert_pem"`
	KeyPEM      string `yaml:"key_pem"`
	SkipVerify  bool   `yaml:"skip_verify"`
	TimeoutSecs int    `yaml:"timeout_secs"`
}

func GetTLSConf(conf TLS) (*tls.Config, error) {

	caCertPEM, e := ioutil.ReadFile(conf.CACertPEM)
	if e != nil {
		log.Println("Could not open caCertPem", e)
		return nil, e
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCertPEM)

	cerPEM, e := tls.LoadX509KeyPair(conf.CertPEM, conf.KeyPEM)
	if e != nil {
		log.Println("Could not open certPem", e)
		return nil, e
	}

	t := tls.Config{
		Certificates:       []tls.Certificate{cerPEM},
		InsecureSkipVerify: true,
		RootCAs:            certPool}

	t.BuildNameToCertificate()

	return &t, e
}

func GetHTTPSClient(tlsConf *tls.Config) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConf,
			TLSNextProto:    make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
			Dial: func(network string, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, 10*time.Second)
			},
		},
	}
}
