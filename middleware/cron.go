package middleware

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New(cron.WithSeconds(), cron.WithLocation(time.UTC))
	log.Println("Cron Time Zone:", c.Location())

	// UTC　(3分おき)
	// interval := "1 * * * * ?"

	// test for send mail
	// go c.AddFunc(interval, controllers.MailBySendGrid)

	// go c.AddFunc(interval, func() { fmt.Println("testです") })

	go c.Start()
	defer c.Stop()
}
