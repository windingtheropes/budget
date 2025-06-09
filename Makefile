TAG=windingtherope/budget-backend:latest
BUILD_PLATFORMS=linux/arm64,linux/amd64
build:
	docker buildx build --platform $(BUILD_PLATFORMS) -t $(TAG) . --push 