FROM opensuse/leap:15.1

RUN zypper update -y && \
    zypper install -y \
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
           glibc-locale \
           go \
           libtool \
           libvirt-devel \
           lsof \
           make \
           net-tools \
           ninja \
           patch \
           perl \
           pkgconfig \
           python3 \
           python3-pip \
           python3-setuptools \
           python3-wheel \
           rpm-build \
           screen \
           strace \
           sudo \
           vim && \
    zypper clean --all && \
    mkdir -p /usr/libexec/ccache-wrappers && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/cc && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/$(basename /usr/bin/gcc)

RUN pip3 install \
         meson==0.49.0

ENV LANG "en_US.UTF-8"

ENV MAKE "/usr/bin/make"
ENV NINJA "/usr/bin/ninja"
ENV PYTHON "/usr/bin/python3"

ENV CCACHE_WRAPPERSDIR "/usr/libexec/ccache-wrappers"
