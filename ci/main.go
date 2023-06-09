package main

import (
	"context"
	"os"

	"dagger.io/dagger"
	"github.com/wingyplus/must"
)

func main() {
	ctx := context.Background()

	client := must.Must(
		dagger.Connect(
			ctx,
			dagger.WithLogOutput(os.Stderr),
			dagger.WithWorkdir(".."),
		),
	)
	defer client.Close()

	must.Must(
		client.
			Pipeline("go").Pipeline("test").
			Container().
			From("golang:1.20").
			WithMountedDirectory("/go_headver", client.Host().Directory(".")).
			WithWorkdir("/go_headver").
			WithExec([]string{"go", "test", "./..."}).
			Sync(ctx),
	)
}
