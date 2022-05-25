// Copyright 2019 NetApp, Inc. All Rights Reserved.

package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"

	"github.com/netapp/trident/cli/api"
	"github.com/netapp/trident/utils"
)

var (
	allSnapshots         bool
	allSnapshotsInVolume string
)

func init() {
	deleteCmd.AddCommand(deleteSnapshotCmd)
	deleteSnapshotCmd.Flags().BoolVar(&allSnapshots, "all", false, "Delete all snapshots")
	deleteSnapshotCmd.Flags().StringVar(&allSnapshotsInVolume, "volume", "", "Delete all snapshots in volume")
}

var deleteSnapshotCmd = &cobra.Command{
	Use:     "snapshot <volume/snapshot> [<volume/snapshot>...]",
	Short:   "Delete one or more volume snapshots from Trident",
	Aliases: []string{"s", "snap", "snapshots"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if OperatingMode == ModeTunnel {
			command := []string{"delete", "snapshot"}
			if allSnapshotsInVolume != "" {
				command = append(command, "--volume", allSnapshotsInVolume)
			}
			if allSnapshots {
				command = append(command, "--all")
			}
			TunnelCommand(append(command, args...))
			return nil
		} else {
			return snapshotDelete(args)
		}
	},
}

func snapshotDelete(snapshotIDs []string) error {
	var err error

	if allSnapshotsInVolume != "" {
		// Make sure --volume isn't being used along with specific snapshots
		if len(snapshotIDs) > 0 {
			return errors.New("cannot use --volume switch and specify individual snapshots")
		}

		// Get list of snapshot IDs in the specified volume so we can delete them all
		snapshotIDs, err = GetSnapshots(allSnapshotsInVolume)
		if err != nil {
			return err
		}

	} else if allSnapshots {
		// Make sure --all isn't being used along with specific snapshots
		if len(snapshotIDs) > 0 {
			return errors.New("cannot use --all switch and specify individual snapshots")
		}

		// Get list of snapshot IDs so we can delete them all
		snapshotIDs, err = GetSnapshots("")
		if err != nil {
			return err
		}
	} else {
		// Not using --all or --volume, each snapshot must have volume specified
		if len(snapshotIDs) == 0 {
			return errors.New("volume/snapshot not specified")
		}
		for _, snapshotID := range snapshotIDs {
			if !strings.ContainsRune(snapshotID, '/') {
				return utils.InvalidInputError(fmt.Sprintf("invalid snapshot ID: %s; Please use the format "+
					"<volume name>/<snapshot name>", snapshotID))
			}
		}
	}

	for _, snapshotID := range snapshotIDs {
		url := BaseURL() + "/snapshot/" + snapshotID

		response, responseBody, err := api.InvokeRESTAPI("DELETE", url, nil, Debug)
		if err != nil {
			return err
		} else if response.StatusCode != http.StatusOK {
			return fmt.Errorf("could not delete snapshot %s: %v", snapshotID,
				GetErrorFromHTTPResponse(response, responseBody))
		}
	}

	return nil
}
