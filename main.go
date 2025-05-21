// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
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
	// LOAD ENVS LATER SKIP FOR NOW
	db, err := database.ConnectDataBase()
	if err != nil {
		return err
	}
	return router.StartRouter(db)
}
