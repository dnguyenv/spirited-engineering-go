# Build from golang-centos7 base image
# dnguyenv@us.ibm.com
FROM openshift/base-centos7

ARG GOLANG_VERSION
ENV GOLANG_VERSION ${GOLANG_VERSION:-1.8.1}
ENV BUILDER_VERSION 1.0
ENV HOME /opt/app-root
ENV GOPATH $HOME/gopath
ENV PATH $PATH:$GOROOT/bin:$GOBIN

LABEL io.k8s.description="Spirited Engineering Go s2i experiment. Based on GO ${GOLANG_VERSION}." \
      io.k8s.display-name="golang builder v${GOLANG_VERSION}" \
      io.openshift.expose-services="8080:http" \
      io.openshift.tags="openshift,golang"

# Install golang runtime on CentOS7
RUN curl https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz -o /tmp/go.tar.gz && tar -C /usr/local -zxf /tmp/go.tar.gz

COPY ./s2i/bin/ /usr/libexec/s2i

# Make the content of /opt/app-root owned by user 1001
RUN chown -R 1001:1001 /opt/app-root

WORKDIR /opt/app-root

# Default user was created in the openshift/base-centos7 image
USER 1001

# Expose/set the default port for this test application
EXPOSE 8080