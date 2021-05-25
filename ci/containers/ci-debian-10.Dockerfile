# THIS FILE WAS AUTO-GENERATED
#
#  $ lcitool dockerfile debian-10 libvirt+dist,libvirt-go
#
# https://gitlab.com/libvirt/libvirt-ci/-/commit/1c5d87ecd2283614a8b0c31cead0b6d7883afd28

FROM docker.io/library/debian:10-slim

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y eatmydata && \
    eatmydata apt-get dist-upgrade -y && \
    eatmydata apt-get install --no-install-recommends -y \
            ca-certificates \
            ccache \
            gcc \
            git \
            golang \
            libc6-dev \
            libvirt-dev \
            locales \
            pkgconf && \
    eatmydata apt-get autoremove -y && \
    eatmydata apt-get autoclean -y && \
    sed -Ei 's,^# (en_US\.UTF-8 .*)$,\1,' /etc/locale.gen && \
    dpkg-reconfigure locales && \
    dpkg-query --showformat '${Package}_${Version}_${Architecture}\n' --show > /packages.txt && \
    mkdir -p /usr/libexec/ccache-wrappers && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/cc && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/gcc

ENV LANG "en_US.UTF-8"
ENV CCACHE_WRAPPERSDIR "/usr/libexec/ccache-wrappers"
