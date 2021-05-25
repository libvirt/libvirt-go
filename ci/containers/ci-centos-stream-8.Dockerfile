# THIS FILE WAS AUTO-GENERATED
#
#  $ lcitool dockerfile centos-stream-8 libvirt+dist,libvirt-go
#
# https://gitlab.com/libvirt/libvirt-ci/-/commit/1c5d87ecd2283614a8b0c31cead0b6d7883afd28

FROM quay.io/centos/centos:stream8

RUN dnf update -y && \
    dnf install 'dnf-command(config-manager)' -y && \
    dnf config-manager --set-enabled -y powertools && \
    dnf install -y centos-release-advanced-virtualization && \
    dnf install -y epel-release && \
    dnf install -y \
        ca-certificates \
        ccache \
        gcc \
        git \
        glibc-devel \
        glibc-langpack-en \
        golang \
        libvirt-devel \
        pkgconfig && \
    dnf autoremove -y && \
    dnf clean all -y && \
    rpm -qa | sort > /packages.txt && \
    mkdir -p /usr/libexec/ccache-wrappers && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/cc && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/gcc

ENV LANG "en_US.UTF-8"
ENV CCACHE_WRAPPERSDIR "/usr/libexec/ccache-wrappers"
