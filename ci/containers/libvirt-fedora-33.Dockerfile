FROM registry.fedoraproject.org/fedora:33

RUN dnf update -y && \
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
    mkdir -p /usr/libexec/ccache-wrappers && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/cc && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/$(basename /usr/bin/gcc)

ENV LANG "en_US.UTF-8"
ENV CCACHE_WRAPPERSDIR "/usr/libexec/ccache-wrappers"
