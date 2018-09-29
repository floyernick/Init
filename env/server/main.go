package server

import (
	"Init/env/config"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
)

func RunServer(config config.ServerConfig, mux *http.ServeMux) error {

	server := &http.Server{
		Addr:         config.Port,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      mux,
	}

	if !config.SecureConn {

		return server.ListenAndServe()

	} else {

		cert, err := tls.LoadX509KeyPair(config.SecureCert, config.SecureKey)

		if err != nil {
			return err
		}

		certPool := x509.NewCertPool()

		ca, err := ioutil.ReadFile(config.SecureCA)

		if err != nil {
			return err
		}

		if !certPool.AppendCertsFromPEM(ca) {
			return errors.New("invalid CA certificate")
		}

		TLSConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			ClientCAs:    certPool,
		}

		server.TLSConfig = TLSConfig

		return server.ListenAndServeTLS("", "")

	}

}
