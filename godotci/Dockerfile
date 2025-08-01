FROM ubuntu AS desktop

ARG VERSION=4.3-stable

# Install required packages
RUN apt update
RUN apt install -y zip
RUN apt install -y curl
RUN apt install -y wget
RUN apt install -y fontconfig
RUN apt install -y git

RUN mkdir /godot

COPY download_godot.sh .
RUN sh download_godot.sh ${VERSION}
RUN rm download_godot.sh

# Clean APT
RUN rm -rf /var/lib/apt/lists/*

RUN godot --version

FROM desktop AS android
RUN apt update
RUN apt install -y openjdk-17-jdk

RUN mkdir /android
RUN wget https://dl.google.com/android/repository/commandlinetools-linux-13114758_latest.zip -O commandlinetools.zip
RUN unzip commandlinetools.zip 
RUN rm commandlinetools.zip
RUN mv cmdline-tools /android
RUN yes | /android/cmdline-tools/bin/sdkmanager --licenses --sdk_root=/android
RUN /android/cmdline-tools/bin/sdkmanager --sdk_root=/android "platform-tools" "build-tools;34.0.0" "platforms;android-34" "cmdline-tools;latest" "cmake;3.10.2.4988404" "ndk;23.2.8568313"

COPY debug.keystore /android/debug.keystore
COPY editor_settings /root/.config/godot/editor_settings
RUN cat /root/.config/godot/editor_settings
RUN version="${VERSION}"; short_version="$(echo "$version" | grep -oE '^[0-9]+\.[0-9]+')"; mv /root/.config/godot/editor_settings /root/.config/godot/editor_settings-$short_version.tres

# Clean APT
RUN rm -rf /var/lib/apt/lists/*