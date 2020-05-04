FROM fedora:rawhide

RUN dnf update -y --nogpgcheck fedora-gpg-keys && \
    dnf update -y && \
    dnf install -y \
        autoconf \
        automake \
        bash \
        bash-completion \
        ca-certificates \
        ccache \
        chrony \
        cppi \
        gcc \
        gdb \
        gettext \
        gettext-devel \
        git \
        glibc-devel \
        glibc-langpack-en \
        golang \
        libtool \
        libvirt-devel \
        lsof \
        make \
        meson \
        net-tools \
        ninja-build \
        patch \
        perl \
        pkgconfig \
        python3 \
        python3-setuptools \
        python3-wheel \
        rpm-build \
        screen \
        strace \
        sudo \
        vim && \
    dnf autoremove -y && \
    dnf clean all -y && \
    mkdir -p /usr/libexec/ccache-wrappers && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/cc && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/$(basename /usr/bin/gcc)

ENV LANG "en_US.UTF-8"

ENV MAKE "/usr/bin/make"
ENV NINJA "/usr/bin/ninja"
ENV PYTHON "/usr/bin/python3"

ENV CCACHE_WRAPPERSDIR "/usr/libexec/ccache-wrappers"
