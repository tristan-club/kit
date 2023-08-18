package mailblacklist

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/tristan-club/kit/redisutil"
	"regexp"
	"strings"
)

const (
	RKEmailBlacklist = "blacklist:email"
)

var emailRE = regexp.MustCompile(`^(?:[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*)@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])$`)
var mailDomainRE = regexp.MustCompile(`@(.+)$`)

type Client struct {
	rdb *redis.Client
}

var cli *Client

func InitEmailBlackListChecker(redisSvc string, redisDB string) error {
	rdb, err := redisutil.NewClient(redisSvc, redisDB)
	if err != nil {
		return err
	}
	cli = &Client{rdb: rdb}
	return nil
}

func CheckEmailInBlackList(email string) (flag bool, err error) {

	email = strings.ToLower(email)

	if !emailRE.MatchString(email) {
		return false, fmt.Errorf("invalid email format")
	}

	domain := mailDomainRE.FindStringSubmatch(email)
	if len(domain) != 2 {
		return false, fmt.Errorf("invalid domain format")
	}

	if cli == nil {
		return false, fmt.Errorf("client not ready")
	}

	flag, err = cli.rdb.SIsMember(context.Background(), RKEmailBlacklist, domain[1]).Result()
	if err != nil && err != redis.Nil {
		return false, fmt.Errorf("get blacklist error: %s", err.Error())
	}

	return flag, nil
}

func AddEmailBlacklist(emails []string) error {
	if cli == nil {
		return fmt.Errorf("client not ready")
	}

	pipeliner := cli.rdb.Pipeline()

	for _, email := range emails {
		pipeliner.SAdd(context.Background(), RKEmailBlacklist, email)
	}

	_, err := pipeliner.Exec(context.Background())
	if err != nil {
		return err
	}
	return nil

}
