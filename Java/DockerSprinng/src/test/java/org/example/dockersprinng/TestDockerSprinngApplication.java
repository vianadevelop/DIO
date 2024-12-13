package org.example.dockersprinng;

import org.springframework.boot.SpringApplication;

public class TestDockerSprinngApplication {

    public static void main(String[] args) {
        SpringApplication.from(DockerSprinngApplication::main).with(TestcontainersConfiguration.class).run(args);
    }

}
