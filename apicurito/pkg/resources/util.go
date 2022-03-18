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
	"fmt"

	"github.com/apicurio/apicurio-operators/apicurito/pkg/apis/apicur/v1alpha1"
)

// DefineGeneratorName
// Define the name of the generator pod
func DefineGeneratorName(a *v1alpha1.Apicurito) string {
	return defineResourceName(a, "generator")
}

// DefineUIName
// Define the name of the UI pod
func DefineUIName(a *v1alpha1.Apicurito) string {
	return defineResourceName(a, "ui")
}

func defineResourceName(a *v1alpha1.Apicurito, suffix string) string {
	return fmt.Sprintf("%s-%s", a.Name, suffix)
}
