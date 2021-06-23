package prueba_test

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "testing"
)

type A struct {
    X string
    Y string
}

type B struct {
    X string
    Y string
}

type AB struct {
    B B `gorm:"embedded"` // Embedded struct B before struct A
    A A `gorm:"embedded"`
}

var DB *gorm.DB

func connectDB() error {
    var err error
    spec := "slumberuser:password@tcp(localhost:3306)/slumber"
    DB, err = gorm.Open("mysql", spec+"?parseTime=true&loc=UTC&charset=utf8")
    DB.LogMode(true) // Print SQL statements
    //defer DB.Close()
    if err != nil {
        return err
    }
    return nil
}

// cd tests
// go test -run TestGormEmbed
func TestGormEmbed(t *testing.T) {
    if err := connectDB(); err != nil {
        t.Errorf("error connecting to db %v", err)
    }
    values := []interface{}{&A{}, &B{}}
    for _, value := range values {
        DB.DropTable(value)
    }
    if err := DB.AutoMigrate(values...).Error; err != nil {
        panic(fmt.Sprintf("No error should happen when create table, but got %+v", err))
    }
    DB.Save(&A{X: "AX1", Y: "AY1"})
    DB.Save(&A{X: "AX2", Y: "AY2"})
    DB.Save(&B{X: "BX1", Y: "BY1"})
    DB.Save(&B{X: "BX2", Y: "BY2"})

    //select * from `as` join `bs`;
    //    # x,y,x,y
    //    # AX1,AY1,BX1,BY1
    //    # AX2,AY2,BX1,BY1
    //    # AX1,AY1,BX2,BY2
    //    # AX2,AY2,BX2,BY2

    var abs []AB
    DB.Select("*").
    Table("as").
    Joins("join bs").
        Find(&abs)
    for _, ab := range abs {
        fmt.Println(ab.A, ab.B)
    }
    // if it worked it should print

    //{AX1 AY1} {BX1 BY1}
    //{AX2 AY2} {BX1 BY1}
    //{AX1 AY1} {BX2 BY2}
    //{AX2 AY2} {BX2 BY2}

    // but actually prints
    //{BX1 BY1} {AX1 AY1}
    //{BX1 BY1} {AX2 AY2}
    //{BX2 BY2} {AX1 AY1}
    //{BX2 BY2} {AX2 AY2}
}