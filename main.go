// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"context"

	"github.com/aidenfine/tny/config"
	"github.com/aidenfine/tny/database"
	"github.com/aidenfine/tny/tny-src/router"
)

func main() {
	err := setup()
	if err != nil {
		panic(err)
	}

}

func setup() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}
	db, err := database.ConnectDatabase()
	if err != nil {
		return err
	}
	ctx := context.Background()

	return router.StartRouter(ctx, db)
}
