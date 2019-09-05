/*
 * Copyright (c) 2019, Oracle and/or its affiliates. All rights reserved.
 * Licensed under the Universal Permissive License v 1.0 as shown at
 * http://oss.oracle.com/licenses/upl.
 */

package v1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	coherence "github.com/oracle/coherence-operator/pkg/apis/coherence/v1"
)

var _ = Describe("Testing LoggingSpec struct", func() {

	Context("Copying a LoggingSpec using DeepCopyWithDefaults", func() {
		var original *coherence.LoggingSpec
		var defaults *coherence.LoggingSpec
		var clone *coherence.LoggingSpec

		JustBeforeEach(func() {
			clone = original.DeepCopyWithDefaults(defaults)
		})

		When("original and defaults are nil", func() {
			BeforeEach(func() {
				original = nil
				defaults = nil
			})

			It("the copy should be nil", func() {
				Expect(clone).Should(BeNil())
			})
		})

		When("defaults is nil", func() {
			BeforeEach(func() {
				original = &coherence.LoggingSpec{
					Level:          int32Ptr(9),
					ConfigFile:     stringPtr("logging.properties"),
					ConfigMapName:  stringPtr("loggingMap"),
					FluentdEnabled: boolPtr(true),
				}

				defaults = nil
			})

			It("should copy the original Level", func() {
				Expect(*clone.Level).To(Equal(*original.Level))
			})

			It("should copy the original ConfigFile", func() {
				Expect(*clone.ConfigFile).To(Equal(*original.ConfigFile))
			})

			It("should copy the original ConfigMapName", func() {
				Expect(*clone.ConfigMapName).To(Equal(*original.ConfigMapName))
			})

			It("should copy the original FluentdEnabled", func() {
				Expect(*clone.FluentdEnabled).To(Equal(*original.FluentdEnabled))
			})
		})

		When("original is nil", func() {
			BeforeEach(func() {
				defaults = &coherence.LoggingSpec{
					Level:          int32Ptr(9),
					ConfigFile:     stringPtr("logging.properties"),
					ConfigMapName:  stringPtr("loggingMap"),
					FluentdEnabled: boolPtr(true),
				}

				original = nil
			})

			It("should copy the defaults Level", func() {
				Expect(*clone.Level).To(Equal(*defaults.Level))
			})

			It("should copy the defaults ConfigFile", func() {
				Expect(*clone.ConfigFile).To(Equal(*defaults.ConfigFile))
			})

			It("should copy the defaults ConfigMapName", func() {
				Expect(*clone.ConfigMapName).To(Equal(*defaults.ConfigMapName))
			})

			It("should copy the defaults FluentdEnabled", func() {
				Expect(*clone.FluentdEnabled).To(Equal(*defaults.FluentdEnabled))
			})
		})

		When("all original fields are set", func() {
			BeforeEach(func() {
				original = &coherence.LoggingSpec{
					Level:          int32Ptr(9),
					ConfigFile:     stringPtr("logging.properties"),
					ConfigMapName:  stringPtr("loggingMap"),
					FluentdEnabled: boolPtr(true),
				}

				defaults = &coherence.LoggingSpec{
					Level:          int32Ptr(7),
					ConfigFile:     stringPtr("logging2.properties"),
					ConfigMapName:  stringPtr("loggingMap2"),
					FluentdEnabled: boolPtr(false),
				}
			})

			It("should copy the original Level", func() {
				Expect(*clone.Level).To(Equal(*original.Level))
			})

			It("should copy the original ConfigFile", func() {
				Expect(*clone.ConfigFile).To(Equal(*original.ConfigFile))
			})

			It("should copy the original ConfigMapName", func() {
				Expect(*clone.ConfigMapName).To(Equal(*original.ConfigMapName))
			})

			It("should copy the original FluentdEnabled", func() {
				Expect(*clone.FluentdEnabled).To(Equal(*original.FluentdEnabled))
			})
		})

		When("original Level is nil", func() {
			BeforeEach(func() {
				original = &coherence.LoggingSpec{
					Level:          nil,
					ConfigFile:     stringPtr("logging.properties"),
					ConfigMapName:  stringPtr("loggingMap"),
					FluentdEnabled: boolPtr(true),
				}

				defaults = &coherence.LoggingSpec{
					Level:          int32Ptr(7),
					ConfigFile:     stringPtr("logging2.properties"),
					ConfigMapName:  stringPtr("loggingMap2"),
					FluentdEnabled: boolPtr(false),
				}
			})

			It("should copy the defaults Level", func() {
				Expect(*clone.Level).To(Equal(*defaults.Level))
			})

			It("should copy the original ConfigFile", func() {
				Expect(*clone.ConfigFile).To(Equal(*original.ConfigFile))
			})

			It("should copy the original ConfigMapName", func() {
				Expect(*clone.ConfigMapName).To(Equal(*original.ConfigMapName))
			})

			It("should copy the original FluentdEnabled", func() {
				Expect(*clone.FluentdEnabled).To(Equal(*original.FluentdEnabled))
			})
		})

		When("original ConfigFile is nil", func() {
			BeforeEach(func() {
				original = &coherence.LoggingSpec{
					Level:          int32Ptr(9),
					ConfigFile:     nil,
					ConfigMapName:  stringPtr("loggingMap"),
					FluentdEnabled: boolPtr(true),
				}

				defaults = &coherence.LoggingSpec{
					Level:          int32Ptr(7),
					ConfigFile:     stringPtr("logging2.properties"),
					ConfigMapName:  stringPtr("loggingMap2"),
					FluentdEnabled: boolPtr(false),
				}
			})

			It("should copy the original Level", func() {
				Expect(*clone.Level).To(Equal(*original.Level))
			})

			It("should copy the defaults ConfigFile", func() {
				Expect(*clone.ConfigFile).To(Equal(*defaults.ConfigFile))
			})

			It("should copy the original ConfigMapName", func() {
				Expect(*clone.ConfigMapName).To(Equal(*original.ConfigMapName))
			})

			It("should copy the original FluentdEnabled", func() {
				Expect(*clone.FluentdEnabled).To(Equal(*original.FluentdEnabled))
			})
		})

		When("original ConfigMapName is nil", func() {
			BeforeEach(func() {
				original = &coherence.LoggingSpec{
					Level:          int32Ptr(9),
					ConfigFile:     stringPtr("logging.properties"),
					ConfigMapName:  nil,
					FluentdEnabled: boolPtr(true),
				}

				defaults = &coherence.LoggingSpec{
					Level:          int32Ptr(7),
					ConfigFile:     stringPtr("logging2.properties"),
					ConfigMapName:  stringPtr("loggingMap2"),
					FluentdEnabled: boolPtr(false),
				}
			})

			It("should copy the original Level", func() {
				Expect(*clone.Level).To(Equal(*original.Level))
			})

			It("should copy the original ConfigFile", func() {
				Expect(*clone.ConfigFile).To(Equal(*original.ConfigFile))
			})

			It("should copy the defaults ConfigMapName", func() {
				Expect(*clone.ConfigMapName).To(Equal(*defaults.ConfigMapName))
			})

			It("should copy the original FluentdEnabled", func() {
				Expect(*clone.FluentdEnabled).To(Equal(*original.FluentdEnabled))
			})
		})

		When("original FluentdEnabled is nil", func() {
			BeforeEach(func() {
				original = &coherence.LoggingSpec{
					Level:          int32Ptr(9),
					ConfigFile:     stringPtr("logging.properties"),
					ConfigMapName:  stringPtr("loggingMap"),
					FluentdEnabled: nil,
				}

				defaults = &coherence.LoggingSpec{
					Level:          int32Ptr(7),
					ConfigFile:     stringPtr("logging2.properties"),
					ConfigMapName:  stringPtr("loggingMap2"),
					FluentdEnabled: boolPtr(true),
				}
			})

			It("should copy the original Level", func() {
				Expect(*clone.Level).To(Equal(*original.Level))
			})

			It("should copy the original ConfigFile", func() {
				Expect(*clone.ConfigFile).To(Equal(*original.ConfigFile))
			})

			It("should copy the original ConfigMapName", func() {
				Expect(*clone.ConfigMapName).To(Equal(*original.ConfigMapName))
			})

			It("should copy the defaults FluentdEnabled", func() {
				Expect(*clone.FluentdEnabled).To(Equal(*defaults.FluentdEnabled))
			})
		})
	})
})
