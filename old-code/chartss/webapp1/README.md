# webapp1

![Version: 0.1.1](https://img.shields.io/badge/Version-0.1.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.16.0](https://img.shields.io/badge/AppVersion-1.16.0-informational?style=flat-square)

A Helm chart for Kubernetes

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| appName | string | `"myhelmapp"` |  |
| namespace | string | `"default"` |  |
| configmap.name | string | `"helmappconfigmapv1.1"` |  |
| configmap.data.CUSTOM_HEADER | string | `"This app was deployed with Helm!"` |  |
| image.name | string | `"devopsjourney1/mywebapp"` |  |
| image.tag | string | `"latest"` |  |

