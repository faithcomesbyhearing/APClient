# APClient_installer.ps1
param(
    [string]$AppName = "APClient"
)
Write-Host "Installing $AppName..." -ForegroundColor Green

# Create application directory
$appDir = "$env:LOCALAPPDATA\$AppName"
Write-Host "Creating directory: $appDir"
New-Item -ItemType Directory -Force -Path $appDir | Out-Null

# Copy executable (assuming it's in same directory as this script)
$exePath = Join-Path $PSScriptRoot "APClient.exe"
if (Test-Path $exePath) {
    Copy-Item $exePath -Destination "$appDir\$AppName.exe"
    Write-Host "Copied executable to $appDir"
} else {
    Write-Error "Could not find $AppName.exe in script directory"
    exit 1
}

# Create desktop shortcut
$WshShell = New-Object -comObject WScript.Shell
$desktopShortcut = $WshShell.CreateShortcut("$env:USERPROFILE\Desktop\$AppName.lnk")
$desktopShortcut.TargetPath = "$appDir\$AppName.exe"
$desktopShortcut.WorkingDirectory = $appDir
$desktopShortcut.Save()

# Create Start Menu shortcut
$startMenuShortcut = $WshShell.CreateShortcut("$env:APPDATA\Microsoft\Windows\Start Menu\Programs\$AppName.lnk")
$startMenuShortcut.TargetPath = "$appDir\$AppName.exe"
$startMenuShortcut.WorkingDirectory = $appDir
$startMenuShortcut.Save()

Write-Host "Installation complete!" -ForegroundColor Green
Write-Host "- Desktop shortcut created"
Write-Host "- Start Menu entry created"
Write-Host "- Application installed to: $appDir"
