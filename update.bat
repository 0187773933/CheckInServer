@echo off
setlocal

:: Define variables
set "REPO_URL=https://github.com/0187773933/CheckInServer"
set "BASE_DIR=C:\APPLICATIONS_2"
set "APP_DIR=%BASE_DIR%\CheckInServer"

:: Ensure base directory exists
if not exist "%BASE_DIR%" (
    mkdir "%BASE_DIR%"
)

:: Check if the repository exists, clone if not
if not exist "%APP_DIR%" (
    echo Cloning repository...
    git clone "%REPO_URL%" "%APP_DIR%"
)

:: Navigate to application directory
cd /d "%APP_DIR%"

:: Remove existing main.exe if it exists
if exist main.exe (
    del main.exe
)

:: Pull latest changes
git pull origin main

:: Build the Go project for Windows
go build -o main.exe

echo Build completed.

endlocal