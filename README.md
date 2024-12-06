# Edge Enablement Interview Application
## Prerequisites
- It is preferred to be running on Linux, MacOS, or have [WSL2 on Windows](https://learn.microsoft.com/en-us/windows/wsl/install)
- You will need to have a recent version of go (1.20+ is preferred) in your `$PATH`
  - [Download Go](https://go.dev/doc/install)
- You will need the Openshift CLI (oc) in your `$PATH`
  - [Linux x86](https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-linux.tar.gz)
  - [Mac x86](https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-mac.tar.gz)
  - [Apple Silicon Mac](https://mirror.openshift.com/pub/openshift-v4/aarch64/clients/ocp/stable/openshift-client-mac-arm64.tar.gz)
  - [Windows x86](https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-windows.zip)
- You will need an IDE such as [Visual Studio Code](https://code.visualstudio.com/download)
- You will need access to your terminal
- Having `make` installed will allow easier running of the necessary commands

## Instructions
1. Repository Prep  
  Fork this repository and clone your fork
  ```bash
  $ git clone https://github.com/your-user/ee-interview.git
  ```

  Once the repository has been cloned, change the git reference in config/openshift_build.yaml to match your repository

2. Set your kubeconfig to the file provided by your interviewer(s)
  ```bash
  $ export KUBECONFIG=kubeconfig.txt
  ```

3. Build the application in the cluster
  > **Note**  
  This will build the application via the Dockerfile and push it to the cluster

  ```bash
  ## Using make
  $ make oc-build

  ## Using oc
  $ oc apply -f config/openshift_build.yaml
  $ oc start-build serverapp

  ## You can follow along with the build log by supplying the build name to the oc command
  $ oc start-build serverapp
  build.build.openshift.io/serverapp-1 started

  $ oc logs build/serverapp-1 --follow
  ```

4. Deploy the application
  ```bash
  ## Using make
  $ make deploy

  ## Using oc
  $ oc apply -f config/openshift_deployment.yaml
  ...
  $ oc expose svc/serverapp
  ```

5. Once the application is deployed, your interviewer(s) will guide you through the next steps

6. (Optional) Cleanup  
  Your interviewer(s) will likely take care of cleanup for sake of time
  ```bash
  ## You can cleanup the deployments using the following commands:
  $ oc delete -f config/
  $ oc delete route serverapp
  ```