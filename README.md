# rocktemplate

## copy

1. copy repo and replace all 'rocktemplate' to 'YourRepoName'

2. add e2e tools in test/scripts/installE2eTools.sh for github CI

3. image name 
    .github/workflows/call-release-image.ymal  line 30-35 , rename image name

4. github/workflows/auto-nightly-ci.yaml  line 3 

5. create badge for github/workflows/auto-nightly-ci.yaml, github/workflows/badge.yaml

## local develop

1. `make build_local_image`

2. `make e2e_init`

3. `make e2e_run`

4. check proscope, browser vists http://NodeIP:4040
