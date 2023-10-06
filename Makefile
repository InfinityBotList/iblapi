# Build revision should be unique for each build AND not the git commit hash
# because we want to be able to build the same commit multiple times and
# have the build revision be different each time.

BUILDTIME := $(shell date '+%Y-%m-%d | %H:%M:%S')
BUILDREV := ${BUILDTIME}
REPONAME := github.com/InfinityBotList/iblapi
PROJECTNAME := iblapi
GOFLAGS := -trimpath -ldflags="-s -w -X '$(REPONAME)/cmd.BuildRev=$(BUILDREV)' -X '$(REPONAME)/cmd.BuildTime=$(BUILDTIME)' -X '$(REPONAME)/cmd.ProjectName=$(PROJECTNAME)'"
GOFLAGS_DBG := -trimpath -ldflags="-X '$(REPONAME)/cmd.BuildRev=$(BUILDREV)' -X '$(REPONAME)/cmd.BuildTime=$(BUILDTIME)'"

COMBOS := linux/386 linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64 windows/386 freebsd/amd64

all:
	CGO_ENABLED=0 go build -v $(GOFLAGS_DBG)
install: all
	cp -rf iblapi /usr/bin/iblapi
publish:
	mkdir -p bin

	for combo in $(COMBOS); do \
		echo "$$combo"; \
		CGO_ENABLED=0 GOOS=$${combo%/*} GOARCH=$${combo#*/} go build -o bin/$$combo/iblapi $(GOFLAGS); \
		sha512sum bin/$$combo/iblapi > bin/$$combo/iblapi.sha512; \
	done

	# Rename all the windows binaries to .exe
	for folder in bin/windows/*; do \
		mv -vf $$folder/iblapi $$folder/iblapi.exe; \
	done

	rm -rf /silverpelt/cdn/ibl/dev/downloads/iblapi
	mkdir -p /silverpelt/cdn/ibl/dev/downloads/iblapi
	mv -vf bin/* /silverpelt/cdn/ibl/dev/downloads/iblapi
	echo -n "$(BUILDREV)" > /silverpelt/cdn/ibl/dev/downloads/iblapi/current_rev
	chown -Rv ibl:ibl /silverpelt/cdn/ibl/
	rm -rf bin
	
