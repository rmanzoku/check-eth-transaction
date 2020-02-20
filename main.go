package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mackerelio/checkers"
	"github.com/rmanzoku/ethbq"
	"golang.org/x/net/context"
)

var (
	projectID = ""
	addresses = "0x0000000000000000000000000000000000000000"
	since     = time.Now().UTC().Add(-10 * time.Minute).Format("2006-01-02 15:04:05")
	until     = time.Now().UTC().Format("2006-01-02 15:04:05")
)

func run() (err error) {
	ctx := context.TODO()
	c, err := ethbq.NewClient(ctx, projectID)
	if err != nil {
		return
	}

	al := strings.Split(addresses, ",")
	for i, a := range al {
		al[i] = strconv.Quote(a)
	}
	in := strings.Join(al, ",")

	f := "SELECT `hash`," + `from_address, receipt_status FROM %s WHERE block_timestamp > "%s" AND block_timestamp < "%s" AND from_address IN (%s)`
	query := fmt.Sprintf(f, ethbq.TransactionsTable, since, until, in)
	it, err := c.Query(query)
	if err != nil {
		return
	}

	tx := []*ethbq.Transaction{}
	err = ethbq.UnmarshalTransactions(it, &tx)
	if err != nil {
		return
	}

	for _, t := range tx {
		if !t.Success() {
			return fmt.Errorf("%s has failure tx: %s", t.FromAddress, t.Hash)
		}
	}

	return nil
}

func main() {
	flag.StringVar(&projectID, "p", projectID, "projectID")
	flag.StringVar(&addresses, "a", addresses, "addresses")
	flag.Parse()

	chr := checkers.Ok("command ok!")
	err := run()
	if err != nil {
		chr = checkers.Critical(err.Error())
	}
	chr.Exit()
}
