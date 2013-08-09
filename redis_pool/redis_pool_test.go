package redis_connection_pool

import "os"
import "testing"
import "github.com/alecthomas/log4go"

var redis_master = os.Getenv("REDIS_MASTER")

const method_expected_a_equal_b_from_url = "[%s][%s] Expected '%v', Actual '%v'"

func Test_initRedisConnection_Redis_master(t *testing.T) {
	tag := "initRedisConnection - Redis_master"

	client, err := initRedisConnection(redis_master)
	if nil != err {
		t.Errorf(method_expected_a_equal_b_from_url, tag, "err", err, nil)
		return
	}

	if nil == client {
		t.Errorf(method_expected_a_equal_b_from_url, tag, "client", client, "!nil")
		return
	}
}

func Test_initRedisConnection_Invalid_host(t *testing.T) {
	tag := "initRedisConnection - Invalid_host"

	// Suppress error/warning messages
	SetLogLevel(log4go.CRITICAL)
	defer ResetLogLevel()

	client, err := initRedisConnection("localhost:9999")
	if nil == err {
		t.Errorf(method_expected_a_equal_b_from_url, tag, "err", err, "!nil")
		return
	}

	if nil != client {
		t.Errorf(method_expected_a_equal_b_from_url, tag, "client", client, nil)
		return
	}
}