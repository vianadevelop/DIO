package org.example.dockersprinng.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.ui.Model;

@Controller
public class HelloController {

    @GetMapping("/")
    public String hello(Model model) {
        model.addAttribute("message", "Hello, Docker!");
        return "index"; // Referência ao arquivo HTML em src/main/resources/templates/
    }
}
