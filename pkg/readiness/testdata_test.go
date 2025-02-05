/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package readiness_test

import (
	externaldatav1beta1 "github.com/open-policy-agent/frameworks/constraint/pkg/apis/externaldata/v1beta1"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	mutationsv1alpha1 "github.com/open-policy-agent/gatekeeper/apis/mutations/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Templates and constraints in testdata/.
var testTemplates = []*templates.ConstraintTemplate{
	makeTemplate("k8sallowedrepos"),
	makeTemplate("k8srequiredlabels"),
}

var testConstraints = []*unstructured.Unstructured{
	makeConstraint("ns-must-have-gk", "K8sRequiredLabels"),
	makeConstraint("prod-repo-is-openpolicyagent", "K8sAllowedRepos"),
}

// Templates and constraint in testdata/post/.
var postTemplates = []*templates.ConstraintTemplate{
	makeTemplate("k8shttpsonly"),
}

var postConstraints = []*unstructured.Unstructured{
	makeConstraint("ingress-https-only", "K8sHttpsOnly"),
}

var testAssignMetadata = []*mutationsv1alpha1.AssignMetadata{
	makeAssignMetadata("demo"),
}

var testModifySet = []*mutationsv1alpha1.ModifySet{
	makeModifySet("demo"),
}

var testAssignImage = []*mutationsv1alpha1.AssignImage{
	makeAssignImage("demo"),
}

var testAssign = []*mutationsv1alpha1.Assign{
	makeAssign("demo"),
}

var testProvider = []*externaldatav1beta1.Provider{
	makeProvider("demo"),
}

var testNS = makeNS("demo")

func makeTemplate(name string) *templates.ConstraintTemplate {
	return &templates.ConstraintTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
}

func makeConstraint(name string, kind string) *unstructured.Unstructured {
	u := unstructured.Unstructured{}
	u.SetAPIVersion("constraints.gatekeeper.sh/v1beta1")
	u.SetKind(kind)
	u.SetName(name)
	return &u
}

func makeAssignMetadata(name string) *mutationsv1alpha1.AssignMetadata {
	return &mutationsv1alpha1.AssignMetadata{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "mutations.gatekeeper.sh/v1alpha1",
			Kind:       "AssignMetadata",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: mutationsv1alpha1.AssignMetadataSpec{
			Location: "metadata.labels.demolabel",
		},
	}
}

func makeModifySet(name string) *mutationsv1alpha1.ModifySet {
	return &mutationsv1alpha1.ModifySet{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "mutations.gatekeeper.sh/v1alpha1",
			Kind:       "ModifySet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: mutationsv1alpha1.ModifySetSpec{
			Location: "spec.some.set",
		},
	}
}

func makeAssignImage(name string) *mutationsv1alpha1.AssignImage {
	return &mutationsv1alpha1.AssignImage{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "mutations.gatekeeper.sh/v1alpha1",
			Kind:       "AssignImage",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: mutationsv1alpha1.AssignImageSpec{
			Location: "spec.containers[name:*].image",
		},
	}
}

func makeAssign(name string) *mutationsv1alpha1.Assign {
	return &mutationsv1alpha1.Assign{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "mutations.gatekeeper.sh/v1alpha1",
			Kind:       "Assign",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: mutationsv1alpha1.AssignSpec{
			Location: "spec.dnsPolicy",
		},
	}
}

func makeProvider(name string) *externaldatav1beta1.Provider {
	return &externaldatav1beta1.Provider{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "externaldata.gatekeeper.sh/v1alpha1",
			Kind:       "Provider",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: externaldatav1beta1.ProviderSpec{
			URL:     "http://demo",
			Timeout: 1,
		},
	}
}

func makeDeployment(name string) *appsv1.Deployment {
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:latest",
						},
					},
				},
			},
		},
	}
}

func makeNS(name string) *corev1.Namespace {
	return &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
}
