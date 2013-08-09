go-redis_pool
=============

Redis Connection Pool written in GO


Install
=======

	go get -u "github.com/gnagel/go-redis_pool/redis_pool"


GO Usage
========

	import "github.com/gnagel/go-redis_pool/redis_pool"
	import redis "github.com/fzzy/radix/redis"
	
	// Setup the pool
	urls := []string{"127.0.0.1:6379"}
	mode := redis_pool.LAZY // SIMPLE, AGRESSIVE, etc
	pool := redis_pool.NewPool(mode, urls)
	
	// Pop a Redis connection from the pool
	factory := pool.Pop();
	defer pool.Push(factory)
	
	// Get the redis client from the factory
	// This will re-connect if the client was previously disconnected
	client, err := factory.Client()
	
	// No connection available!
	if  nil != err {
		panic(err)
	}
	
	// Ping the server
	client.Append("ping")
	// Get the response
	reply := client.GetReply()
	
	// Connection error? Then tell the factory to invalidate the Redis connection
	if nil != reply.Err {
		// Close the connection
		factory.Close(reply.Err)
		
		// Exit your test/go routine/app/etc here
		return reply.Err
	}
	
	// ...
	// ...
	// ...
	
	


Authors:
========

Glenn Nagel <glenn@mercury-wireless.com>, <gnagel@rundsp.com>


Credits:
========

Juhani Åhman's redis implementation [rocks](https://github.com/fzzy/radix)
