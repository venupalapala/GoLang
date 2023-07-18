package main

import (
	"fmt"
	"go-app/team1/controller"
	"go-app/team1/repository"
	"go-app/team1/service"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	transactionsRepository repository.TransactionsRepository = repository.New()
	transactionsService    service.TransactionsService       = service.New(transactionsRepository)
	transactionsController controller.TransactionsController = controller.New(transactionsService)
)

func main() {
	fmt.Println("Test")

	r := gin.Default()

	r.GET("/transactions/:tx_id", func(ctx *gin.Context) {

		fmt.Printf("Host : %s", ctx.Request.Host)
		transactionId := ctx.Param("tx_id")
		ctx.JSON(200, transactionsController.FindTransactionByTransactionId(transactionId))
	})

	// http://{hostname}:{port}/transactions/account/{acc_id}
	// http://{hostname}:{port}/transactions/account/{acc_id}?fromDateTime="yyyy-mm-dd hh:mm:ss"
	// http://{hostname}:{port}/transactions/account/{acc_id}?fromDateTime="yyyy-mm-dd hh:mm:ss"&toDateTime="yyyy-mm-dd hh:mm:ss"
	r.GET("/transactions/account/:acc_id", func(ctx *gin.Context) {
		accountId := ctx.Param("acc_id")
		fromDateTime := ctx.Query("fromDateTime")
		toDateTime := ctx.Query("toDateTime")
		status := ctx.Query("status")
		if len(fromDateTime) > 0 && len(toDateTime) > 0 {
			theFromTime, err := time.Parse(time.DateTime, fromDateTime)
			if err != nil {
				panic(err.Error())
			}
			theToTime, err := time.Parse(time.DateTime, toDateTime)
			if err != nil {
				panic(err.Error())
			}

			ctx.JSON(200, transactionsController.FindTransacationsByAccountIdBetweenDateTimeAndStatus(accountId, theFromTime, theToTime, status))

		} else if len(fromDateTime) > 0 {
			theTime, err := time.Parse(time.DateTime, fromDateTime)
			if err != nil {
				panic(err.Error())
			}

			ctx.JSON(200, transactionsController.FindTransactionsByAccountIdFromDateTimeAndStatus(accountId, theTime, status))

		} else {
			ctx.JSON(200, transactionsController.FindTransactionsByAccountId(accountId, status))
		}

	})

	r.Run(":9090")

}
