package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

func initDB() (err error) {
	dsn := "root:admin1234@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed:err:%v\n", err)
		return
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(150)
	return
}

func queryRowDemo() {
	sqlStr := "select id,name,age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed,err=%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

func queryMultiRowDemo() {
	sqlStr := "select id,name,age from user where id>?"
	var users []user
	err := db.Select(&users, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed,err=%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

func insertRowDemo() {
	sqlStr := "insert into user(name,age) values (?,?)"
	ret, err := db.Exec(sqlStr, "张三", 20)
	if err != nil {
		fmt.Printf("insert failed, err=%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n",theID)
}
	
func updateRowDemo() {
	sqlStr := "update user set age= ? where id = ?"
	ret, err := db.Exec(sqlStr, 11,1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}

func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("deleted failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
}

func transactionDemo() error {
    // 开始一个新的事务
    tx, err := db.Beginx()
    if err != nil {
        return err // 如果开始事务失败，返回错误
    }

    // 使用 defer 确保在结束时正确处理事务
    defer func() {
        if p := recover(); p != nil {
            // 如果发生 panic，则回滚事务
            tx.Rollback()
            panic(p) // 重新抛出 panic，以便上层调用处理
        } else if err != nil {
            // 如果发生错误，回滚事务
            fmt.Println("Rollback due to error:", err)
            tx.Rollback()
        } else {
            // 如果没有错误，提交事务
            fmt.Println("Committing transaction")
            err = tx.Commit()
        }
    }()

    _, err = tx.Exec("INSERT INTO user(name,age) VALUES(?, ?)", "lisi1",10)
    if err != nil {
        return err // 记录插入失败，返回错误
    }

    _, err = tx.Exec("INSERT INTO user(id,name,age) VALUES(?, ?, ?)", 6,"lisi2",10)
    if err != nil {
        return err // 记录插入失败，返回错误
    }

    return nil // 如果没有错误，返回 nil
}