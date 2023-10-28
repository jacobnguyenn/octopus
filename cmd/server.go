package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/spf13/cobra"

	pkgerrors "github.com/pkg/errors"

	"ddd-sample/registry"
)

func newServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "server",
		DisableAutoGenTag: true,
		SilenceErrors:     true,
		SilenceUsage:      true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			log.Println("starting server")
			var wg sync.WaitGroup
			errCh := make(chan error, 1)

			srv := registry.InitializeServer()
			l, err := net.Listen(
				"tcp",
				fmt.Sprintf(
					"%s:%s",
					"0.0.0.0",
					"8080",
				),
			)
			if err != nil {
				return pkgerrors.Wrap(err, "error creating listener")
			}
			defer func() {
				_ = l.Close()
			}()
			wg.Add(1)
			go func() {
				srvErr := srv.Serve(ctx, l)
				errCh <- pkgerrors.Wrap(srvErr, "serve")
				wg.Done()
			}()
			wg.Wait()
			close(errCh)
			var resErr error
			for err := range errCh {
				resErr = errors.Join(resErr, err)
			}
			return resErr
		},
	}
}
