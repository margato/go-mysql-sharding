package main

func main() {
	shards := CreateDatabaseShardsConnections()
	repository := NewCustomerGormRepository(shards)
	router := NewRouter(repository)

	_ = router.Run()
}
