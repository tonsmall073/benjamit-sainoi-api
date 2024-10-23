@echo off
setlocal enabledelayedexpansion

set IMAGE_NAME=tonsmall073/benjamitsainoi-backend
set VERSION_FILE=version.txt

if not exist %VERSION_FILE% (
    set EXIST_FILE=FALSE
    echo Version file not found. Creating version file with initial version v1.0.0.
    echo v1.0.0 > %VERSION_FILE%
    set VERSION=v1.0.0
) else (
    set EXIST_FILE=TRUE
    set /p VERSION=<%VERSION_FILE%
    echo Current version is %VERSION%
    
    echo Choose version to bump:
    echo 1. Major
    echo 2. Minor
    echo 3. Patch
    set /p choice="Enter your choice (1-3, or press Enter to skip): "

    if "!choice!"=="" (
        echo No version bump selected. Exiting.
        exit /b
    )
)

if "%VERSION%"=="" (
    echo Current version is empty. Setting to v1.0.0.
    set VERSION=v1.0.0
)

if "%EXIST_FILE%"=="TRUE" (
    for /f "tokens=1,2,3 delims=." %%a in ("%VERSION:~1%") do (
        set major=%%a
        set minor=%%b
        set patch=%%c
    )

    if not defined major set major=0
    if not defined minor set minor=0
    if not defined patch set patch=0

    if "%choice%"=="1" (
        set /a major+=1
        set minor=0
        set patch=0
        echo Updated major version to !major!
    ) else if "%choice%"=="2" (
        set /a minor+=1
        set patch=0
        echo Updated minor version to !minor!
    ) else if "%choice%"=="3" (
        set /a patch+=1
        echo Updated patch version to !patch!
    ) else (
        echo Invalid choice. No changes made.
        exit /b
    )

    set VERSION=v!major!.!minor!.!patch!
)

echo !VERSION! > %VERSION_FILE%

echo Final version to be written: !VERSION!

@REM ------------------------------------------Docker Deploy------------------------------------------

set TAG_LATEST=latest
set TAG_VERSION=!VERSION!

docker build -t %IMAGE_NAME%:%TAG_LATEST% .
docker tag %IMAGE_NAME%:%TAG_LATEST% %IMAGE_NAME%:%TAG_VERSION%

docker push %IMAGE_NAME%:%TAG_LATEST%
docker push %IMAGE_NAME%:%TAG_VERSION%

docker-compose pull
docker-compose up -d

echo Deployment complete with version !VERSION!
exit /b
