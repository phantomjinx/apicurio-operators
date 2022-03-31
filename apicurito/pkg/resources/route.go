/*
 * Copyright (C) 2020 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package resources

import (
	"github.com/RHsyseng/operator-utils/pkg/resource"

	api "github.com/apicurio/apicurio-operators/apicurito/pkg/apis/apicur/v1"
	routev1 "github.com/openshift/api/route/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func generatorRoute(a *api.Apicurito) (r resource.KubernetesResource) {

	hostname := a.Spec.GeneratorRouteHostname
	if len(hostname) == 0 {
		hostname = DefineGeneratorName(a)
	}

	r = &routev1.Route{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Route",
			APIVersion: routev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      DefineGeneratorName(a),
			Namespace: a.Namespace,
			Labels:    labels,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(a, schema.GroupVersionKind{
					Group:   api.SchemeGroupVersion.Group,
					Version: api.SchemeGroupVersion.Version,
					Kind:    a.Kind,
				}),
			},
		},
		Spec: routev1.RouteSpec{
			Host: hostname,
			Path: "/api/v1",
			TLS:  &routev1.TLSConfig{Termination: routev1.TLSTerminationEdge},
			To: routev1.RouteTargetReference{
				Kind: "Service",
				Name: DefineGeneratorName(a),
			},
		},
	}

	return
}

func apicuritoRoute(a *api.Apicurito) (r resource.KubernetesResource) {

	hostname := a.Spec.UIRouteHostname
	if len(hostname) == 0 {
		hostname = DefineUIName(a)
	}

	r = &routev1.Route{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Route",
			APIVersion: routev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      DefineUIName(a),
			Namespace: a.Namespace,
			Labels:    labels,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(a, schema.GroupVersionKind{
					Group:   api.SchemeGroupVersion.Group,
					Version: api.SchemeGroupVersion.Version,
					Kind:    a.Kind,
				}),
			},
		},
		Spec: routev1.RouteSpec{
			Host: hostname,
			TLS:  &routev1.TLSConfig{Termination: routev1.TLSTerminationEdge},
			To: routev1.RouteTargetReference{
				Kind: "Service",
				Name: DefineUIName(a),
			},
		},
	}

	return
}
