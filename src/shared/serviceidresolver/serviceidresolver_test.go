package serviceidresolver

import (
	"context"
	"errors"
	"fmt"
	"github.com/otterize/intents-operator/src/operator/api/v1alpha3"
	serviceidresolvermocks "github.com/otterize/intents-operator/src/shared/serviceidresolver/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
)

type MatchingLabelsSelectorMatcher struct {
	expected client.MatchingLabelsSelector
}

func (m *MatchingLabelsSelectorMatcher) Matches(x interface{}) bool {
	if x == nil {
		return false
	}
	matchingLabels, ok := x.(client.MatchingLabelsSelector)
	if !ok {
		return false
	}
	return m.expected.String() == matchingLabels.String()
}

func (m *MatchingLabelsSelectorMatcher) String() string {
	return m.expected.String()
}

type ServiceIdResolverTestSuite struct {
	suite.Suite
	Client   *serviceidresolvermocks.MockClient
	Resolver *Resolver
}

func (s *ServiceIdResolverTestSuite) SetupTest() {
	controller := gomock.NewController(s.T())
	s.Client = serviceidresolvermocks.NewMockClient(controller)
	s.Resolver = NewResolver(s.Client)
}

func (s *ServiceIdResolverTestSuite) TestResolveClientIntentToPod_PodExists() {
	serviceName := "coolservice"
	namespace := "coolnamespace"
	SAName := "backendservice"

	intent := v1alpha3.ClientIntents{Spec: &v1alpha3.IntentsSpec{Service: v1alpha3.Service{Name: serviceName}}, ObjectMeta: metav1.ObjectMeta{Namespace: namespace}}
	ls, err := intent.BuildPodLabelSelector()
	s.Require().NoError(err)

	pod := corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: serviceName, Namespace: namespace}, Spec: corev1.PodSpec{ServiceAccountName: SAName}}

	s.Client.EXPECT().List(
		gomock.Any(),
		gomock.AssignableToTypeOf(&corev1.PodList{}),
		&MatchingLabelsSelectorMatcher{client.MatchingLabelsSelector{Selector: ls}},
	).Do(func(_ any, podList *corev1.PodList, _ ...any) {
		podList.Items = append(podList.Items, pod)
	})

	resolvedPod, err := s.Resolver.ResolveClientIntentToPod(context.Background(), intent)
	resultSAName := resolvedPod.Spec.ServiceAccountName
	s.Require().NoError(err)
	s.Require().Equal(SAName, resultSAName)
}

func (s *ServiceIdResolverTestSuite) TestResolveClientIntentToPod_PodDoesntExist() {
	serviceName := "coolservice"
	namespace := "coolnamespace"

	intent := v1alpha3.ClientIntents{Spec: &v1alpha3.IntentsSpec{Service: v1alpha3.Service{Name: serviceName}}, ObjectMeta: metav1.ObjectMeta{Namespace: namespace}}
	ls, err := intent.BuildPodLabelSelector()
	s.Require().NoError(err)

	s.Client.EXPECT().List(
		gomock.Any(),
		gomock.AssignableToTypeOf(&corev1.PodList{}),
		&MatchingLabelsSelectorMatcher{client.MatchingLabelsSelector{Selector: ls}},
	).Do(func(_ any, podList *corev1.PodList, _ ...any) {})

	pod, err := s.Resolver.ResolveClientIntentToPod(context.Background(), intent)
	s.Require().Equal(err, ErrPodNotFound)
	s.Require().Equal(corev1.Pod{}, pod)
}

func (s *ServiceIdResolverTestSuite) TestGetPodAnnotatedName_PodExists() {
	podName := "coolpod"
	podNamespace := "coolnamespace"
	serviceName := "coolservice"

	pod := corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: podName, Namespace: podNamespace, Annotations: map[string]string{viper.GetString(serviceNameOverrideAnnotationKey): serviceName}}}
	id, err := s.Resolver.ResolvePodToServiceIdentity(context.Background(), &pod)
	s.Require().NoError(err)
	s.Require().Equal(serviceName, id.Name)
}

func (s *ServiceIdResolverTestSuite) TestGetPodAnnotatedName_PodMissingAnnotation() {
	podName := "coolpod"
	podNamespace := "coolnamespace"

	pod := corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: podName, Namespace: podNamespace}}

	id, err := s.Resolver.ResolvePodToServiceIdentity(context.Background(), &pod)
	s.Require().Nil(err)
	s.Require().Equal(podName, id.Name)
}

func (s *ServiceIdResolverTestSuite) TestDeploymentNameWithDotsReplacedByUnderscore() {
	deploymentName := "cool-versioned-application.4.2.0"
	podName := "cool-pod-1234567890-12345"
	serviceName := "cool-versioned-application_4_2_0"
	podNamespace := "cool-namespace"

	// Create a pod with reference to the deployment with dots in the name
	myPod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: podNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind:       "Deployment",
					Name:       deploymentName,
					APIVersion: "apps/v1",
				},
			},
		},
	}

	deploymentAsObject := unstructured.Unstructured{}
	deploymentAsObject.SetName(deploymentName)
	deploymentAsObject.SetNamespace(podNamespace)
	deploymentAsObject.SetKind("Deployment")
	deploymentAsObject.SetAPIVersion("apps/v1")

	emptyObject := &unstructured.Unstructured{}
	emptyObject.SetKind("Deployment")
	emptyObject.SetAPIVersion("apps/v1")
	s.Client.EXPECT().Get(gomock.Any(), types.NamespacedName{Name: deploymentName, Namespace: podNamespace}, emptyObject).Do(
		func(_ context.Context, _ types.NamespacedName, obj *unstructured.Unstructured, _ ...any) error {
			deploymentAsObject.DeepCopyInto(obj)
			return nil
		})

	service, err := s.Resolver.ResolvePodToServiceIdentity(context.Background(), &myPod)
	s.Require().NoError(err)
	s.Require().Equal(serviceName, service.Name)
}

func (s *ServiceIdResolverTestSuite) TestDeploymentReadForbidden() {
	deploymentName := "best-deployment-ever"
	podName := "cool-pod-1234567890-12345"
	podNamespace := "cool-namespace"

	// Create a pod with reference to the deployment with dots in the name
	myPod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: podNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind:       "Deployment",
					Name:       deploymentName,
					APIVersion: "apps/v1",
				},
			},
		},
	}

	emptyObject := &unstructured.Unstructured{}
	emptyObject.SetKind("Deployment")
	emptyObject.SetAPIVersion("apps/v1")

	forbiddenError := apierrors.NewForbidden(schema.GroupResource{Group: "apps", Resource: "Deployment"}, deploymentName, errors.New("forbidden"))
	s.Client.EXPECT().Get(gomock.Any(), types.NamespacedName{Name: deploymentName, Namespace: podNamespace}, emptyObject).Return(forbiddenError)

	service, err := s.Resolver.ResolvePodToServiceIdentity(context.Background(), &myPod)
	s.Require().NoError(err)
	s.Require().Equal(deploymentName, service.Name)
}

func (s *ServiceIdResolverTestSuite) TestDeploymentRead() {
	deploymentName := "best-deployment-ever"
	podName := "cool-pod-1234567890-12345"
	podNamespace := "cool-namespace"

	// Create a pod with reference to the deployment with dots in the name
	myPod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: podNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind:       "Deployment",
					Name:       deploymentName,
					APIVersion: "apps/v1",
				},
			},
		},
	}

	deploymentAsObject := unstructured.Unstructured{}
	deploymentAsObject.SetName(deploymentName)
	deploymentAsObject.SetNamespace(podNamespace)
	deploymentAsObject.SetKind("Deployment")
	deploymentAsObject.SetAPIVersion("apps/v1")

	emptyObject := &unstructured.Unstructured{}
	emptyObject.SetKind("Deployment")
	emptyObject.SetAPIVersion("apps/v1")
	s.Client.EXPECT().Get(gomock.Any(), types.NamespacedName{Name: deploymentName, Namespace: podNamespace}, emptyObject).Do(
		func(_ context.Context, _ types.NamespacedName, obj *unstructured.Unstructured, _ ...any) error {
			deploymentAsObject.DeepCopyInto(obj)
			return nil
		})

	service, err := s.Resolver.ResolvePodToServiceIdentity(context.Background(), &myPod)
	s.Require().NoError(err)
	s.Require().Equal(deploymentName, service.Name)
}

func (s *ServiceIdResolverTestSuite) TestJobWithNoParent() {
	jobName := "my-crappy-1001-job-name-1234567890-12345"
	podName := "cool-pod-1234567890-12345"
	podNamespace := "cool-namespace"
	imageName := "cool-image"

	// Create a pod with reference to the deployment with dots in the name
	myPod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: podNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind:       "Job",
					Name:       jobName,
					APIVersion: "batch/v1",
				},
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "cool-container",
					Image: fmt.Sprintf("353146681200.dkr.ecr.us-west-2.amazonaws.com/%s:some-tag", imageName),
				},
			},
		},
	}

	deploymentAsObject := unstructured.Unstructured{}
	deploymentAsObject.SetName(jobName)
	deploymentAsObject.SetNamespace(podNamespace)
	deploymentAsObject.SetKind("Job")
	deploymentAsObject.SetAPIVersion("batch/v1")

	emptyObject := &unstructured.Unstructured{}
	emptyObject.SetKind("Job")
	emptyObject.SetAPIVersion("batch/v1")
	s.Client.EXPECT().Get(gomock.Any(), types.NamespacedName{Name: jobName, Namespace: podNamespace}, emptyObject).Do(
		func(_ context.Context, _ types.NamespacedName, obj *unstructured.Unstructured, _ ...any) error {
			deploymentAsObject.DeepCopyInto(obj)
			return nil
		})

	viper.Set(useImageNameForServiceIDForJobs, false)
	service, err := s.Resolver.ResolvePodToServiceIdentity(context.Background(), &myPod)
	s.Require().NoError(err)
	s.Require().Equal(jobName, service.Name)

	s.Client.EXPECT().Get(gomock.Any(), types.NamespacedName{Name: jobName, Namespace: podNamespace}, emptyObject).Do(
		func(_ context.Context, _ types.NamespacedName, obj *unstructured.Unstructured, _ ...any) error {
			deploymentAsObject.DeepCopyInto(obj)
			return nil
		})

	viper.Set(useImageNameForServiceIDForJobs, true)
	service, err = s.Resolver.ResolvePodToServiceIdentity(context.Background(), &myPod)
	s.Require().NoError(err)
	s.Require().Equal(imageName, service.Name)
}

func (s *ServiceIdResolverTestSuite) TestUserSpecifiedAnnotationForServiceName() {
	annotationName := "coolAnnotationName"
	expectedEnvVarName := "OTTERIZE_SERVICE_NAME_OVERRIDE_ANNOTATION"
	_ = os.Setenv(expectedEnvVarName, annotationName)
	s.Require().Equal(annotationName, viper.GetString(serviceNameOverrideAnnotationKey))
	_ = os.Unsetenv(expectedEnvVarName)
	s.Require().Equal(ServiceNameOverrideAnnotationKeyDefault, viper.GetString(serviceNameOverrideAnnotationKey))
}

func TestServiceIdResolverTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceIdResolverTestSuite))
}
