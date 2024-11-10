package domingo

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo/config"
	"log"
	"net"
	"time"
)

// Credentials holds user credentials for authentication
type Credentials struct {
	UserID   string
	Password string
}

type Client struct {
	conn        net.Conn
	credentials Credentials
	hostname    string
	port        string
}

// NewClient initializes a new Client with TLS
func NewClient(cfg *config.EPPConfig) (*Client, error) {
	clientCert, err := tls.LoadX509KeyPair(cfg.ClientCertFile, cfg.ClientKeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate and private key: %w", err)
	}

	certPool := x509.NewCertPool()

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		RootCAs:            certPool,
		ServerName:         cfg.Hostname,
		MinVersion:         tls.VersionTLS13,
		CipherSuites:       []uint16{tls.TLS_AES_256_GCM_SHA384},
		InsecureSkipVerify: true,
	}

	address := fmt.Sprintf("%s:%s", cfg.Hostname, cfg.Port)
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 10 * time.Second}, "tcp", address, tlsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	credentials := Credentials{
		UserID:   cfg.UserID,
		Password: cfg.Password,
	}

	client := &Client{
		conn:        conn,
		credentials: credentials,
		hostname:    cfg.Hostname,
		port:        cfg.Port,
	}

	_, err = client.Read()
	if err != nil {
		log.Fatalf("Can not read greeting initial response: %v", err)
	}
	return client, nil
}

func (c *Client) Credentials() Credentials {
	return c.credentials
}

// Close terminates the TLS connection
func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Send(data []byte) error {
	_, err := c.conn.Write(data)
	return err
}

func (c *Client) Read() (string, error) {
	response := make([]byte, 4096)
	n, err := c.conn.Read(response)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}
	response = response[:n]
	return string(response), nil
}
