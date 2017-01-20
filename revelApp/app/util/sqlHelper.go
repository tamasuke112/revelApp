package util

import (
        "io/ioutil"
        "fmt"
)

func GetSql(sqlId string) string{

	var tempSqlId	string = sqlId +".sql"
	var basePath    string = "revelApp/app/sql/"

    sql, err := ioutil.ReadFile(basePath+tempSqlId)

        if err != nil {
        panic(err)
    }

    fmt.Println("== parsing Query start     ==")
    fmt.Println(string(sql))
	fmt.Println("== parsing Query end               ==")

    return string(sql)
}
