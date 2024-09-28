
# Set the version for the release
VERSION=$(git describe --tags)

# Define the architectures to build for
ARCHITECTURES=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64")

# Create a directory for the release
mkdir -p release

# Build for each architecture
for ARCH in "${ARCHITECTURES[@]}"; do
    OS=$(echo $ARCH | cut -d'/' -f1)
    ARCH=$(echo $ARCH | cut -d'/' -f2)
    OUTPUT="release/ifnew-$VERSION-$OS-$ARCH"

    echo "Building for $OS/$ARCH..."
    GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT
done

# Create a tarball for the release
tar -czvf ifnew-$VERSION.tar.gz -C release .


