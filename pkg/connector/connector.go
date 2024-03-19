package connector

import (
	"context"
	"crypto/tls"
	"io"

	"github.com/conductorone/baton-panorama/pkg/panorama"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
)

type Panorama struct {
	Client *panorama.Client
}

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced from the upstream service.
func (d *Panorama) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{
		newUserBuilder(d.Client),
		newGroupBuilder(d.Client),
	}
}

// Asset takes an input AssetRef and attempts to fetch it using the connector's authenticated http client
// It streams a response, always starting with a metadata object, following by chunked payloads for the asset.
func (d *Panorama) Asset(ctx context.Context, asset *v2.AssetRef) (string, io.ReadCloser, error) {
	return "", nil, nil
}

// Metadata returns metadata about the connector.
func (d *Panorama) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "My Baton Connector",
		Description: "The template implementation of a baton connector",
	}, nil
}

// Validate is called to ensure that the connector is properly configured. It should exercise any API credentials
// to be sure that they are valid.
func (d *Panorama) Validate(ctx context.Context) (annotations.Annotations, error) {
	_, _, err := d.Client.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// New returns a new instance of the connector.
func New(ctx context.Context, panoramaUrl, username, password string, ignoreBadCertificate bool) (*Panorama, error) {
	clientOptions := []uhttp.Option{}
	if ignoreBadCertificate {
		l := ctxzap.Extract(ctx)
		l.Warn("Ignoring bad certificate. This should only be used for testing purposes.")

		clientOptions = append(clientOptions, uhttp.WithTLSClientConfig(
			&tls.Config{
				InsecureSkipVerify: true, // #nosec G402
			},
		))
	}

	httpClient, err := uhttp.NewBasicAuth(username, password).GetClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	client, err := panorama.New(panoramaUrl, httpClient)
	if err != nil {
		return nil, err
	}

	return &Panorama{
		Client: client,
	}, nil
}
