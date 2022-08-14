FROM fedora:36

# Install sys deps
RUN yum -y install vim golang git tinygo

# Mountpoint for actual code
RUN mkdir /src

