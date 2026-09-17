// Package mongo wraps the MongoDB driver connection used by all repositories.
package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Client owns the driver client and the selected database.
type Client struct {
	client *mongo.Client
	DB     *mongo.Database
}

// Connect opens a connection and verifies it with a ping.
func Connect(ctx context.Context, uri, dbName string) (*Client, error) {
	opts := options.Client().
		ApplyURI(uri).
		SetServerSelectionTimeout(5 * time.Second).
		SetConnectTimeout(5 * time.Second)

	c, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}
	if err := c.Ping(ctx, readpref.Primary()); err != nil {
		_ = c.Disconnect(ctx)
		return nil, fmt.Errorf("mongo ping: %w", err)
	}
	return &Client{client: c, DB: c.Database(dbName)}, nil
}

// Ping reports whether the primary is reachable.
func (c *Client) Ping(ctx context.Context) error {
	return c.client.Ping(ctx, readpref.Primary())
}

// Close disconnects the driver client.
func (c *Client) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}
