FROM centos:8

RUN dnf install 'dnf-command(config-manager)' -y && \
    dnf config-manager --set-enabled -y PowerTools && \
    dnf install -y epel-release && \
    dnf update -y && \
    dnf install -y \
        autoconf \
        automake \
        bash \
        bash-completion \
        ca-certificates \
        ccache \
        chrony \
        gcc \
        gdb \
        gettext \
        gettext-devel \
        git \
        glib2-devel \
        glibc-devel \
        glibc-langpack-en \
        gnutls-devel \
        golang \
        libnl3-devel \
        libtirpc-devel \
        libtool \
        libvirt-devel \
        libxml2 \
        libxml2-devel \
        libxslt \
        lsof \
        make \
        net-tools \
        ninja-build \
        patch \
        perl \
        perl-App-cpanminus \
        pkgconfig \
        python3 \
        python3-docutils \
        python3-pip \
        python3-setuptools \
        python3-wheel \
        rpcgen \
        rpm-build \
        screen \
        strace \
        sudo \
        vim \
        xz && \
    dnf autoremove -y && \
    dnf clean all -y && \
    mkdir -p /usr/libexec/ccache-wrappers && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/cc && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/$(basename /usr/bin/gcc)

RUN pip3 install \
         meson==0.54.0

ENV LANG "en_US.UTF-8"

ENV MAKE "/usr/bin/make"
ENV NINJA "/usr/bin/ninja"
ENV PYTHON "/usr/bin/python3"

ENV CCACHE_WRAPPERSDIR "/usr/libexec/ccache-wrappers"
