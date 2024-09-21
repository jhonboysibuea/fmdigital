package com.soal.controller;

import com.soal.model.User;
import com.soal.repository.UserRepository;
import com.soal.util.JwtUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;

import java.util.Optional;
import java.util.UUID;

@RestController
@RequestMapping("/api")
public class UserController {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @PostMapping("/register")
    public ResponseEntity<?> register(@RequestBody User user) {
        if (userRepository.existsByPhoneNumber(user.getPhoneNumber())) {
            return ResponseEntity.status(409).body("Phone Number already registered");
        }
        user.setPin(passwordEncoder.encode(user.getPin()));
        userRepository.save(user);
        return ResponseEntity.ok(user);
    }

    @PostMapping("/login")
    public ResponseEntity<?> login(@RequestBody User loginRequest) {
        Optional<User> user = userRepository.findByPhoneNumber(loginRequest.getPhoneNumber());
        if (user.isPresent() && passwordEncoder.matches(loginRequest.getPin(), user.get().getPin())) {
            String token = JwtUtils.generateToken(user.get().getId());
            return ResponseEntity.ok(token);
        }
        return ResponseEntity.status(401).body("Phone Number and PIN doesn't match.");
    }
}
