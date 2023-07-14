# Mongo
this is a simple project about how to use mongoDB with go

## How to use

### How to crate a Sql or Update

```go
    query := Db.Where("_id", Db.ConvertStringToObjectId("60595428721f2314cbbf6c65"))
    update := Db.Update().Set("app", "nezha_core")
```
when you want to find or update some data from mongoDB,you can use the code to create sql or update values.

### 1,Find

```go
    query := Db.Where("_id", Db.ConvertStringToObjectId("60595428721f2314cbbf6c65"))
	list := make([]UserSetting, 0)
	mongoTemplate.FindMany(ctx, query.GetCriteria(), &list)

	for _, v := range list {
		fmt.Println(v)
	}
	
	var userSetting UserSetting
	mongoTemplate.FindOne(ctx, query.GetCriteria(), &userSetting)
	fmt.Println(userSetting)
```
when you want to find some data from mongoDB,you can use the code above.
the all Find method will set the values to your struct.

### 2.Update

```go
    query := Db.Where("_id", Db.ConvertStringToObjectId("60595428721f2314cbbf6c65"))
    update := Db.Update().Set("app", "nezha_core")
    mongoTemplate.UpdateOne(ctx, query.GetCriteria(), update.GetUpdate())
```
when you want to update some data from mongoDB,you can use the code above.

