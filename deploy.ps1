$IMAGE_NAME = "tonsmall073/benjamitsainoi-backend"
$VERSION_FILE = "version.txt"

if (-Not (Test-Path $VERSION_FILE)) {
    $EXIST_FILE="FALSE"
    Write-Host "Version file not found. Creating version file with initial version v1.0.0."
    "v1.0.0" | Out-File $VERSION_FILE
    $VERSION = "v1.0.0" 
} else {
    $EXIST_FILE="TRUE"
    $VERSION = Get-Content $VERSION_FILE
    Write-Host "Current version is $VERSION"
    
    Write-Host "Choose version to bump:"
    Write-Host "1. Major"
    Write-Host "2. Minor"
    Write-Host "3. Patch"
    $choice = Read-Host "Enter your choice (1-3, or press Enter to skip)"
    
    if (-Not $choice) {
        Write-Host "No version bump selected. Exiting."
        exit 0
    }
}
if($EXIST_FILE -eq "TRUE"){
    function bump_version {
        param (
            [ref]$version
        )
    
        if (-Not $version.Value) {
            Write-Host "Current version is empty. Setting to v1.0.0."
            $version.Value = "v1.0.0"
        }
    
        $parts = $version.Value -replace 'v', '' -split '\.'
        $major = [int]$parts[0]
        $minor = [int]$parts[1]
        $patch = [int]$parts[2]
    
        switch ($choice) {
            1 {
                $major++
                $minor = 0
                $patch = 0
                Write-Host "Updated major version to $major"
            }
            2 {
                $minor++
                $patch = 0
                Write-Host "Updated minor version to $minor"
            }
            3 {
                $patch++
                Write-Host "Updated patch version to $patch"
            }
            default {
                Write-Host "Invalid choice. No changes made."
                exit 0
                return
            }
        }
    
        $version.Value = "v$major.$minor.$patch"
        Write-Host "New version set to: $($version.Value)"
    }
    
    bump_version ([ref]$VERSION)
}

Set-Content -Path $VERSION_FILE -Value $VERSION

Write-Host "Final version to be written: ${VERSION}"

# ------------------------------------------Docker Deploy------------------------------------------

$TAG_LATEST = "latest"
$TAG_VERSION = $VERSION

docker build -t "${IMAGE_NAME}:${TAG_LATEST}" .
docker tag "${IMAGE_NAME}:${TAG_LATEST}" "${IMAGE_NAME}:${TAG_VERSION}"

docker push "${IMAGE_NAME}:${TAG_LATEST}"
docker push "${IMAGE_NAME}:${TAG_VERSION}"

docker-compose pull
docker-compose up -d

Write-Host "Deployment complete with version ${VERSION}!"