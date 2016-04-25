# concourse-meetup-goserv

#### Description

This repository contain the demo use during the concourse meetup.
Basically is just an HTTP Server serving `/bar` which respond `hello world`

####Concourse

#####Pipeline

######Jobs
* 1 
  * Check develop branch and is there is new commit we run the unit-testing
  * If unit-testing on develop is passed
    * we deploy to PWS in the dev organisation (TBD)
* 2
  * Check new code on master, run unit-testing against
  * If unit-testing on master is successful run and we manually trigger `tag-master` job
    * We downstream master to develop
    * We create a github-release with the binary
    * We deploy to PWS (prod app) (TBD)

### Docker image use
* https://hub.docker.com/r/getourneau/alpine-bash-git/ https://hub.docker.com/r/getourneau/alpine-golang-bash/ https://github.com/shinji62/alpine-docker-bash-git
* https://github.com/shinji62/alpine-docker-bash-git

