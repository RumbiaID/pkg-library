This library will be used as main library for helpers


````
go get github.com/RumbiaID/pkg-library
````

to update:
````
go get -u github.com/RumbiaID/pkg-library@v.x.x
````

How to Use CQRS

1. In the main package you are using, make sure you import "github.com/RumbiaID/pkg-library/app/pkg/database".

2. During the declaration of using the database, you would run a function from the database folder called NewDatabase(driver string, cfg *Config, isCqrs bool)

   2.1 Config of db consists of:
    ````
    type Config struct {
    DbHost   string
    DbUser   string
    DbPass   string
    DbName   string
    DbPort   string
    DbPrefix string
    }
    `````
   2.2 Driver you can use: "postgres" / "pgsql", "mysql", "sqlserver", "oracle".
   
   2.3 isCqrs you may fill with a Boolean value of false or true.

This declaration of the Database will return a struct consisting of:
```
type Database struct {
    db     *gorm.DB
    isCqrs bool
}
```

3. If you input the isCqrs as true, you may use CqrsDB(driver string, cfg *Config) to insert the replica value.

4. For migrating purposes, you would use methods from the Database struct:
````
    MigrateDB(dst ...interface{})
    DownMigrate(all bool, dst ...interface{})
    DropColumnDB(dst interface{}, columnTarget string)
    RenameColumnDB(dst interface{}, oldname, columnTarget string)
    DownIndexDB(dst interface{}, columnTarget string)
    WipeTable(dst interface{})
    DeleteTable(dst ...interface{})
````    

5. Depending on the isCqrs value in the Database struct, the migration would run in master only or master-replica.
