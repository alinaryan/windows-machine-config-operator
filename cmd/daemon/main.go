//go:build windows

/*
Copyright 2021.

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

package main

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	rootCmd = &cobra.Command{
		Use:   "wicd.exe",
		Short: "Windows Instance Config Daemon",
		Long: "The Windows Instance Config Daemon performs multiple functions related to maintaining the expected " +
			"state of a Windows Node.",
	}
	kubeconfig string
	namespace  string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "Path to kubeconfig with required permissions")
	rootCmd.MarkPersistentFlagRequired("kubeconfig")
	rootCmd.PersistentFlags().StringVar(&namespace, "namespace", "",
		"The namespace that required cluster resources, such as the ConfigMap, will be located in. This is the "+
			"namespace that WMCO is deployed in")
	rootCmd.MarkPersistentFlagRequired("namespace")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}
