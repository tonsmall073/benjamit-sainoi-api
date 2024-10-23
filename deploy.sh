IMAGE_NAME="tonsmall073/benjamitsainoi-backend"
VERSION_FILE="version.txt"

if [[ ! -f $VERSION_FILE ]]; then
    EXIST_FILE="FALSE"
    echo "Version file not found. Creating version file with initial version v1.0.0."
    echo "v1.0.0" > $VERSION_FILE
    VERSION="v1.0.0"
else
    EXIST_FILE="TRUE"
    VERSION=$(cat $VERSION_FILE)
    echo "Current version is $VERSION"
    echo "Choose version to bump:"
    echo "1. Major"
    echo "2. Minor"
    echo "3. Patch"
    read -rp "Enter your choice (1-3, or press Enter to skip): " choice

    if [[ -z "$choice" ]]; then
        echo "No version bump selected. Exiting."
        exit 0
    fi
fi

bump_version() {
    if [[ -z "$VERSION" ]]; then
        echo "Current version is empty. Setting to v1.0.0."
        VERSION="v1.0.0"
    fi

    major=$(echo "${VERSION#v}" | cut -d '.' -f 1)
    minor=$(echo "${VERSION#v}" | cut -d '.' -f 2)
    patch=$(echo "${VERSION#v}" | cut -d '.' -f 3)

    major=${major:-0}
    minor=${minor:-0}
    patch=${patch:-0}

    case $choice in
        1) 
            major=$((major + 1))
            minor=0
            patch=0
            echo "Updated major version to $major"
            ;;
        2) 
            minor=$((minor + 1))
            patch=0
            echo "Updated minor version to $minor"
            ;;
        3) 
            patch=$((patch + 1))
            echo "Updated patch version to $patch"
            ;;
        *) 
            echo "Invalid choice. No changes made."
            exit 0
            return
            ;;
    esac

    VERSION="v$major.$minor.$patch"
}

if [[ $EXIST_FILE == "TRUE" ]]; then
    bump_version
fi

echo "$VERSION" > "$VERSION_FILE"
echo "Final version to be written: $VERSION"

# ------------------------------------------Docker Deploy------------------------------------------

TAG_LATEST="latest"
TAG_VERSION="$VERSION"

docker build -t "$IMAGE_NAME:$TAG_LATEST" .
docker tag "$IMAGE_NAME:$TAG_LATEST" "$IMAGE_NAME:$TAG_VERSION"

docker push "$IMAGE_NAME:$TAG_LATEST"
docker push "$IMAGE_NAME:$TAG_VERSION"

docker-compose pull
docker-compose up -d

echo "Deployment complete with version $VERSION!"