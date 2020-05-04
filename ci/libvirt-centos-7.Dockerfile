FROM centos:7

RUN echo -e '[openvz]\n\
name=OpenVZ addons\n\
baseurl=https://download.openvz.org/virtuozzo/releases/openvz-7.0.11-235/x86_64/os/\n\
enabled=1\n\
gpgcheck=1\n\
skip_if_unavailable=0\n\
metadata_expire=6h\n\
priority=90\n\
includepkgs=libprl*' > /etc/yum.repos.d/openvz.repo && \
    echo -e '-----BEGIN PGP PUBLIC KEY BLOCK-----\n\
Version: GnuPG v2.0.22 (GNU/Linux)\n\
\n\
mI0EVl80nQEEAKrEeyeTCwrzS9kYedZ/sAc/GUqlb81C7pA9SaR3fyck5mVw1Ogk\n\
YdmNBPM2kY7QDxR9F0EpSpnxSCAXZXugsQ8KzZ0DRLVeBDQyGs9IGK5hI0zzxIil\n\
BzfvIexLiQQhLy7YlIi8Jt/uUqKkW0pIMNMGcduY97VATtczpncpkmSzABEBAAG0\n\
SFZpcnR1b3p6byBUZWFtIChHUEcga2V5IHNpZ25hdHVyZSBmb3IgcGFja2FnZXMp\n\
IDxzZWN1cml0eUB2aXJ0dW96em8uY29tPoi5BBMBAgAjBQJWXzSdAhsDBwsJCAcD\n\
AgEGFQgCCQoLBBYCAwECHgECF4AACgkQygt9GUTNrSruIgP/er70Eyo73A1gfrjv\n\
oPUkyo4rslVRZu3qqCwoMFtJc/Z/UxWgEka1buorlcGLa6eO/EZ49c0n+KGa4Kvt\n\
EUboIq0yEu5i0FyAj92ifm+hNhoAbGfm0cZ4/fD0oGr3l8OsQo4+iHX4xAPwFe7Y\n\
zABuB8I1ZDZ4OIp5tDfTTuF2LT24jQRWXzSdAQQAog2Aqb+Ptl68O7cQhWLjVGkj\n\
yyigZrdeReLx3HloKJPBeQ/kA6uvMJc/IYS3uppMWXv9v+QenS6uhP1TUJ2k9FvM\n\
t94MQZfALN7Vpf8AF+UeWu4Ru+y4BNzcFhrPhIFNFChOR2QqW6FkgE57D9I177NC\n\
oJMyrlNe8wcGa178An8AEQEAAYifBBgBAgAJBQJWXzSdAhsMAAoJEMoLfRlEza0q\n\
bKwD/3+OFVIEXnIv5XgdGRNX5fHggsUN1bb8gva7HANRlKdd4LD8foDM3F/yv/3V\n\
igG14D5EjKz56SaBDNgiI4++hOzb2M8jhAsR86jxkXFrrP1U3ZNRKg6av9DPFAPS\n\
WEiJKtQrZDJloqtyi/mmRa1VsV7RYR0VPJjhK/R8EQ7Ysshy\n\
=fRMg\n\
-----END PGP PUBLIC KEY BLOCK-----' > /etc/pki/rpm-gpg/RPM-GPG-KEY-OpenVZ && \
    rpm --import /etc/pki/rpm-gpg/RPM-GPG-KEY-OpenVZ && \
    yum install -y epel-release && \
    yum update -y && \
    yum install -y \
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
        glibc-common \
        glibc-devel \
        golang \
        libtool \
        libvirt-devel \
        lsof \
        make \
        net-tools \
        ninja-build \
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
    yum autoremove -y && \
    yum clean all -y && \
    mkdir -p /usr/libexec/ccache-wrappers && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/cc && \
    ln -s /usr/bin/ccache /usr/libexec/ccache-wrappers/$(basename /usr/bin/gcc)

RUN pip3 install \
         meson==0.49.0

ENV LANG "en_US.UTF-8"

ENV MAKE "/usr/bin/make"
ENV NINJA "/usr/bin/ninja-build"
ENV PYTHON "/usr/bin/python3"

ENV CCACHE_WRAPPERSDIR "/usr/libexec/ccache-wrappers"
