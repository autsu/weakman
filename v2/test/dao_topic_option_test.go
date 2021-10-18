package test

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
	"vote/v2/dao"
	"vote/v2/model"
	"vote/v2/pkg"

	"github.com/jmoiron/sqlx"
)

func TestTopicOptionQueryById(t *testing.T) {
	var to *dao.TopicOptionDao
	option, total, err := to.QueryByTopicId("1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("total: %v, options: %v \n", total, option)
}

func TestOptionShowParticipantById(t *testing.T) {
	var to *dao.TopicOptionDao
	name, err := to.ShowParticipantById("1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(name)
}

// 并发下票数是否正常
func TestSingleVote(t *testing.T) {
	var to *dao.TopicOptionDao
	var wg sync.WaitGroup
	var num = 100

	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			to.SingleVote(&model.VoteRecord{
				Uid:      i,
				OptionId: 39,
				Time:     time.Now(),
			}, 1, "34")
		}(i)
	}

	wg.Wait()
}

func TestSingleVote1(t *testing.T) {
	var to *dao.TopicOptionDao
	to.SingleVote(&model.VoteRecord{
		Uid:      1,
		OptionId: 39,
		Time:     time.Now(),
	}, 1, "34")
}

func TestSingleVote2(t *testing.T) {
	var to *dao.TopicOptionDao
	to.SingleVote(&model.VoteRecord{
		Uid:      1,
		OptionId: 39,
		Time:     time.Now(),
	}, 1, "34")
}

func update(db *sqlx.Tx, num int32) {
	sql := `
update topic_option set number = ? 
where id = ?
`
	_, err := db.Exec(sql, num, 39)
	if err != nil {
		log.Fatalln(err)
	}

}

func Test123(t *testing.T) {
	db, err := pkg.NewMysql()
	if err != nil {
		log.Fatalln(err)
	}

	sql := `
update topic_option set number = ? 
where id = ?
`

	go func() {
		tx, err := db.Beginx()
		fmt.Println("事务2 开始执行")
		if err != nil {
			log.Fatalln(err)
		}

		_, err = tx.Exec(sql, 111, 39)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("事务2 update 执行完成")

		if err == nil {
			// 事务2 先执行，然后卡着不提交
			time.Sleep(time.Second * 5)
			tx.Commit()
			fmt.Println("事务2 提交")
		}
	}()

	time.Sleep(time.Second)
	tx, err := db.Beginx()
	fmt.Println("事务1 开始执行")
	if err != nil {
		log.Fatalln(err)
	}

	// 事务1 后执行，因为 事务2 还没有 commit，所以这里
	// 应该会被阻塞，直到 事务1 commit
	_, err = tx.Exec(sql, 222, 39)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("事务1 update 执行完成")

	if err == nil {
		//time.Sleep(time.Second * 5)
		tx.Commit()
		fmt.Println("事务1 提交")
	}

	time.Sleep(time.Second * 5)
}

func Test123_(t *testing.T) {
	db, err := pkg.NewMysql()
	if err != nil {
		log.Fatalln(err)
	}

	sql := `
update topic_option set number = ? 
where id = ?
`

	go func() {
		db1, err := pkg.NewMysql()
		if err != nil {
			log.Fatalln(err)
		}
		tx, err := db1.Beginx()
		fmt.Println("事务2 开始执行")
		if err != nil {
			log.Fatalln(err)
		}

		_, err = tx.Exec(sql, 111, 39)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("事务2 update 执行完成")

		if err == nil {
			// 事务2 先执行，然后卡着不提交
			time.Sleep(time.Second * 5)
			tx.Commit()
			fmt.Println("事务2 提交")
		}
	}()

	time.Sleep(time.Second)
	tx, err := db.Beginx()
	fmt.Println("事务1 开始执行")
	if err != nil {
		log.Fatalln(err)
	}

	// 事务1 后执行，因为 事务2 还没有 commit，所以这里
	// 应该会被阻塞，直到 事务1 commit
	_, err = tx.Exec(sql, 222, 39)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("事务1 update 执行完成")

	if err == nil {
		//time.Sleep(time.Second * 5)
		tx.Commit()
		fmt.Println("事务1 提交")
	}

	time.Sleep(time.Second * 5)
}
