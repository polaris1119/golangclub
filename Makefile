GitCommitLog=`git log --pretty=oneline -n 1`
GitRelease=`git describe --dirty --all`
BuildTime=`date '+%Y-%m-%d %H:%M:%S'`
LdFlags="-X 'main.gitCommitLog=${GitCommitLog}' -X 'main.gitRelease=${GitRelease}' -X 'main.buildTime=${BuildTime}'"

build: fmt
	@echo "building project..."
	go build -ldflags ${LdFlags} github.com/polaris1119/golangclub
	@echo "build finished!"

fmt:
	gofmt -w .