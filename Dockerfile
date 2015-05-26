FROM java:7
COPY . /usr/src/myapp
WORKDIR /usr/src/myapp
# RUN javac MyHello.java
EXPOSE 80
CMD ["java", "-jar PlaneGame.jar"]