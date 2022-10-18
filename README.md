# rocktemplate

## copy

1. copy repo and replace all 'rocktemplate' to 'YourRepoName'

2. add e2e tools in test/scripts/installE2eTools.sh for github CI

3. image name 
    .github/workflows/call-release-image.ymal  line 30-35 , rename image name

4. github/workflows/auto-nightly-ci.yaml  line 3 

5. create badge for github/workflows/auto-nightly-ci.yaml, github/workflows/badge.yaml

6. spidernet.io  -> settings -> secrets -> actions -> grant secret to repo

7. repo -> packages -> package settings -> Change package visibility

8. repo -> settings -> pages -> add branch 'github_pages', directory 'docs'

8. repo -> settings -> branch -> add protection rules for 'main' and 'github_pages'

9 define CRD in pkg/k8s/v1, update "./tools/golang/crdControllerGen.sh" "./tools/golang/crdSdkGen.sh" and `make update_crd_sdk`



## local develop

1. `make build_local_image`

2. `make e2e_init`

3. `make e2e_run`

4. check proscope, browser vists http://NodeIP:4040
