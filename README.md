# ACME webhook for Scaleway DNS

This webhook allows Scaleway users to use the DNS01 challenge solving when using `cert-manager` in kubernetes.

## Installation

To install with helm, run:

```bash
$ git clone https://github.com/touchifyapp/cert-manager-webhook-scaleway.git
$ cd cert-manager-webhook-scaleway/deploy/cert-manager-webhook-scaleway
$ helm install --name cert-manager-webhook-scaleway .
```

Without helm, use:

```bash
$ make rendered-manifest.yaml
$ kubectl apply -f _out/rendered-manifest.yaml
```

## Configuration

### Generate your secret key

Login to your Scaleway account and create a token from the `credentials` page. A `secret_key` and an `access_key` will be displayed on your screen, the `secret_key` will be used in your `kubernetes` secret.

Reference: https://www.scaleway.com/docs/generate-an-api-token/

### Create a kubernetes secret

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: scaleway-secret-key
type: Opaque
stringData:
  token: SECRET_KEY_FROM_SCALEWAY
```

### Create a new Issuer/ClusterIssuer

```yaml
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: letsencrypt-scaleway
  namespace: default
spec:
  acme:
    server: https://acme-staging-v02.api.letsencrypt.org/directory
    email: certmaster@company.com
    privateKeySecretRef:
      name: letsencrypt-scaleway-account-key
    solvers:
    - dns01:
        webhook:
          groupName: acme.company.com
          solverName: scaleway
          config:
            organizationId: 12345678-1234-1234-1234-123456789012
            secretKeySecretRef:
              name: scaleway-secret-key
              key: token
```

### Testing your issuer

```yaml
apiVersion: cert-manager.io/v1alpha2
kind: Certificate
metadata:
  name: test-letsencrypt-crt
  namespace: default
spec:
  secretName: company-com-tls
  commonName: company.com
  issuerRef:
    name: letsencrypt-scaleway
    kind: Issuer
  dnsNames:
  - company.com
  - www.company.com
```

## Contributing

### Running the test suite

First, you need to provide your own secret key:
1. Generate your secret key as explained below ([more info](#generate-your-secret-key))
2. Fill in the appropriate values in `testdata/scaleway-solver/secretkey.yml` and `testdata/scaleway-solver/config.json`

Then, you can run the test suite with:

```bash
$ TEST_ZONE_NAME=example.com make test
```
