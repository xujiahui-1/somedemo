package main

import (
	"fmt"
	"testing"
	"time"
)

/*
建造者模式
	当要构建的对象很大并且需要多个步骤时，使用构建器模式，有助于减小构造函数的大小。
	就是连续调用的感觉，类似。Coffee.builder().name("Latti").price("30").build()

*/

// 例；假设我们要链接数据库
type DBPool struct {
	dsn             string
	maxOpenConn     int
	maxIdleConn     int
	maxConnLifeTime time.Duration
}

// DBPoolBuilder 数据库连接池建造者
type DBPoolBuilder struct {
	DBPool DBPool
	err    error
}

// Builder 建造者,给予各种属性默认值
func Builder() *DBPoolBuilder {
	b := new(DBPoolBuilder)
	b.DBPool.dsn = "127.0.0.1"
	b.DBPool.maxConnLifeTime = 1 * time.Second
	b.DBPool.maxOpenConn = 30
	return b

}

// 各种属性的配制方法，返回DBPoolBuilder
func (b *DBPoolBuilder) DSN(dsn string) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if dsn == "" {
		b.err = fmt.Errorf("invalid dsn,current is %s", dsn)
	}
	b.DBPool.dsn = dsn
	return b
}
func (b *DBPoolBuilder) MaxOpenConn(connNum int) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if connNum < 1 {
		b.err = fmt.Errorf("invalid MaxOpenConn, current is %d", connNum)
	}

	b.DBPool.maxOpenConn = connNum
	return b
}
func (b *DBPoolBuilder) MaxConnLifeTime(lifeTime time.Duration) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if lifeTime < 1*time.Second {
		b.err = fmt.Errorf("connection max life time can not litte than 1 second, current is %v", lifeTime)
	}

	b.DBPool.maxConnLifeTime = lifeTime
	return b
}
func (b *DBPoolBuilder) Build() (*DBPool, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.DBPool.maxOpenConn < b.DBPool.maxIdleConn {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", b.DBPool.maxOpenConn, b.DBPool.maxIdleConn)
	}
	return &b.DBPool, nil
}

func Test_build(t *testing.T) {
	dbpool, err := Builder().DSN("localhost:3306").MaxConnLifeTime(50).MaxOpenConn(30).MaxConnLifeTime(0 * time.Second).Build()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbpool)
}
