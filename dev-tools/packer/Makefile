BUILDID=$(shell git rev-parse HEAD)

.PHONY: all
all: aibeat/deb aibeat/rpm aibeat/darwin aibeat/win aibeat/bin \
	build/upload/build_id.txt

ES_BEATS=../../vendor/github.com/elastic/beats
include $(ES_BEATS)/dev-tools/packer/scripts/Makefile


.PHONY: aibeat
aibeat: build/upload
	# cross compile on ubuntu
	docker run --rm \
		-v $(abspath build):/build \
		-v $(abspath $(ES_BEATS)/dev-tools/packer/xgo-scripts):/scripts \
		-v $(abspath ../..):/source \
		-e PACK=$@ \
		-e BEFORE_BUILD=before_build.sh \
		-e SOURCE=/source \
		-e BUILDID=${BUILDID} \
		tudorg/beats-builder \
		github.com/aryou//aibeat
	# linux builds on debian 6
	docker run --rm \
		-v $(abspath build):/build \
		-v $(abspath $(ES_BEATS)/dev-tools/packer/xgo-scripts):/scripts \
		-v $(abspath ../..):/source \
		-e PACK=$@ \
		-e BEFORE_BUILD=before_build.sh \
		-e SOURCE=/source \
		-e BUILDID=${BUILDID} \
		tudorg/beats-builder-deb6 \
		github.com/aryou//aibeat
