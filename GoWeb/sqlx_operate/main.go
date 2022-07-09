package main

import (
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	ID   int    `db:"id"`
	Age  int    `db"age"`
	Name string `db"name"`
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func insertUserDemo() (err error) {
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "七米",
			"age":  28,
		})
	return
}

func namedQuery() {
	sqlStr := "SELECT * FROM user WHERE name=:name"
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七米"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

	u := user{
		Name: "七米",
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

func (u user) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err := db.Exec(query, args...)
	return err
}

// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*user) error {
	_, err := db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}

// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int) (users []user, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []user, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init DB faied, err: %v\n", err)
		return
	}
	fmt.Println("init DB success...")

	//defer db.Close()
	//u1 := user{Name: "xx", Age: 18}
	//u2 := user{Name: "xxx", Age: 28}
	//u3 := user{Name: "xxxx", Age: 38}
	//
	//// 方法1
	//users := []*user{&u1, &u2, &u3}
	//err = BatchInsertUsers(users)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	//}
	//
	//// 方法2
	//users2 := []interface{}{u1, u2, u3}
	//err = BatchInsertUsers2(users2)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	//}
	//
	//// 方法3
	//users3 := []*User{&u1, &u2, &u3}
	//err = BatchInsertUsers3(users3)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	//}

	//users, err := QueryByIDs([]int{1, 2, 3})
	//if err != nil {
	//	fmt.Printf("QueryByIDs failed, err:%v\n", err)
	//	return
	//}
	//for _, user := range users {
	//	fmt.Printf("user:%#v\n", user)
	//}

	//users, err := QueryAndOrderByIDs([]int{1, 2, 3})
	//if err != nil {
	//	fmt.Printf("QueryByIDs failed, err:%v\n", err)
	//	return
	//}
	//for _, user := range users {
	//	fmt.Printf("user:%#v\n", user)
	//}
}
