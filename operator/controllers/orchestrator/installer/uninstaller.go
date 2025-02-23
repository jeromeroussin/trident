// Copyright 2021 NetApp, Inc. All Rights Reserved.

package installer

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/netapp/trident/cli/cmd"
	k8sclient "github.com/netapp/trident/cli/k8s_client"
)

func (i *Installer) UninstallTrident() error {
	// 1. preview CSI Trident --> uninstall preview CSI Trident
	// 2. preview CSI Trident & legacy Trident --> uninstall preview CSI Trident
	// 3. CSI Trident --> uninstall CSI Trident
	// 4. legacy Trident --> uninstall legacy Trident
	//
	// if csiPreview, uninstall csiPreview
	// else if csi, uninstall csi
	// else if legacy, uninstall legacy

	// Check if legacy Trident is installed
	legacyTridentInstalled, _, err := i.client.CheckDeploymentExistsByLabel(TridentLegacyLabel, true)
	if err != nil {
		return fmt.Errorf("could not check if legacy Trident is installed; %v", err)
	}

	// Check if preview CSI Trident is installed
	csiPreviewTridentInstalled, _, err := i.client.CheckStatefulSetExistsByLabel(TridentCSILabel, true)
	if err != nil {
		return fmt.Errorf("could not check if preview CSI Trident is installed; %v", err)
	}

	if legacyTridentInstalled && csiPreviewTridentInstalled {
		log.Warning("Both legacy and CSI Trident are installed.  CSI Trident will be uninstalled, and " +
			"the uninstaller will run again to remove legacy Trident before running the Trident installer.")
	}

	// Set the global csi variable, which controls things like RBAC and app labels
	// Should not use csiPreviewTridentInstalled || csiTridentInstalled as it give false when CSI trident
	// installation is deleted
	csi = !legacyTridentInstalled

	// Set the app labels (CSI takes precedence)
	if csi {
		appLabel = TridentCSILabel
		appLabelKey = TridentCSILabelKey
		appLabelValue = TridentCSILabelValue
	} else {
		appLabel = TridentLegacyLabel
		appLabelKey = TridentLegacyLabelKey
		appLabelValue = TridentLegacyLabelValue
	}
	nodeLabel := TridentNodeLabel

	// First handle the deployment (legacy, CSI) / statefulset (preview CSI)
	if csiPreviewTridentInstalled {
		if err := i.client.DeleteTridentStatefulSet(appLabel); err != nil {
			return fmt.Errorf("could not delete Trident CSI preview deployment; %v", err)
		}
	} else {
		if err := i.client.DeleteTridentDeployment(appLabel); err != nil {
			return fmt.Errorf("could not delete Trident CSI deployment; %v", err)
		}
	}

	// Next handle all the other common CSI components (daemonset, service).  Some/all of these may
	// not be present if uninstalling legacy Trident or preview CSI Trident, in which case we log
	// warnings only.
	if err := i.client.DeleteTridentDaemonSet(nodeLabel); err != nil {
		return fmt.Errorf("could not delete Trident daemonset; %v", err)
	}

	if err := i.client.DeleteTridentService(getServiceName(), appLabel, i.namespace); err != nil {
		return fmt.Errorf("could not delete Trident service; %v", err)
	}

	if err := i.client.DeleteTridentSecret(getProtocolSecretName(), appLabel, i.namespace); err != nil {
		return fmt.Errorf("could not delete Trident secret; %v", err)
	}

	csiDriverName := getCSIDriverName()
	if i.client.ServerVersion().AtLeast(KubernetesVersionMinV1CSIDriver) {
		if err := i.client.DeleteCSIDriverCR(csiDriverName, appLabel); err != nil {
			return fmt.Errorf("could not delete Trident CSI driver custom resource; %v", err)
		}
	} else {
		if err := i.client.DeleteBetaCSIDriverCR(csiDriverName, appLabel); err != nil {
			return fmt.Errorf("could not delete Trident Beta CSI driver custom resource; %v", err)
		}
	}

	// Delete Trident RBAC objects
	if err := i.removeRBACObjects(csi); err != nil {
		return fmt.Errorf("could not delete all Trident's RBAC objects; %v", err)
	}

	if err := i.client.DeleteTridentPodSecurityPolicy(getPSPName(), appLabel); err != nil {
		return err
	}

	log.Info("The uninstaller did not delete Trident's namespace in case it is going to be reused.")

	return nil
}

func (i *Installer) UninstallCSIPreviewTrident() error {
	appLabel = TridentCSILabel
	appLabelKey = TridentCSILabelKey
	appLabelValue = TridentCSILabelValue

	return i.client.DeleteTridentStatefulSet(appLabel)
}

func (i *Installer) UninstallLegacyTrident() error {
	appLabel = TridentLegacyLabel
	appLabelKey = TridentLegacyLabelKey
	appLabelValue = TridentLegacyLabelValue

	if err := i.client.DeleteTridentDeployment(appLabel); err != nil {
		return err
	}

	// legacy Trident will not use Trident-CSI names
	return i.removeRBACObjects(false)
}

// removeRBACObjects removes any ClusterRoleBindings, ClusterRoles,
// ServicesAccounts and OpenShiftSCCs associated with legacy Trident or Trident-CSI.
func (i *Installer) removeRBACObjects(csi bool) error {
	// Delete cluster role binding
	if err := i.client.DeleteTridentClusterRoleBinding(getClusterRoleBindingName(csi), appLabel); err != nil {
		return err
	}

	// Delete cluster role
	if err := i.client.DeleteTridentClusterRole(getClusterRoleName(csi), appLabel); err != nil {
		return err
	}

	// Delete service account
	if err := i.client.DeleteTridentServiceAccount(getServiceAccountName(csi), appLabel,
		i.namespace); err != nil {
		return err
	}

	// If OpenShift, delete Trident Security Context Constraint(s)
	if i.client.Flavor() == k8sclient.FlavorOpenShift {
		if err := i.client.DeleteOpenShiftSCC(getOpenShiftSCCUserName(), getOpenShiftSCCName(),
			appLabelValue); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) ObliviateCRDs() error {
	return cmd.ObliviateCRDs(i.client, i.tridentCRDClient, k8sTimeout)
}
