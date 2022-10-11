include Makefile.defs

.PHONY: all
all:


# ----------------

define BUILD_BIN
echo "begin to build bin under $(CMD_BIN_DIR)" ; \
   mkdir -p $(DESTDIR_BIN) ; \
   BIN_NAME_LIST=$$( cd $(CMD_BIN_DIR) && ls ) ; \
   for BIN_NAME in $${BIN_NAME_LIST} ; do \
  		rm -f $(DESTDIR_BIN)/$${BIN_NAME} ; \
  		$(GO_BUILD) -o $(DESTDIR_BIN)/$${BIN_NAME}  $(CMD_BIN_DIR)/$${BIN_NAME}/main.go ; \
  		(($$?!=0)) && echo "error, failed to build $${BIN_NAME}" && exit 1 ; \
  		echo "succeeded to build '$${BIN_NAME}' to $(DESTDIR_BIN)/$${BIN_NAME}" ; \
  	 done
endef


.PHONY: build_all_bin
build_all_bin: CMD_BIN_DIR:=cmd
build_all_bin:
	@ $(BUILD_BIN)


# ==========================

define BUILD_FINAL_IMAGE
echo "Build Image $(IMAGE_NAME):$(IMAGE_TAG)" ; \
		docker build  \
				--build-arg RACE=1 \
				--build-arg NOSTRIP=1 \
				--build-arg NOOPT=1 \
				--build-arg GIT_COMMIT_VERSION=$(GIT_COMMIT_VERSION) \
				--build-arg GIT_COMMIT_TIME=$(GIT_COMMIT_TIME) \
				--build-arg VERSION=$(GIT_COMMIT_VERSION) \
				--file $(DOCKERFILE_PATH) \
				--tag ${IMAGE_NAME}:$(IMAGE_TAG) . ; \
		echo "build success for $${i}:$(IMAGE_TAG) "
endef


.PHONY: build_local_image
build_local_image: build_local_agent_image

.PHONY: build_local_agent_image
build_local_agent_image: IMAGE_NAME := ${REGISTER}/${GIT_REPO}/agent
build_local_agent_image: DOCKERFILE_PATH := $(ROOT_DIR)/images/agent/Dockerfile
build_local_agent_image: IMAGE_TAG := $(GIT_COMMIT_VERSION)
build_local_agent_image:
	@ $(BUILD_FINAL_IMAGE)


#---------

define BUILD_BASE_IMAGE
TAG=` git ls-tree --full-tree HEAD -- $(IMAGEDIR) | awk '{ print $$3 }' ` ; \
		echo "Build base image $(BASE_IMAGES):$${TAG}" ; \
		docker build  \
				--build-arg USE_PROXY_SOURCE=true \
				--file $(DOCKERFILE_PATH) \
				--output type=docker \
				--tag $(BASE_IMAGES):$${TAG}  $(IMAGEDIR) ; \
		(($$?==0)) || { echo "error , failed to build base image" ; exit 1 ;} ; \
		echo "build success $(BASE_IMAGES):$${TAG} "
endef

.PHONY: build_local_base_image
build_local_base_image: build_local_agent_base_image

.PHONY: build_local_agent_base_image
build_local_agent_base_image: DOCKERFILE_PATH := $(ROOT_DIR)/images/agent-base/Dockerfile
build_local_agent_base_image: BASE_IMAGES := ${REGISTER}/${GIT_REPO}/agent-base
build_local_agent_base_image:
	@ $(BUILD_BASE_IMAGE)



.PHONY: update_images_dockerfile_golang
update_images_dockerfile_golang:
	GO_VERSION=$(GO_VERSION) $(ROOT_DIR)/tools/images/update-golang-image.sh

