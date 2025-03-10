#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

# Function to install a Helm chart and wait for its readiness
install_helm_chart() {
  local chart_dir=$1
  local release_name=$2
  local namespace=$3

  echo "Installing Helm chart from $chart_dir..."
  helm dependency update "$chart_dir"
  helm upgrade --install "$release_name" "$chart_dir" --namespace "$namespace" --create-namespace --timeout 30m
}

# Install helm chart for cilium
install_helm_chart "$SCRIPT_DIR/helm/cilium" "cilium" "choreo-system"

# Install choreo-opensource-dp
install_helm_chart "$SCRIPT_DIR/helm/choreo" "choreo" "choreo-system"

echo "Both Helm charts have been installed successfully! Please note that completing the full installation process may take some time."
