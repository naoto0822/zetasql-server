FROM ubuntu:bionic AS build-env
RUN apt-get update && apt-get install -y --no-install-recommends curl ca-certificates gnupg
RUN curl https://bazel.build/bazel-release.pub.gpg | apt-key add -
RUN echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN apt-get update && apt-get install --no-install-recommends -y make g++ bazel-3.7.2 git python3

RUN update-alternatives --install /usr/bin/bazel bazel /usr/bin/bazel-3.7.2 100
RUN update-alternatives --install /usr/bin/python python /usr/bin/python3 100

WORKDIR /work

# cache build
#COPY .bazelrc CROSSTOOL WORKSPACE /work/

#RUN bazel build '@com_google_zetasql//zetasql/public:sql_formatter'
#RUN bazel build '@com_google_zetasql//zetasql/public:parse_helpers'
#RUN bazel build '@com_google_zetasql//zetasql/public:analyzer'

#COPY BUILD *.cc *.h *.go pkg /work/
ADD . /work/
RUN cd /work && bazel build :zetasql-server

RUN ls /work
RUN ls /work/bazel-out
RUN ls /work/bazel-out/k8-fastbuild/bin

FROM gcr.io/distroless/cc
COPY --from=build-env /work/bazel-out/k8-fastbuild/bin/zetasql-server_/zetasql-server ./
ENTRYPOINT ["./zetasql-server"]
