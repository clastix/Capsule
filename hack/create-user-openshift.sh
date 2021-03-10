#!/bin/bash

# This script uses Kubernetes CertificateSigningRequest (CSR) to a generate a
# certificate signed by the Kubernetes CA itself.
# It requires cluster admin permission.
#
# e.g.: ./create-user-openshift.sh alice oil
#       where `oil` is the Tenant and `alice` the owner

# Check if OpenSSL is installed
if [[ ! -x "$(command -v openssl)" ]]; then
    echo "Error: openssl not found"
    exit 1
fi

# Check if kubectl is installed
if [[ ! -x "$(command -v kubectl)" ]]; then
    echo "Error: kubectl not found"
    exit 1
fi

# Check if oc is installed
if [[ ! -x "$(command -v oc)" ]]; then
    echo "Error: kubectl not found"
    exit 1
fi

USER=$1
TENANT=$2

if [[ -z ${USER} ]]; then
    echo "User has not been specified!"
    exit 1
fi

if [[ -z ${TENANT} ]]; then
    echo "Tenant has not been specified!"
    exit 1
fi

GROUP=capsule.clastix.io

TMPDIR=$(mktemp -d)
echo "creating certs in TMPDIR ${TMPDIR} "

openssl genrsa -out ${TMPDIR}/tls.key 2048
openssl req -new -key ${TMPDIR}/tls.key -subj "/CN=${USER}/O=${GROUP}" -out ${TMPDIR}/${USER}-${TENANT}.csr

# Clean any previously created CSR for the same user.
kubectl delete csr ${USER}-${TENANT} 2>/dev/null || true

# Create a new CSR file.
cat <<EOF > ${TMPDIR}/${USER}-${TENANT}-csr.yaml
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: ${USER}-${TENANT}
spec:
  groups:
  - system:authenticated
  request: $(cat ${TMPDIR}/${USER}-${TENANT}.csr | base64 | tr -d '\n')
  usages:
  - digital signature
  - key encipherment
  - client auth
EOF

# Create the CSR
kubectl apply -f ${TMPDIR}/${USER}-${TENANT}-csr.yaml

# Approve and fetch the signed certificate
# In OCP why must use oc adm certificate approve
oc adm certificate approve ${USER}-${TENANT}
kubectl get csr ${USER}-${TENANT} -o jsonpath='{.status.certificate}' | base64 --decode > ${TMPDIR}/tls.crt


# Create the kubeconfig file
CONTEXT=$(kubectl config current-context)
CLUSTER=$(kubectl config view -o jsonpath="{.contexts[?(@.name == \"$CONTEXT\"})].context.cluster}")
SERVER=$(kubectl config view -o jsonpath="{.clusters[?(@.name == \"${CLUSTER}\"})].cluster.server}")

#create context for the new user:
oc config set-credentials ${USER} --client-certificate=${TMPDIR}/tls.crt --client-key=${TMPDIR}/tls.key --embed-certs --kubeconfig=${USER}-${TENANT}.kubeconfig

#set current context for new user
oc config set-context ${USER} --cluster=$(oc config view -o jsonpath='{.clusters[0].name}') --namespace=default --user=${USER}  --kubeconfig=${USER}-${TENANT}.kubeconfig

echo "kubeconfig file is:" ${USER}-${TENANT}.kubeconfig
echo "to use it as ${USER}: 'oc config use-context ${USER} --kubeconfig=${USER}-${TENANT}.kubeconfig'"
