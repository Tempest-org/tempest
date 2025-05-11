package datasource

import (
	"context"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
)

func Migrate(ctx context.Context, url string, migrationsDir string) (string, error){
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS(migrationsDir),
		),
	)
	if err != nil {
		return "", err
	}
	defer workdir.Close()

	client ,err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		return "", err
	}

	res, err := client.MigrateApply(ctx, &atlasexec.MigrateApplyParams{
		URL:url,
	})
	if err != nil {
		return "", err
	}


	return res.Summary("    "), nil
}
