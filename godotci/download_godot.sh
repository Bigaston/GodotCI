#!/bin/bash

VERSION=$1
EXPORT_TEMPLATE_VERSION=$(echo "$VERSION" | sed 's/-/./g')

# Godot
wget "https://github.com/godotengine/godot-builds/releases/download/${VERSION}/Godot_v${VERSION}_linux.x86_64.zip"
unzip "Godot_v${VERSION}_linux.x86_64.zip"
mv "Godot_v${VERSION}_linux.x86_64" /usr/bin/godot
chmod +x /usr/bin/godot
rm "Godot_v${VERSION}_linux.x86_64.zip"
echo "${EXPORT_TEMPLATE_VERSION}" > /usr/bin/godot.version

# Export Templates
wget "https://github.com/godotengine/godot-builds/releases/download/${VERSION}/Godot_v${VERSION}_export_templates.tpz"
mkdir -p "/root/.local/share/godot/export_templates/"
unzip "Godot_v${VERSION}_export_templates.tpz"

cp -r "./templates/" "/root/.local/share/godot/export_templates/${EXPORT_TEMPLATE_VERSION}/"

rm "./Godot_v${VERSION}_export_templates.tpz"
rm -r "./templates"