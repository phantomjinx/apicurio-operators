package apis

import (
	api "github.com/apicurio/apicurio-operators/apicurito/pkg/apis/apicur/v1"
	"github.com/apicurio/apicurio-operators/apicurito/pkg/controller/apicurito"
	consolev1 "github.com/openshift/api/console/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, api.SchemeBuilder.AddToScheme)
	if err := apicurito.ConsoleYAMLSampleExists(); err == nil {
		AddToSchemes = append(AddToSchemes, consolev1.Install)
	}
}
