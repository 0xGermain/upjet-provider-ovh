// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	database "saagie.io/provider-ovh/internal/controller/database/database"
	iprestrictions "saagie.io/provider-ovh/internal/controller/kube/iprestrictions"
	kube "saagie.io/provider-ovh/internal/controller/kube/kube"
	nodepool "saagie.io/provider-ovh/internal/controller/kube/nodepool"
	providerconfig "saagie.io/provider-ovh/internal/controller/providerconfig"
	s3credentials "saagie.io/provider-ovh/internal/controller/user/s3credentials"
	s3policy "saagie.io/provider-ovh/internal/controller/user/s3policy"
	user "saagie.io/provider-ovh/internal/controller/user/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		iprestrictions.Setup,
		kube.Setup,
		nodepool.Setup,
		providerconfig.Setup,
		s3credentials.Setup,
		s3policy.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
