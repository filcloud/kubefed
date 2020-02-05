## Generate helm chart artifact and repo index

```bash
helm package charts/kubefed/ --version 0.1.0-rc6-patch
mv kubefed-0.1.0-rc6-patch.tgz charts/
helm repo index charts/ --url https://github.com/filcloud/kubefed/releases/download/v0.1.0-rc6-patch --merge charts/index.yaml
```

Then, you should release `v0.1.0-rc6-patch` on `https://github.com/filcloud/kubefed`, with uploaded `kubefed-0.1.0-rc6-patch.tgz`.

## Install helm chart

```bash
helm repo add kubefed-charts https://raw.githubusercontent.com/filcloud/kubefed/filcloud/charts

kubectl create namespace kubefed-system

helm install kubefed kubefed-charts/kubefed --version 0.1.0-rc6-patch --namespace kubefed-system --set controllermanager.tag=canary

kubectl -n kubefed-system get all
```
