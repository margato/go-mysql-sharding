package main

import "hash/fnv"

func HashShard(id string, numShards int) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(id))
	shardId := int(h.Sum32()) % numShards
	return shardId
}
