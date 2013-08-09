//
// Redis Connection Pool written in GO
//

package redis_pool

import "time"
import "github.com/fzzy/radix/redis"
import "github.com/alecthomas/log4go"
import cp "github.com/gnagel/go-connection_pool/connection_pool"

//
// Constants for connecting to Redis & Logging
//
const timeout = time.Duration(10) * time.Second
const log_client_premade_success = "[RedisConnectionFactory][Client] - Pre-made availble for %s"
const log_client_premade_failure = "[RedisConnectionFactory][Client] - No saved connection available for %s"
const log_open_failed = "[RedisConnectionFactory][Open] - Failed to connect to %s, error = %#v"
const log_open_success = "[RedisConnectionFactory][Open] - Opened new connection to %s"
const log_closed = "[RedisConnectionFactory][Close] - Closed connection to %s, error = %#v"

//
// Connection Wrapper for Redis
//
type RedisConnectionFactory struct {
	Url string "Redis URL this factory will connect to"

	Logger *log4go.Logger "Handle to the logger we are using"

	client *redis.Client "Connection to a Redis, may be nil"
}

func newLazyFactory(url string, Logger *log4go.Logger) *RedisConnectionFactory {
	// Create a new factory instance
	factory := &RedisConnectionFactory{Url: nextUrl(), p.Logger}

	// Return the factory
	return factory, nil
}

func newAgressiveFactory(url string, Logger *log4go.Logger) (*RedisConnectionFactory, error) {
	// Create a new factory instance
	factory := newLazyFactory(url, logger)

	// Open the connection to Redis
	client, err := factory.Client()
	if nil != err {
		return nil, err
	}

	// Ping the server
	client.Append("ping")

	// Get the response
	reply := client.GetReply()

	// Connection error? Then tell the factory to invalidate the Redis connection
	if nil != reply.Err {
		// Close the connection
		factory.Close(reply.Err)

		return nil, reply.Err
	}

	// Return the factory
	return factory, nil
}

//
// Get a connection to Redis
//
func (p *RedisConnectionFactory) Client() (*redis.Client, error) {
	// If the connection is valid, return it
	if nil != p.client {
		// Log the event
		p.Logger.Trace(log_client_premade_success, p.Url)

		// Return the connection
		return p.client, nil
	}

	// Log the event
	p.Logger.Warn(log_client_premade_failure, p.Url)

	// Open a new connection to redis
	if err := Open(); nil != err {
		// Errors are already logged in Open()
		return nil, err
	}

	// Return the new redis connection
	return p.client, nil
}

//
// Open a new connection to redis
//
func (p *RedisConnectionFactory) Open() error {
	// Connect to Redis
	client, err := redis.DialTimeout("tcp", p.Url, timeout)

	if nil != err {
		// Log the event
		log.Critical(log_open_failed, p.Url, err)

		// Return the error
		return err
	}

	// Log the event
	p.Logger.Info(log_open_success, p.Url)

	return client, err
}

//
// Close the connection to redis
//
func (p *RedisConnectionFactory) Close(err error) {
	// Log the event
	log.Warn(log_closed, p.Url, err)

	// Close the connection
	if nil != p.client {
		p.client.Close()
	}

	// Set the pointer to nil
	p.client = nil
}
