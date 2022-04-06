# Package `store`
Package `store` declares interface `DataGateway`, which is a high-level abstraction of data storage.

![DataStoreAbstraction](https://github.com/yimsoijoi/todong/blob/main/assets/todogin_store.png?raw=true)

# Interface `DataGateway`
Both `gormDataGateway` and `redisDataGateway` structs implement `DataGateway`.

# Interface `GormStore` (implemented by struct `gormStore`)
Interface `GormStore` abstracts `GormDB`, and is more convinient to use than `GormDB`.

# Interface `GormDB`
Interface `GormDB` represents the methods used by `*gorm.DB`