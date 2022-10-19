# rocktemplate

## copy

0. grep "====modify====" * -RHn --colour  and modify all of them

1. copy repo and replace all 'rocktemplate' to 'YourRepoName'

2. add e2e tools in test/scripts/installE2eTools.sh for github CI

3. image name 
    .github/workflows/call-release-image.ymal  line 30-35 , rename image name

4. github/workflows/auto-nightly-ci.yaml  line 3 

5. create badge for github/workflows/auto-nightly-ci.yaml, github/workflows/badge.yaml

6. spidernet.io  -> settings -> secrets -> actions -> grant secret to repo

7. repo -> packages -> package settings -> Change package visibility

8. repo -> settings -> pages -> add branch 'github_pages', directory 'docs'

9. repo -> settings -> branch -> add protection rules for 'main' and 'github_pages'

9 redefine CRD in pkg/k8s/v1, and `make update_crd_sdk`



## local develop

1. `make build_local_image`

2. `make e2e_init`

3. `make e2e_run`

4. check proscope, browser vists http://NodeIP:4040

5. apply cr

        cat <<EOF > mybook.yaml
        apiVersion: rocktemplate.spidernet.io/v1
        kind: Mybook
        metadata:
          name: test
        spec:
          ipVersion: 4
          subnet: "1.0.0.0/8"
        EOF
        kubectl apply -f mybook.yaml


