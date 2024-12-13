package org.example.dockersprinng;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

@SpringBootApplication
@Controller
public class DockerSprinngApplication {

    public static void main(String[] args) {
        SpringApplication.run(DockerSprinngApplication.class, args);
    }

    // Alterando o mapeamento para /home
    @GetMapping("/home")
    public String index() {
        return "index";  // Retorna a página index (pode ser a mesma ou uma página diferente)
    }
}
