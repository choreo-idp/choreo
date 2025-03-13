/*
 * Copyright (c) 2025, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package kubernetes

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"

	choreov1 "github.com/choreo-idp/choreo/api/v1"
	"github.com/choreo-idp/choreo/internal/dataplane"
	"github.com/choreo-idp/choreo/internal/ptr"
)

var _ = Describe("makeCronJob", func() {
	var (
		deployCtx *dataplane.DeploymentContext
		cronJob   *batchv1.CronJob
	)

	// Prepare fresh DeploymentContext before each test
	BeforeEach(func() {
		deployCtx = newTestDeploymentContext()
	})

	JustBeforeEach(func() {
		cronJob = makeCronJob(deployCtx)
	})

	Context("for a ScheduledTask component", func() {
		BeforeEach(func() {
			deployCtx.Component.Spec.Type = choreov1.ComponentTypeScheduledTask
		})

		It("should create a CronJob with correct name and namespace", func() {
			Expect(cronJob).NotTo(BeNil())
			Expect(cronJob.Name).To(Equal("my-component-my-main-track-a43a18e7"))
			Expect(cronJob.Namespace).To(Equal("dp-test-organiza-my-project-test-environ-04bdf416"))
		})

		expectedLabels := map[string]string{
			"organization-name":     "test-organization",
			"project-name":          "my-project",
			"environment-name":      "test-environment",
			"component-name":        "my-component",
			"component-type":        "ScheduledTask",
			"deployment-track-name": "my-main-track",
			"deployment-name":       "my-deployment",
			"managed-by":            "choreo-deployment-controller",
			"belong-to":             "user-workloads",
		}

		It("should create a CronJob with valid labels", func() {
			Expect(cronJob.Labels).To(BeComparableTo(expectedLabels))
		})

		It("should create a CronJob with correct Job labels", func() {
			Expect(cronJob.Spec.JobTemplate.Labels).To(BeComparableTo(expectedLabels))
		})

		It("should create a CronJob with correct Pod labels", func() {
			Expect(cronJob.Spec.JobTemplate.Spec.Template.Labels).To(BeComparableTo(expectedLabels))
		})

		It("should create a CronJob with correct Spec", func() {
			By("checking the schedule")
			// This will be empty if the configuration is not provided
			Expect(cronJob.Spec.Schedule).To(Equal(""))

			By("checking the concurrency policy")
			Expect(cronJob.Spec.ConcurrencyPolicy).To(Equal(batchv1.ForbidConcurrent))

			By("checking the timezone")
			Expect(cronJob.Spec.TimeZone).To(Equal(ptr.String("Etc/UTC")))
		})

		It("should create a CronJob with correct Job template", func() {
			By("checking the ActiveDeadlineSeconds")
			Expect(cronJob.Spec.JobTemplate.Spec.ActiveDeadlineSeconds).To(Equal(ptr.Int64(300)))

			By("checking the BackoffLimit")
			Expect(cronJob.Spec.JobTemplate.Spec.BackoffLimit).To(Equal(ptr.Int32(4)))

			By("checking the TTLSecondsAfterFinished")
			Expect(cronJob.Spec.JobTemplate.Spec.TTLSecondsAfterFinished).To(Equal(ptr.Int32(360)))
		})

		It("should create a CronJob with a correct container", func() {
			containers := cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers
			By("checking the container length")
			Expect(containers).To(HaveLen(1))

			By("checking the container")
			Expect(containers[0].Name).To(Equal("main"))
			Expect(containers[0].Image).To(Equal("my-image:latest"))
		})
	})

	Context("for a ScheduledTask component with a configuration", func() {
		BeforeEach(func() {
			deployCtx.Component.Spec.Type = choreov1.ComponentTypeScheduledTask
			deployCtx.DeployableArtifact.Spec.Configuration = &choreov1.Configuration{
				Application: &choreov1.Application{
					Task: &choreov1.TaskConfig{
						Disabled: true,
						Schedule: &choreov1.TaskSchedule{
							Cron:     "*/5 * * * *",
							Timezone: "Asia/Colombo",
						},
					},
				},
			}
		})

		It("should create a CronJob with correct schedule", func() {
			Expect(cronJob.Spec.Schedule).To(Equal("*/5 * * * *"))
		})

		It("should create a CronJob with correct timezone", func() {
			Expect(cronJob.Spec.TimeZone).To(Equal(ptr.String("Asia/Colombo")))
		})

		It("should create a CronJob with correct Suspend value", func() {
			Expect(cronJob.Spec.Suspend).To(Equal(ptr.Bool(true)))
		})
	})
})
