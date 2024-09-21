package com.soal.controller;

import com.soal.model.Transaction;
import com.soal.model.User;
import com.soal.repository.TransactionRepository;
import com.soal.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.*;

import java.util.Optional;
import java.util.UUID;

@RestController
@RequestMapping("/api")
public class TransactionController {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private TransactionRepository transactionRepository;

    @PostMapping("/topup")
    public ResponseEntity<?> topUp(@RequestBody Transaction request, Authentication auth) {
        User user = getUserFromAuth(auth);
        if (user == null) {
            return ResponseEntity.status(401).body("Unauthenticated");
        }

        int balanceBefore = user.getBalance();
        user.setBalance(balanceBefore + request.getAmount());

        Transaction transaction = new Transaction();
        transaction.setUser(user);
        transaction.setType("CREDIT");
        transaction.setAmount(request.getAmount());
        transaction.setBalanceBefore(balanceBefore);
        transaction.setBalanceAfter(user.getBalance());
        transaction.setRemarks("Top-up");

        userRepository.save(user);
        transactionRepository.save(transaction);

        return ResponseEntity.ok(transaction);
    }

    @PostMapping("/pay")
    public ResponseEntity<?> pay(@RequestBody Transaction request, Authentication auth) {
        User user = getUserFromAuth(auth);
        if (user == null) {
            return ResponseEntity.status(401).body("Unauthenticated");
        }

        if (user.getBalance() < request.getAmount()) {
            return ResponseEntity.status(400).body("Balance is not enough");
        }

        int balanceBefore = user.getBalance();
        user.setBalance(balanceBefore - request.getAmount());

        Transaction transaction = new Transaction();
        transaction.setUser(user);
        transaction.setType("DEBIT");
        transaction.setAmount(request.getAmount());
        transaction.setBalanceBefore(balanceBefore);
        transaction.setBalanceAfter(user.getBalance());
        transaction.setRemarks(request.getRemarks());

        userRepository.save(user);
        transactionRepository.save(transaction);

        return ResponseEntity.ok(transaction);
    }

    // Util method to extract user from the authenticated token
    private User getUserFromAuth(Authentication auth) {
        UUID userId = UUID.fromString(auth.getName());
        return userRepository.findById(userId).orElse(null);
    }
}
